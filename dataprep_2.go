package main

import (
	"github.com/google/uuid"
	"log"
	"time"
)

type inputData struct {
	id        string
	timestamp int64
}

func prepareData(ic <-chan uuid.UUID) <-chan inputData {
	oc := make(chan inputData)
	go func() {
		for id := range ic {
			input := inputData{
				id:        id.String(),
				timestamp: time.Now().UnixNano(),
			}
			log.Printf("Data ready for processing:%v", input)
			oc <- input
		}
	}()
	return oc
}
