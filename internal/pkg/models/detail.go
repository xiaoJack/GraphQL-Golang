package models

import "time"

type Detail struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Price       float64   `json:"price"`
	CreatedTime time.Time `json:"created_time"`
}
