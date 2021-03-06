package avro

//This is my favourite job: writing error-prone stupid mappings manually :P

import (
	"errors"
	"fmt"
	"github.com/maxromanovsky/game_scraper/domain/entity"
	"time"
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
	msg.From = ems.From
	msg.To = ems.To
	msg.Subject = ems.Subject
	msg.Raw = ems.Raw
	msg.DateReceived = time.Unix(ems.DateReceived, 0)

	switch ems.Source {
	case EmailSourceGOG:
		msg.Source = entity.GOG
	case EmailSourceAppleAppStore:
		msg.Source = entity.AppleAppStore
	default:
		return nil, errors.New(fmt.Sprintf("Unsupported message source: %s", ems.Source.String()))
	}

	for _, p := range ems.Parts {
		part := entity.BodyPart{PartId: p.PartId, MimeType: p.MimeType, Filename: p.Filename, Body: p.Body}
		part.Headers = p.Headers
		msg.Parts = append(msg.Parts, part)
	}

	return &msg, nil
}
