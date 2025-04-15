package models

import "time"

type ReconciliationResult struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	TransactionID  string    `json:"transaction_id"`
	Status         string    `json:"status"` // "matched", "missing_internal", "missing_external", etc
	InternalAmount float64   `json:"internal_amount"`
	ExternalAmount float64   `json:"external_amount"`
	ReconciledAt   time.Time `json:"reconciled_at"`
}
