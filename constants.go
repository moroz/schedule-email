package main

import (
	"log"
	"os"
)

func MustGetenv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("FATAL: Environment variable %s is not set!", key)
	}
	return value
}

var API_TOKEN = MustGetenv("JMAP_TOKEN")
var USER_ID = MustGetenv("JMAP_USER_ID")

const SESSION_URL = "https://api.fastmail.com/jmap/session"
const API_URL = "https://api.fastmail.com/jmap/api/"

var DRAFT_MAILBOX_ID = MustGetenv("DRAFT_MAILBOX_ID")
