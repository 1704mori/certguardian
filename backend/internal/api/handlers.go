package api

import (
	"fmt"
	"net/http"

	"github.com/1704mori/certguardian/internal/db"
	domainTypings "github.com/1704mori/certguardian/internal/domain"
	"github.com/1704mori/certguardian/internal/repository"
	"github.com/1704mori/certguardian/internal/sslcheck"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	DB    *db.Database
	repos *repositories
}

type repositories struct {
	domain *repository.DomainRepository
	cert   *repository.CertificateRepository
}

func NewHandler(
	db *db.Database,
	repos *repositories,
) *Handler {
	return &Handler{
		DB:    db,
		repos: repos,
	}
}

func (h *Handler) addDomain(c *gin.Context) {
	var domain domainTypings.Info

	if err := c.ShouldBindJSON(&domain); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := h.repos.domain.Find(domain.CommonName); err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Domain already exists"})
		return
	}

	certInfo, err := sslcheck.FromDomain(domain.CommonName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldn't get certificate information"})
		return
	}

	err = h.repos.domain.Add(domain.CommonName, domainTypings.Info{
		CommonName: domain.CommonName,
		ValidTo:    certInfo.ValidTo,
		ValidFrom:  certInfo.ValidFrom,
		IsExpired:  certInfo.IsExpired,
		Issuer:     certInfo.Issuer,
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

	d, err := h.repos.domain.Find(domain)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, d)
}

func (h *Handler) listDomain(c *gin.Context) {
	d, err := h.repos.domain.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": d})
}

func (h *Handler) deleteDomain(c *gin.Context) {
	domain := c.Param("domain")

	if domain == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid domain"})
		return
	}

	err := h.repos.domain.Delete(domain)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Domain deleted"})
}

// certificates
func (h *Handler) addCertificate(c *gin.Context) {
	var certs struct {
		Directories []string `json:"directories" binding:"required"`
	}

	if err := c.ShouldBindJSON(&certs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dirCertificates, err := sslcheck.FindCertificates(certs.Directories)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(dirCertificates)

	err = h.repos.cert.Add(dirCertificates)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Directories added successfully"})
}

func (h *Handler) listCertificates(c *gin.Context) {
	d, err := h.repos.cert.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": d[0].Directories})
}
