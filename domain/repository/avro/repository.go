package avro

import (
	"github.com/actgardner/gogen-avro/v7/container"
	"github.com/maxromanovsky/game_scraper/domain/entity"
	"io"
	"log"
	"os"
)

type EmailMessageRepository struct {
	fileName string
}

func NewEmailMessageRepository(fileName string) *EmailMessageRepository {
	return &EmailMessageRepository{fileName}
}

func (r *EmailMessageRepository) Save(codec container.Codec, recordsPerBlock int64, messages <-chan *entity.EmailMessage, callback func()) {
	fileWriter, err := os.Create(r.fileName)
	if err != nil {
		log.Fatalf("Error opening file writer: %v", err)
	}

	containerWriter, err := NewEmailMessageSchemaWriter(fileWriter, codec, recordsPerBlock)
	if err != nil {
		log.Fatalf("Error opening container writer: %v", err)
	}

	for m := range messages {
		record, err := ToEmailMessageSchema(m)
		if err != nil {
			log.Fatalf("Error converting message to record: %s", m.Id)
		}

		err = containerWriter.WriteRecord(record)
		if err != nil {
			log.Fatalf("Error writing record to file: %v", err)
		}
	}

	err = containerWriter.Flush()
	if err != nil {
		log.Fatalf("Error flushing last block to file: %v", err)
	}

	err = fileWriter.Close()
	if err != nil {
		log.Fatalf("Error closing file: %v", err)
	}

	callback()
}

func (r *EmailMessageRepository) Load(messages chan<- *entity.EmailMessage, callback func()) {
	fileReader, err := os.Open(r.fileName)
	if err != nil {
		log.Fatalf("Error opening file reader: %v", err)
	}

	// Create a new OCF reader
	reader, err := NewEmailMessageSchemaReader(fileReader)
	if err != nil {
		log.Fatalf("Error creating file reader: %v", err)
	}

	// Read the records back until the file is finished
	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				log.Println("Avro: EOF")
				err = fileReader.Close()
				if err != nil {
					log.Fatalf("Error closing file: %v", err)
				}
				close(messages)
				if callback != nil {
					callback()
				}
				return
			}
			log.Fatalf("Error reading file: %v", err)
		}

		message, err := FromEmailMessageSchema(record)
		if err != nil {
			log.Fatalf("Error converting record to message: %s", record.Id)
		}
		messages <- message
	}
}
