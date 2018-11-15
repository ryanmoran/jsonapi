package jsonapi

import (
	"fmt"
	"strconv"
)

type Error struct {
	ID     string `json:"id,omitempty"`
	Status Status `json:"status,omitempty"`
	Code   string `json:"code,omitempty"`
	Title  string `json:"title,omitempty"`
	Detail string `json:"detail,omitempty"`
}

type Errors []Error

type Status int

func (s Status) MarshalJSON() ([]byte, error) {
	quotedStatus := fmt.Sprintf("%q", strconv.Itoa(int(s)))
	return []byte(quotedStatus), nil
}
