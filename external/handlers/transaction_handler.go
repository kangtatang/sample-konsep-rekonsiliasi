package handler

import (
	"go-big-external/models"
	"math/rand"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SeedExternalData(c *gin.Context, db *gorm.DB) {
	for i := 0; i < 5000; i++ {
		tx := models.ExternalTransaction{
			Reference: "EXT-" + strconv.Itoa(100000+i),
			Source:    "external",
			Amount:    float64(rand.Intn(1000000)) / 100,
			Date:      time.Now().AddDate(0, 0, -rand.Intn(30)), // random date within the last 30 days
			Note:      "Dummy transaction " + strconv.Itoa(i),
			Status:    []string{"SUCCESS", "FAILED", "PENDING"}[rand.Intn(3)],
		}
		db.Create(&tx)
	}
	c.JSON(200, gin.H{"message": "Data eksternal berhasil di-generate"})
}

func GetExternalTransactions(c *gin.Context, db *gorm.DB) {
	var transactions []models.ExternalTransaction
	if err := db.Find(&transactions).Error; err != nil {
		c.JSON(500, gin.H{"error": "Gagal mengambil data"})
		return
	}
	c.JSON(200, transactions)
}
