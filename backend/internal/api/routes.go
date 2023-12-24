package api

import (
	"github.com/1704mori/certguardian/internal/db"
	"github.com/1704mori/certguardian/internal/repository"
	"github.com/gin-gonic/gin"
)

func setupRoutes(router *gin.RouterGroup, db *db.Database) {
	domainRepo := repository.NewDomainRepository(db)
	certRepo := repository.NewCertificateRepository(db)
	h := NewHandler(db, &repositories{
		domain: domainRepo,
		cert:   certRepo,
	})

	domain := router.Group("/domain")
	{
		domain.GET("/", h.listDomain)
		domain.GET("/:domain", h.findDomain)
		domain.POST("/", h.addDomain)
		domain.DELETE("/:domain", h.deleteDomain)
	}

	// cert := router.Group("/cert")
	//
	//	{
	//		cert.POST("/", h.addCertificate)
	//	}
}
