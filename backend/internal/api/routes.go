package api

import (
	"github.com/1704mori/certguardian/internal/db"
	"github.com/1704mori/certguardian/internal/repository"
	"github.com/gin-gonic/gin"
)

func setupRoutes(router *gin.RouterGroup, db *db.Database) {
	certRepo := repository.NewCertificateRepository(db)
	h := NewHandler(db, certRepo)

	router.GET("/", h.listDomain)
	router.GET("/:domain", h.findDomain)
	router.POST("/", h.addDomain)
}
