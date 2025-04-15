package models

import "time"

// type ExternalTransaction struct {
// 	ID        uint `gorm:"primaryKey"`
// 	Reference string
// 	Amount    float64
// 	Status    string
// }

type ExternalTransaction struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Amount    float64   `json:"amount"`
	Date      time.Time `json:"date"`
	Source    string    `json:"source"` // "internal" / "external"
	Reference string    `json:"reference"`
	Note      string    `json:"note"`
	Status    string    `json:"status"` // "SUCCESS" / "FAILED" / "PENDING"
	CreatedAt time.Time
}
