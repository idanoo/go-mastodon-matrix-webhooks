package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ip2location/ip2location-go/v9"
	"github.com/joho/godotenv"
)

var MATRIX_WEBHOOK_URL string
var MATRIX_WEBHOOK_API_KEY string
var MATRIX_CHANNEL string
var PORT string
var IP2LOCATION_API_KEY string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	MATRIX_WEBHOOK_URL = os.Getenv("MATRIX_WEBHOOK_URL")
	if MATRIX_WEBHOOK_URL == "" {
		log.Fatal("MATRIX_WEBHOOK_URL empty or invalid")
	}

	MATRIX_WEBHOOK_API_KEY = os.Getenv("MATRIX_WEBHOOK_API_KEY")
	if MATRIX_WEBHOOK_API_KEY == "" {
		log.Fatal("MATRIX_WEBHOOK_API_KEY empty or invalid")
	}

	MATRIX_CHANNEL = os.Getenv("MATRIX_CHANNEL")
	if MATRIX_CHANNEL == "" {
		log.Fatal("MATRIX_CHANNEL empty or invalid")
	}

	PORT = os.Getenv("PORT")
	if PORT == "" {
		log.Fatal("PORT empty or invalid")
	}

	IP2LOCATION_API_KEY = os.Getenv("IP2LOCATION_API_KEY")
}

func main() {
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":"+PORT, nil); err != nil {
		log.Fatal(err)
	}
}

// Handle requests
func handler(w http.ResponseWriter, r *http.Request) {
	if r.Body != nil {
		var i IdentifyingRequest
		err := json.NewDecoder(r.Body).Decode(&i)
		if err != nil {
			log.Println(err.Error())
			return
		}

		if i.Event == "report.created" {
			var report MastodonReportEvent
			err := json.NewDecoder(r.Body).Decode(&report)
			if err != nil {
				log.Println(err.Error())
				return
			}
			go sendWebhook("New report!")
		} else if i.Event == "account.created" {
			var account MastodonSignUpEvent
			err := json.NewDecoder(r.Body).Decode(&account)
			if err != nil {
				log.Println(err.Error())
				return
			}
			country := ipLookup(account.Object.IP)
			go sendWebhook(fmt.Sprintf("*New Signup* %s has joined from %s", account.Object.Username, country))
		}
	}
}

// sendWebhook - takes msg, sends to matrix
func sendWebhook(msgText string) error {
	log.Println(msgText)

	data := MatrixWebhook{
		Key: MATRIX_WEBHOOK_API_KEY,
	}
	data.Body = msgText
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", MATRIX_WEBHOOK_URL+"/"+MATRIX_CHANNEL, b)
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	return nil
}

// Lookup to country!
func ipLookup(ip string) string {
	if IP2LOCATION_API_KEY == "" {
		return ""
	}

	apipackage := "WS25"
	usessl := true
	addon := "continent,country,region,city,geotargeting,country_groupings,time_zone_info" // leave blank if no need
	lang := "en"                                                                           // leave blank if no need

	ws, err := ip2location.OpenWS(IP2LOCATION_API_KEY, apipackage, usessl)

	if err != nil {
		fmt.Print(err)
		return ""
	}

	res, err := ws.LookUp(ip, addon, lang)

	if err != nil {
		fmt.Print(err)
		return ""
	}

	if res.Response != "OK" {
		fmt.Printf("Error: %s\n", res.Response)
	}

	return res.Country.Name
}
