package main

import (
	"encoding/json"

	"github.com/foxcpp/go-jmap"
)

type MailboxQueryResponse struct {
	AccountID string   `json:"accountId"`
	IDs       []string `json:"ids"`
	Total     int      `json:"total"`
}

func GenericUnmarshaler[T any]() func(args json.RawMessage) (any, error) {
	return func(args json.RawMessage) (any, error) {
		var result T
		err := json.Unmarshal(args, &result)
		return &result, err
	}
}

var Unmarshallers = map[string]jmap.FuncArgsUnmarshal{
	"Mailbox/query": GenericUnmarshaler[MailboxQueryResponse](),
	"Mailbox/get":   GenericUnmarshaler[any](),
	"Email/set":     GenericUnmarshaler[any](),
}

type Address struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email"`
}

type EmailBodyValue struct {
	Value       string `json:"value"`
	IsTruncated bool   `json:"isTruncated"`
}

type EmailBodyPart struct {
	PartID  string        `json:"partId,omitempty"`
	BlobID  string        `json:"blobId,omitempty"`
	Headers []EmailHeader `json:"header,omitempty"`
	Type    string        `json:"type,omitempty"`
}

type EmailHeader struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type MessageParams struct {
	MailboxIDs    map[string]bool           `json:"mailboxIds"`
	Keywords      map[string]bool           `json:"keywords"`
	From          []Address                 `json:"from"`
	To            []Address                 `json:"to"`
	Subject       string                    `json:"subject"`
	BodyValues    map[string]EmailBodyValue `json:"bodyValues"`
	BodyStructure EmailBodyPart             `json:"bodyStructure"`
}
