package main

import (
	"encoding/json"
	"time"

	"github.com/foxcpp/go-jmap"
)

type MailboxQueryResponse struct {
	AccountID string   `json:"accountId"`
	IDs       []string `json:"ids"`
	Total     int      `json:"total"`
}

type EmailSetResponse struct {
	AccountID string           `json:"accountId"`
	Created   map[string]Email `json:"created"`
}

func GenericUnmarshaler[T any]() func(args json.RawMessage) (any, error) {
	return func(args json.RawMessage) (any, error) {
		var result T
		err := json.Unmarshal(args, &result)
		return &result, err
	}
}

var Unmarshallers = map[string]jmap.FuncArgsUnmarshal{
	"Mailbox/query":       GenericUnmarshaler[MailboxQueryResponse](),
	"Mailbox/get":         GenericUnmarshaler[any](),
	"Identity/get":        GenericUnmarshaler[IdentityGetResponse](),
	"Email/set":           GenericUnmarshaler[EmailSetResponse](),
	"EmailSubmission/set": GenericUnmarshaler[any](),
}

type Address struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email"`
}

type Email struct {
	BlobID   string `json:"blobId,omitempty"`
	ID       string `json:"id,omitempty"`
	Size     int    `json:"size"`
	ThreadID string `json:"threadId,omitempty"`
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

type Envelope struct {
	MailFrom      Address    `json:"mailFrom"`
	RecipientTo   []Address  `json:"rcptTo"`
	FutureRelease *time.Time `json:"FUTURERELEASE,omitempty"`
}

type EmailSubmission struct {
	EmailID            string            `json:"emailId"`
	IdentityIDResultOf map[string]string `json:"#identityId,omitempty"`
	IdentityID         string            `json:"identityId,omitempty"`
	Envelope           Envelope          `json:"envelope"`
}

type IdentityGetResponse struct {
	AccountID string     `json:"accountId"`
	List      []Identity `json:"list"`
}

type Identity struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}
