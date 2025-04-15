package handler

import (
	"go-big-internal/config"
	"go-big-internal/sampledata"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GenerateInternalDummy(c *gin.Context) {
	db := config.InitDB()

	data := sampledata.GenerateSampleTransactions(100, "internal")

	for _, tx := range data {
		db.Create(&tx)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Generated internal dummy data", "count": len(data)})
}
