package main

import (
	"bytes"
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
var MATRIX_ACCOUNT_CHANNEL string
var MATRIX_REPORT_CHANNEL string
var PORT string
var IP2LOCATION_FILE string

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

	MATRIX_ACCOUNT_CHANNEL = os.Getenv("MATRIX_ACCOUNT_CHANNEL")
	if MATRIX_ACCOUNT_CHANNEL == "" {
		log.Fatal("MATRIX_ACCOUNT_CHANNEL empty or invalid")
	}

	MATRIX_REPORT_CHANNEL = os.Getenv("MATRIX_REPORT_CHANNEL")
	if MATRIX_REPORT_CHANNEL == "" {
		log.Fatal("MATRIX_REPORT_CHANNEL empty or invalid")
	}

	PORT = os.Getenv("PORT")
	if PORT == "" {
		log.Fatal("PORT empty or invalid")
	}

	IP2LOCATION_FILE = os.Getenv("IP2LOCATION_FILE")
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
		var i MastodonEvent
		err := json.NewDecoder(r.Body).Decode(&i)
		if err != nil {
			log.Println(err.Error())
			return
		}

		if i.Event == "report.created" {
			err = sendWebhook(
				fmt.Sprintf(
					"[New Report](%s): **%s** has reported **%s**: %s",
					fmt.Sprintf(
						"https://mastodon.nz/admin/reports/%s",
						i.Object.ID,
					),
					i.Object.Account.Username,
					i.Object.TargetAccount.Username,
					i.Object.Comment,
				),
				MATRIX_REPORT_CHANNEL,
			)
			if err != nil {
				log.Println(err.Error())
				return
			}
		} else if i.Event == "account.created" {
			country := ipLookup(i.Object.IP)
			err = sendWebhook(
				fmt.Sprintf(
					"[New Signup](%s) %s: **%s** (%s). %s",
					fmt.Sprintf(
						"https://mastodon.nz/admin/accounts/%s",
						i.Object.ID,
					),
					country,
					i.Object.Username,
					i.Object.Email,
					fmt.Sprintf(
						"Notes: %s",
						i.Object.Notes,
					),
				),
				MATRIX_ACCOUNT_CHANNEL,
			)
			if err != nil {
				log.Println(err.Error())
				return
			}
		}
	}
}

// sendWebhook - takes msg, sends to matrix
func sendWebhook(msgText string, channel string) error {
	// log.Println(msgText)
	data := MatrixWebhook{
		Key: MATRIX_WEBHOOK_API_KEY,
	}
	data.Body = msgText
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// log.Printf("Sending %s to %s", b, MATRIX_WEBHOOK_URL+"/"+channel)
	req, err := http.NewRequest("POST", MATRIX_WEBHOOK_URL+"/"+channel, bytes.NewBuffer(b))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Print(err)
		return nil
	}

	defer resp.Body.Close()

	return nil
}

// Lookup to country!
func ipLookup(ip string) string {
	if IP2LOCATION_FILE == "" {
		return ""
	}

	db, err := ip2location.OpenDB(IP2LOCATION_FILE)
	if err != nil {
		log.Print(err)
		return ""
	}
	results, err := db.Get_all(ip)

	if err != nil {
		log.Print(err)
		return ""
	}

	return " from " + results.Country_long
}
