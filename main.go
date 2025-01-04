package main

import (
	"log"
	"os"

	"github.com/emersion/go-imap/v2/imapclient"
)

func MustGetenv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("FATAL: Environment variable %s is not set!", key)
	}
	return value
}

var IMAP_SERVER = MustGetenv("IMAP_SERVER")
var IMAP_USER = MustGetenv("IMAP_USER")
var IMAP_PASSWORD = MustGetenv("IMAP_PASSWORD")

func main() {
	client, err := imapclient.DialTLS(IMAP_SERVER+":993", &imapclient.Options{})
	if err != nil {
		log.Fatal(err)
	}
	_ = client

	cmd := client.Login(IMAP_USER, IMAP_PASSWORD)
	err = cmd.Wait()
	if err != nil {
		log.Fatal(err)
	}
}
