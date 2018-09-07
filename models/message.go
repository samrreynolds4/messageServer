package models

import "time"

type Message struct {
	Message string    `json:"message"`
	Date    time.Time `json:"date"`
	User    string    `json:"user"`
}
