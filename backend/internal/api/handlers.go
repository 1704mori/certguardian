package api

import (
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
// func (h *Handler) addCertificate(c *gin.Context) {
// 	var certs certificates.Metadata

// 	if err := c.ShouldBindJSON(&certs); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	existingDirs, err := h.repos.cert.FindByDirectories(certs.Directories)

// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if len(existingDirs.Directories) > 0 {
// 		// get the dir that is in existingDirs and in certs
// 		var dirs []string
// 		for _, dir := range certs.Directories {
// 			for _, existingDir := range existingDirs.Directories {
// 				if dir == existingDir {
// 					dirs = append(dirs, dir)
// 				}
// 			}
// 		}

// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Directories already exists", "directories": dirs})
// 		return
// 	}

// 	err = h.repos.cert.Add(certs.Directories)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Directories added successfully"})
// }
