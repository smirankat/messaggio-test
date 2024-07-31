package main

import (
	"context"
	"fmt"
	// "net/http"

	"github.com/segmentio/kafka-go"
)

// func handlerKafkaMessage(w http.ResponseWriter, r *http.Request) {

func StartKafka() {
	conf := kafka.ReaderConfig{
		Brokers: []string{"localhost: 9092"},
		Topic:   "mytopic",
		GroupID: "g1",
	}

	reader := kafka.NewReader(conf)

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Println("Some error occured", err)
			continue
		}
		fmt.Println("Message is : ", string(m.Value))
	}

}

// 	respondWithJSON(w, 200, struct{}{})
// }
