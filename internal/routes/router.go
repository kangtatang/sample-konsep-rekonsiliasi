package routes

import (
	handlers "go-big-internal/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.GET("/generate/internal", handlers.GenerateDummyInternalData(db))
	r.POST("/seed", func(c *gin.Context) {
		handlers.GenerateInternalDummy(c)
	})
	r.GET("/reconcile", handlers.ReconcileData(db))
	r.GET("/export", handlers.ExportReconciliationExcel(db))
	r.GET("/reconciliation", handlers.GetReconciliationResults(db))

	return r
}
