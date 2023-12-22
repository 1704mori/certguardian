package api

import (
	"github.com/1704mori/certguardian/internal/db"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	DB *db.Database
}

func NewHandler(db *db.Database) *Handler {
	return &Handler{DB: db}
}

func (h *Handler) addDomain(c *gin.Context) {
	c.JSON(200, gin.H{"message": "todo"})
}

func (h *Handler) listDomain(c *gin.Context) {
	c.JSON(200, gin.H{"message": "todo"})
}
