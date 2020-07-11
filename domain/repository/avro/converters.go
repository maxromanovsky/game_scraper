package avro

import "github.com/maxromanovsky/game_scraper/domain/entity"

func ToEmailMessageSchema(message *entity.EmailMessage) *EmailMessageSchema {
	ems := NewEmailMessageSchema()
	ems.Id = message.Id
	ems.Subject = message.Subject
	return ems
}
