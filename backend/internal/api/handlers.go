package api

import (
	"net/http"

	"github.com/1704mori/certguardian/internal/certificate"
	"github.com/1704mori/certguardian/internal/db"
	"github.com/1704mori/certguardian/internal/repository"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	DB   *db.Database
	repo *repository.CertificateRepository
}

func NewHandler(
	db *db.Database,
	repo *repository.CertificateRepository,
) *Handler {
	return &Handler{
		DB:   db,
		repo: repo,
	}
}

func (h *Handler) addDomain(c *gin.Context) {
	var domain certificate.Metadata

	if err := c.ShouldBindJSON(&domain); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if d, _ := h.repo.Find(domain.CommonName); d != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Domain already exists"})
		return
	}

	err := h.repo.Add(domain.CommonName, certificate.Metadata{
		CommonName: domain.CommonName,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Domain added successfully"})
}

func (h *Handler) findDomain(c *gin.Context) {
	domain := c.Param("domain")

	if domain == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid domain"})
		return
	}

	d, err := h.repo.Find(domain)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, d)
}

func (h *Handler) listDomain(c *gin.Context) {
	c.JSON(200, gin.H{"message": "todo"})
}
