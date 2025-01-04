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
}
