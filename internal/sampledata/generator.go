package sampledata

import (
	"fmt"
	"math/rand"
	"time"

	"go-big-internal/models"

	"github.com/google/uuid"
)

func GenerateSampleTransactions(n int, source string) []models.InternalTransaction {
	rand.Seed(time.Now().UnixNano())
	var transactions []models.InternalTransaction

	for i := 0; i < n; i++ {
		id := uuid.New().String()
		amount := float64(rand.Intn(10000)) / 100.0
		date := time.Now().AddDate(0, 0, -rand.Intn(10)) // random 10 hari ke belakang

		tx := models.InternalTransaction{
			TransactionID: id,
			AccountNumber: fmt.Sprintf("ACC%04d", rand.Intn(10000)),
			Amount:        amount,
			Date:          date,
			Note:          fmt.Sprintf("Sample Transaction #%d", i+1),
			Source:        source,
		}

		transactions = append(transactions, tx)
	}

	return transactions
}
