package entity

import "time"

//go:generate $GOPATH/bin/gogen-avro -containers ../repository/avro ../../assets/avro/email-message.json

type EmailMessage struct {
	Id, From, To, Subject, Raw string
	Source                     EmailSource
	DateReceived               time.Time
	Parts                      []BodyPart
}

type BodyPart struct {
	PartId, MimeType, Filename, Body string
	Headers                          map[string]string //Assumption: no headers with duplicate names
}

type EmailSource int

const (
	GOG EmailSource = iota
	AppleAppStore
)
