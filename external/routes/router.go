package routes

import (
	handlers "go-big-external/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.POST("/api/external/seed", func(c *gin.Context) {
		handlers.SeedExternalData(c, db)
	})

	r.POST("/api/external/new-seed", func(c *gin.Context) {
		handlers.GenerateExternalDummy(c)
	})

	r.GET("/api/external/transactions", func(c *gin.Context) {
		handlers.GetExternalTransactions(c, db)
	})

	return r
}
