package models

import "time"

type Book struct {
	ID          int       `json:"id"`
	Code        string    `json:"code"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Author      string    `json:"author"`
	Publisher   string    `json:"publisher"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
