package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/foxcpp/go-jmap"
)

func PerformJMAPCall(request jmap.Request) (*jmap.Response, error) {
	var bodyBuf bytes.Buffer
	err := json.NewEncoder(&bodyBuf).Encode(request)
	if err != nil {
		return nil, fmt.Errorf("Failed to encode payload to JSON: %w", err)
	}

	req, err := http.NewRequest("POST", API_URL, &bodyBuf)
	if err != nil {
		return nil, fmt.Errorf("Failed to initialize HTTP request: %w", err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+API_TOKEN)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("JMAP request failed: %w", err)
	}

	defer resp.Body.Close()
	var respBody jmap.Response
	err = respBody.Unmarshal(resp.Body, Unmarshallers)
	if err != nil {
		return nil, fmt.Errorf("Failed to decode JMAP response: %w", err)
	}
	return &respBody, nil
}

func main() {
	resp, err := PerformJMAPCall(jmap.Request{
		Using: []string{"urn:ietf:params:jmap:core", "urn:ietf:params:jmap:mail"},
		Calls: []jmap.Invocation{
			{
				Name: "Mailbox/query",
				Args: map[string]string{
					"accountId": USER_ID,
				},
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	resp, err = PerformJMAPCall(jmap.Request{
		Using: []string{"urn:ietf:params:jmap:core", "urn:ietf:params:jmap:mail"},
		Calls: []jmap.Invocation{
			{Name: "Mailbox/get", Args: map[string]any{
				"accountId": USER_ID,
				"ids":       nil,
			}},
		},
	})
	fmt.Printf("%#v\n", resp)
	fmt.Println(err)
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.SetEscapeHTML(false)
	enc.Encode(resp)
}
