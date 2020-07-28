// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type AgoraRole string

const (
	AgoraRolePublisher  AgoraRole = "Publisher"
	AgoraRoleSubscriber AgoraRole = "Subscriber"
)

var AllAgoraRole = []AgoraRole{
	AgoraRolePublisher,
	AgoraRoleSubscriber,
}

func (e AgoraRole) IsValid() bool {
	switch e {
	case AgoraRolePublisher, AgoraRoleSubscriber:
		return true
	}
	return false
}

func (e AgoraRole) String() string {
	return string(e)
}

func (e *AgoraRole) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AgoraRole(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AgoraRole", str)
	}
	return nil
}

func (e AgoraRole) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
