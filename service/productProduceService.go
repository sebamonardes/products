package service

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"strconv"
)

const (
	topic           = "first_topic"
	bootstrapServer = "localhost:9092"
)

func Produce(ctx context.Context) {
	w := kafka.Writer{
		Addr:         kafka.TCP(bootstrapServer),
		Topic:        topic,
		RequiredAcks: kafka.RequireAll,
	}

	for i := 0; i < 10; i++ {
		err := w.WriteMessages(ctx, kafka.Message{
			Key:   []byte(strconv.Itoa(i)),
			Value: []byte("Go Message " + strconv.Itoa(i)),
		})

		if err != nil {
			panic("Error al producir mensaje " + err.Error())
		} else {
			fmt.Println("Mensaje escrito ", i)
		}
	}

}
