package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var MATRIX_WEBHOOK_URL string
var MATRIX_WEBHOOK_API_KEY string
var MATRIX_CHANNEL string
var PORT string

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
}

func main() {
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":"+PORT, nil); err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	var bodyBytes []byte
	var err error

	if r.Body != nil {
		bodyBytes, err = io.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("Body reading error: %v", err)
			return
		}
		defer r.Body.Close()
	}

	if len(bodyBytes) > 0 {
		var i IdentifyingRequest
		err := json.NewDecoder(r.Body).Decode(&i)
		if err != nil {
			log.Println(err.Error())
			return
		}

		if i.Event == "report.created" {
			log.Println("New report!")
		} else if i.Event == "account.created" {
			log.Println("New account!")
		}
	}
}
