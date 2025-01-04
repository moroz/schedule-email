package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func MustGetenv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("FATAL: Environment variable %s is not set!", key)
	}
	return value
}

var IMAP_TOKEN = MustGetenv("IMAP_TOKEN")

const SESSION_URL = "https://api.fastmail.com/jmap/session"

func main() {
	req, err := http.NewRequest("GET", SESSION_URL, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Authorization", "Bearer "+IMAP_TOKEN)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		log.Fatalf("API Returned response code %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	var result any
	json.NewDecoder(resp.Body).Decode(&result)

	fmt.Printf("%#v\n", result)
}
