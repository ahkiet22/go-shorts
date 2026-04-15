package queue

import (
	"go-shorts/internal/repository"
	"log"
)

type ClickQueue struct {
	ch chan string
}

func NewClickQueue(r *repository.UrlRepository) *ClickQueue {
	q := &ClickQueue{
		ch: make(chan string, 1000),
	}

	go func() {
		for code := range q.ch {
			if err := r.IncreaseClick(code); err != nil {
				// log error
				log.Println("increase click error:", err)
			}
		}
	}()

	return q
}
