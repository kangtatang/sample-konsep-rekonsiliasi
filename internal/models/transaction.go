package models

import "time"

type Transaction struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Amount    float64   `json:"amount"`
	Date      time.Time `json:"date"`
	Source    string    `json:"source"` // "internal" / "external"
	Note      string    `json:"note"`
	CreatedAt time.Time
}
