package cron

import (
	"fmt"
	"time"

	"github.com/1704mori/certguardian/internal/db"
	"github.com/1704mori/certguardian/internal/repository"
	"github.com/1704mori/certguardian/internal/sslcheck"
	"github.com/go-co-op/gocron/v2"
	"github.com/rs/zerolog/log"
)

type repositories struct {
	domain *repository.DomainRepository
	cert   *repository.CertificateRepository
}

type Cron struct {
	scheduler gocron.Scheduler
	repos     *repositories
}

func NewCron(db *db.Database) *Cron {
	s, err := gocron.NewScheduler()
	if err != nil {
		panic(err)
	}

	certRepo := repository.NewCertificateRepository(db)
	domainRepo := repository.NewDomainRepository(db)

	return &Cron{
		scheduler: s,
		repos: &repositories{
			cert:   certRepo,
			domain: domainRepo,
		},
	}
}

func (c *Cron) CertificateJob(every time.Duration) {
	j, err := c.scheduler.NewJob(
		gocron.DurationJob(
			every,
		),
		gocron.NewTask(
			func() {
				c.checkForCertificates()
			},
		),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(j.ID())
}

func (c *Cron) Start() {
	c.scheduler.Start()
}

func (c *Cron) checkForCertificates() {
	log.Info().Msg("[cron]: trying to update certificates")
	certs, err := c.repos.cert.List()

	if len(certs[0].Directories) == 0 {
		log.Info().Msg("[cron]: no certificates found to update")
		return
	}

	if err != nil {
		panic(err)
	}

	for dir := range certs[0].Directories {
		certFile, err := sslcheck.SearchDirectoryForCertificates(dir)
		if err != nil {
			return
		}

		for _, cert := range certFile {
			certData, err := sslcheck.FromPEM(cert)
			if err != nil {
				log.Error().Msgf("[cron]: failed to get cert data for %s", cert)
			}

			certs[0].Directories[dir][cert] = *certData
		}
	}

	err = c.repos.cert.Add(certs[0])
	if err != nil {
		log.Error().Msg("[cron]: failed to update certificates")
	}

	log.Info().Msg("[cron]: certificates updated")
}
