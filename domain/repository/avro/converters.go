package avro

import "github.com/maxromanovsky/game_scraper/domain/entity"

func ToEmailMessageSchema(message *entity.EmailMessage) *EmailMessageSchema {
	ems := NewEmailMessageSchema()
	ems.Id = message.Id
	ems.Subject = message.Subject
	return ems
}

func FromEmailMessageSchema(ems *EmailMessageSchema) *entity.EmailMessage {
	msg := entity.EmailMessage{}
	msg.Id = ems.Id
	msg.Subject = ems.Subject
	return &msg
}
