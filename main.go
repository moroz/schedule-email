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

func PrettyPrintJSONResponse(resp any) {
	enc := json.NewEncoder(os.Stdout)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "  ")
	enc.Encode(resp)
}

func main() {
	req := (jmap.Request{
		Using: []string{"urn:ietf:params:jmap:core", "urn:ietf:params:jmap:mail"},
		Calls: []jmap.Invocation{
			{
				Name: "Email/set",
				Args: map[string]any{
					"accountId": USER_ID,
					"create": map[string]any{
						"draft": MessageParams{
							MailboxIDs: map[string]bool{
								DRAFT_MAILBOX_ID: true,
							},
							Keywords: map[string]bool{
								"$draft": true,
							},
							From: []Address{
								{
									Name:  "Karol Moroz",
									Email: "karol@moroz.dev",
								},
							},
							To: []Address{
								{
									Name:  "Karol Moroz",
									Email: "recipient@moroz.dev",
								},
							},
							Subject: "Test Schedule Email",
							BodyStructure: EmailBodyPart{
								PartID: "1",
								Type:   "text/plain",
								Headers: []EmailHeader{
									{"Content-Language", "en"},
								},
							},
							BodyValues: map[string]EmailBodyValue{
								"1": {
									Value: "Test creating email with structs!!!",
								},
							},
						},
					},
				},
			},
		},
	})
	resp, err := PerformJMAPCall(req)
	if err != nil {
		log.Fatal(err)
	}
	PrettyPrintJSONResponse(resp)
}
