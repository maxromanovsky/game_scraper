package main

import (
	"github.com/maxromanovsky/game_scraper/domain/repository/avro"
	"io"
	"log"
	"os"
)

func main() {
	messages := make(chan *avro.EmailMessageSchema)
	done := make(chan struct{})

	go func() {
		defer close(done)

		i := 1
		for m := range messages {
			log.Printf("%d -> %s, %s", i, m.Id, m.Subject)
			i++
		}
	}()

	readAvro(messages)

	<-done
}

func readAvro(messages chan *avro.EmailMessageSchema) {
	fileReader, err := os.Open("email_messages.avro")
	if err != nil {
		log.Fatalf("Error opening file reader: %v", err)
	}

	// Create a new OCF reader
	reader, err := avro.NewEmailMessageSchemaReader(fileReader) //todo: configurable via CLI
	if err != nil {
		log.Fatalf("Error creating OCF file reader: %v\n", err)
	}

	// Read the records back until the file is finished
	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				log.Println("Avro: EOF")
				close(messages)
				return
			}
			log.Fatalf("Error reading OCF file: %v", err)
		}

		messages <- record
	}
}
