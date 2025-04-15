package handler

import (
	"encoding/json"
	"go-big-internal/models"
	"io"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"go-big-internal/utils"
)

// type ExternalTransaction struct {
// 	TransactionID string    `json:"transaction_id"`
// 	AccountNumber string    `json:"account_number"`
// 	Amount        float64   `json:"amount"`
// 	Timestamp     time.Time `json:"timestamp"`
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

func ReconcileData(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		externalData, err := fetchExternalData()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		results := reconcileTransactions(db, externalData)

		saveReconciliationResults(db, results)

		c.JSON(http.StatusOK, gin.H{"message": "Reconciliation completed", "total": len(results)})
	}
}

func fetchExternalData() ([]ExternalTransaction, error) {
	resp, err := http.Get("http://localhost:8081/api/external/transactions")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var externalData []ExternalTransaction
	if err := json.Unmarshal(body, &externalData); err != nil {
		return nil, err
	}

	return externalData, nil
}

func reconcileTransactions(db *gorm.DB, externalData []ExternalTransaction) []models.ReconciliationResult {
	var results []models.ReconciliationResult

	for _, ext := range externalData {
		var internal models.InternalTransaction
		tx := db.Where("transaction_id = ?", ext.ID).First(&internal)

		if tx.Error != nil && tx.Error == gorm.ErrRecordNotFound {
			results = append(results, models.ReconciliationResult{
				TransactionID:  ext.ID,
				Status:         "missing_internal",
				ExternalAmount: ext.Amount,
				ReconciledAt:   time.Now(),
			})
		} else {
			status := "matched"
			if internal.Amount != ext.Amount {
				status = "amount_mismatch"
			}
			results = append(results, models.ReconciliationResult{
				TransactionID:  ext.ID,
				Status:         status,
				InternalAmount: internal.Amount,
				ExternalAmount: ext.Amount,
				ReconciledAt:   time.Now(),
			})
		}
	}

	return results
}

func saveReconciliationResults(db *gorm.DB, results []models.ReconciliationResult) {
	for _, result := range results {
		db.Create(&result)
	}
}

func GenerateDummyInternalData(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		for i := 0; i < 1000; i++ {
			tx := models.InternalTransaction{
				TransactionID: generateTransactionID(i),
				AccountNumber: generateAccountNumber(),
				Amount:        float64(1000 + rand.Intn(9000)),
				Timestamp:     time.Now().Add(-time.Duration(rand.Intn(72)) * time.Hour),
			}
			db.Create(&tx)
		}
		c.JSON(http.StatusOK, gin.H{"message": "Dummy internal transactions generated"})
	}
}

func generateTransactionID(i int) string {
	return time.Now().Format("20060102150405") + "_" + string(rune('A'+i%26))
}

func generateAccountNumber() string {
	return "ACC" + time.Now().Format("150405") + string(rune(rand.Intn(100)+65))
}

func ExportReconciliationExcel(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		startDateStr := c.Query("start_date") // format: 2025-04-01
		endDateStr := c.Query("end_date")     // format: 2025-04-10

		var results []models.ReconciliationResult
		query := db.Model(&models.ReconciliationResult{})

		// Filter by date range
		if startDateStr != "" && endDateStr != "" {
			start, err1 := time.Parse("2006-01-02", startDateStr)
			end, err2 := time.Parse("2006-01-02", endDateStr)
			if err1 == nil && err2 == nil {
				end = end.Add(24 * time.Hour) // Tambah 1 hari agar termasuk hari terakhir
				query = query.Where("reconciled_at BETWEEN ? AND ?", start, end)
			}
		}

		query.Order("reconciled_at desc").Find(&results)

		excelFile, err := utils.ExportToExcel(results)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate excel"})
			return
		}

		c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
		c.Header("Content-Disposition", "attachment; filename=reconciliation.xlsx")
		c.Header("Content-Transfer-Encoding", "binary")
		_ = excelFile.Write(c.Writer)
	}
}

func GetReconciliationResults(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		startDateStr := c.Query("start_date")
		endDateStr := c.Query("end_date")

		var results []models.ReconciliationResult
		query := db.Model(&models.ReconciliationResult{})

		if startDateStr != "" && endDateStr != "" {
			start, err1 := time.Parse("2025-01-02", startDateStr)
			end, err2 := time.Parse("2025-05-02", endDateStr)
			if err1 == nil && err2 == nil {
				end = end.Add(24 * time.Hour)
				query = query.Where("reconciled_at BETWEEN ? AND ?", start, end)
			}
		}

		query.Order("reconciled_at desc").Find(&results)

		c.JSON(http.StatusOK, results)
	}
}

// func ExportReconciliationExcel(db *gorm.DB) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var results []models.ReconciliationResult
// 		db.Find(&results)

// 		excelFile, err := utils.ExportToExcel(results)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate excel"})
// 			return
// 		}

// 		c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
// 		c.Header("Content-Disposition", "attachment; filename=reconciliation.xlsx")
// 		c.Header("Content-Transfer-Encoding", "binary")
// 		_ = excelFile.Write(c.Writer)
// 	}
// }
