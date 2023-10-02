package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/brianvoe/gofakeit"
)

const eventsCount = 100_000

func main() {
	ammos := generateAmmos(eventsCount)

	wg := sync.WaitGroup{}
	sem := make(chan struct{}, 1000)
	for _, r := range ammos {
		r := r
		wg.Add(1)
		go func() {
			sem <- struct{}{}
			defer func() {
				wg.Done()
				<-sem
			}()
			_, err := http.Post("http://localhost:8080/api/v1/event", "application/json", r)
			if err != nil {
				fmt.Println(err)
			}
		}()
	}
	wg.Wait()
}

type Event struct {
	EventType string `json:"event_type"`
	Datetime  string `json:"datetime"`
	Payload   string `json:"payload"`
}

type payload struct {
	To   string `json:"to"`
	Text string `json:"text"`
}

func generateAmmos(size int) []io.Reader {
	res := make([]io.Reader, 0, size)
	for i := 0; i < size; i++ {
		pld := payload{
			To:   gofakeit.Email(),
			Text: gofakeit.BuzzWord(),
		}
		p, _ := json.Marshal(pld)
		evt := Event{
			EventType: "SEND_NOTIFICATION",
			Datetime:  "2023-10-02T08:26:10Z",
			Payload:   string(p),
		}
		buf, _ := json.Marshal(evt)
		res = append(res, bytes.NewReader(buf))
	}
	return res
}
