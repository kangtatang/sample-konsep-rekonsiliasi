package models

import "time"

type InternalTransaction struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	TransactionID string    `gorm:"uniqueIndex" json:"transaction_id"`
	AccountNumber string    `json:"account_number"`
	Amount        float64   `json:"amount"`
	Timestamp     time.Time `json:"timestamp"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Date          time.Time `json:"date"`
	Source        string    `json:"source"` // "internal" / "external"
	Note          string    `json:"note"`
}
