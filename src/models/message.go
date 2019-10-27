package models

import "time"

type (
	Message struct {
		ID        int       `json:"id"`
		Sender    string    `json:"from"`
		Receiver  string    `json:"to"`
		Text      string    `json:"text"`
		CreatedAt time.Time `json:"date_time"`
	}
)
