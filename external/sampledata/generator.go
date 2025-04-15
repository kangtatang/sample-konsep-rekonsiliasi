package sampledata

import (
	"fmt"
	"math/rand"
	"time"

	"go-big-external/models"

	"github.com/google/uuid"
)

func GenerateSampleTransactions(n int, source string) []models.ExternalTransaction {
	rand.Seed(time.Now().UnixNano())
	var transactions []models.ExternalTransaction

	for i := 0; i < n; i++ {
		id := uuid.New().String()
		amount := float64(rand.Intn(10000)) / 100.0
		date := time.Now().AddDate(0, 0, -rand.Intn(10)) // random 10 hari ke belakang

		tx := models.ExternalTransaction{
			ID:        id,
			Amount:    amount,
			Date:      date,
			Note:      fmt.Sprintf("Sample Transaction #%d", i+1),
			Status:    []string{"SUCCESS", "FAILED", "PENDING"}[rand.Intn(3)],
			Reference: fmt.Sprintf("EXT-%s", id[:8]), // ambil 8 karakter pertama dari UUID
			CreatedAt: time.Now(),
			Source:    source,
		}

		transactions = append(transactions, tx)
	}

	return transactions
}
