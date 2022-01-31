package service

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
)

func Consume() {

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{bootstrapServer},
		Topic:   topic,
		GroupID: "my-fourth-aplication",
	})

	err := r.SetOffset(42)
	if err != nil {
		return
	}

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			fmt.Println("Error al leer mensaje : ", err)
			break
		}
		fmt.Printf("Message %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader: ", err)
	}
}
