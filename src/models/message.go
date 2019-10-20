package models

import "time"

type (
	Message struct {
		ID        int       `json:"id"`
		From      string    `json:"from"`
		To        string    `json:"to"`
		Text      string    `json:"text"`
		CreatedAt time.Time `json:"date_time"`
	}
)
