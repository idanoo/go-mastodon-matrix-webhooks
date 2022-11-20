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
var MATRIX_CHANNEL string
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

	MATRIX_CHANNEL = os.Getenv("MATRIX_CHANNEL")
	if MATRIX_CHANNEL == "" {
		log.Fatal("MATRIX_CHANNEL empty or invalid")
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
			err = sendWebhook("New report!")
			if err != nil {
				log.Println(err.Error())
				return
			}
		} else if i.Event == "account.created" {
			country := ipLookup(i.Object.IP)
			err = sendWebhook(
				fmt.Sprintf(
					"New Signup%s: %s (%s) has signed up.\n%s\n%s",
					country,
					i.Object.Username,
					i.Object.Email,
					fmt.Sprintf(
						"Notes: %s",
						i.Object.Notes,
					),
					fmt.Sprintf(
						"https://mastodon.nz/admin/accounts/%s",
						i.Object.ID,
					),
				),
			)
			if err != nil {
				log.Println(err.Error())
				return
			}
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

	log.Printf("Sending %s to %s", b, MATRIX_WEBHOOK_URL+"/"+MATRIX_CHANNEL)
	req, err := http.NewRequest("POST", MATRIX_WEBHOOK_URL+"/"+MATRIX_CHANNEL, bytes.NewBuffer(b))
	if err != nil {
		return err
	}

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

	return " from" + results.Country_long
}
