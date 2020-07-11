package avro

//This is my favourite job: writing error-prone stupid mappings manually :P

import (
	"errors"
	"fmt"
	"github.com/maxromanovsky/game_scraper/domain/entity"
)

func ToEmailMessageSchema(message *entity.EmailMessage) (*EmailMessageSchema, error) {
	ems := NewEmailMessageSchema()
	ems.Id = message.Id
	ems.From = message.From
	ems.To = message.To
	ems.Subject = message.Subject
	ems.Raw = message.Raw
	ems.DateReceived = message.DateReceived.Unix()

	switch message.Source {
	case entity.GOG:
		ems.Source = EmailSourceGOG
	case entity.AppleAppStore:
		ems.Source = EmailSourceAppleAppStore
	default:
		return nil, errors.New(fmt.Sprintf("Unsupported message source: %v", message.Source))
	}

	for _, p := range message.Parts {
		part := NewEmailBodyPart()
		part.PartId = p.PartId
		part.MimeType = p.MimeType
		part.Filename = p.Filename
		part.Body = p.Body
		part.Headers = p.Headers

		ems.Parts = append(ems.Parts, part)
	}

	return ems, nil
}

func FromEmailMessageSchema(ems *EmailMessageSchema) (*entity.EmailMessage, error) {
	msg := entity.EmailMessage{}
	msg.Id = ems.Id
	msg.Subject = ems.Subject
	return &msg, nil
}
