package main

import (
	"fmt"
	"time"

	env "github.com/1704mori/certguardian/internal"
	"github.com/1704mori/certguardian/internal/api"
	"github.com/1704mori/certguardian/internal/cron"
	"github.com/1704mori/certguardian/internal/db"
	"github.com/1704mori/certguardian/internal/version"

	"github.com/alexflint/go-arg"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	arg.MustParse(&env.Args)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	database, err := db.New("data/certguard.db")
	if err != nil {
		log.Fatal().Msgf("Failed to initialize database: %v", err)
	}
	defer database.Close()

	log.Info().Msgf("Starting certguardian version %s", version.Version)

	srv := api.NewServer(database)

	c := cron.NewCron(database)
	c.CertificateJob(10 * time.Second)
	c.Start()

	log.Info().Msgf("Listening on port %v", env.Args.Port)
	srv.Run(fmt.Sprintf(":%s", env.Args.Port))
}
