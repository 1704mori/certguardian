package api

import (
	"github.com/1704mori/certguardian/internal/db"
	"github.com/gin-gonic/gin"
)

func setupRoutes(router *gin.RouterGroup, db *db.Database) {
	h := NewHandler(db)

	router.GET("/", h.listDomain)
}
