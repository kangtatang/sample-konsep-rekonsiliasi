package handler

import (
	"go-big-external/config"
	"go-big-external/sampledata"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GenerateExternalDummy(c *gin.Context) {
	db := config.InitDB()

	data := sampledata.GenerateSampleTransactions(100, "external")

	for _, tx := range data {
		db.Create(&tx)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Generated external dummy data", "count": len(data)})
}
