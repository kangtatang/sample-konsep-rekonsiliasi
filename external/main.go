// File: main.go
package main

import (
	"go-big-external/config"
	"go-big-external/routes"
	"log"
)

func main() {
	// Inisialisasi koneksi ke database
	db := config.InitDB()

	// Jalankan migrasi untuk buat tabel jika belum ada
	if err := config.Migrate(db); err != nil {
		log.Fatal("Migration failed:", err)
	}

	// Setup router dan inject db ke handler
	r := routes.SetupRouter(db)

	// Jalankan server
	if err := r.Run(":8081"); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}

// // File: main.go
// package main

// import (
// 	"fmt"
// 	"log"
// 	"math/rand"
// 	"net/http"
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// type ExternalTransaction struct {
// 	ID            uint      `gorm:"primaryKey" json:"id"`
// 	TransactionID string    `gorm:"uniqueIndex" json:"transaction_id"`
// 	Amount        float64   `json:"amount"`
// 	CreatedAt     time.Time `json:"created_at"`
// }

// var db *gorm.DB

// func initDB() {
// 	dsn := "host=localhost user=postgres password=postgres dbname=go_big_external port=5432 sslmode=disable TimeZone=Asia/Jakarta"
// 	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatal("Gagal konek DB:", err)
// 	}
// 	db = database
// 	db.AutoMigrate(&ExternalTransaction{})
// }

// func generateDummyTransactions(c *gin.Context) {
// 	total := 1000
// 	rand.Seed(time.Now().UnixNano())
// 	for i := 0; i < total; i++ {
// 		tx := ExternalTransaction{
// 			TransactionID: generateTransactionID(i),
// 			Amount:        float64(rand.Intn(100000)) / 100,
// 			CreatedAt:     time.Now().Add(-time.Duration(rand.Intn(72)) * time.Hour),
// 		}
// 		db.Create(&tx)
// 	}
// 	c.JSON(http.StatusOK, gin.H{"message": "Dummy data generated", "total": total})
// }

// func getTransactions(c *gin.Context) {
// 	var txs []ExternalTransaction
// 	db.Order("created_at asc").Limit(1000).Find(&txs)
// 	c.JSON(http.StatusOK, txs)
// }

// func generateTransactionID(i int) string {
// 	return fmt.Sprintf("EXT-TX-%06d", i+1)
// }

// func main() {
// 	initDB()
// 	r := gin.Default()

// 	r.GET("/api/external/transactions", getTransactions)
// 	r.POST("/api/external/generate", generateDummyTransactions)

// 	r.Run(":8081")
// }
