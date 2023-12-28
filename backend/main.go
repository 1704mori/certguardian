package main

import (
	"fmt"
	"os/exec"
	"time"

	env "github.com/1704mori/certguardian/internal"
	"github.com/1704mori/certguardian/internal/api"
	"github.com/1704mori/certguardian/internal/cron"
	"github.com/1704mori/certguardian/internal/db"
	"github.com/1704mori/certguardian/internal/util"
	"github.com/1704mori/certguardian/internal/version"

	"github.com/alexflint/go-arg"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	arg.MustParse(&env.Args)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Info().Msgf("%v", env.Args)

	database, err := db.New("certguard.db")
	if err != nil {
		log.Fatal().Msgf("Failed to initialize database: %v", err)
	}
	defer database.Close()

	log.Info().Msgf("Starting certguardian version %s", version.Version)

	env.BuildEnv()
	cmd := exec.Command("npm", "run", "build")
	cmd.Dir = "/build/frontend"

	stdout, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}

	fmt.Println(string(stdout))

	srv := api.NewServer(database)

	interval, err := util.ConvertToDuration(env.Args.CronInterval)
	if err != nil {
		log.Error().Msgf("[env] could not convert CRON_INTERVAL with value %s using default value", env.Args.CronInterval)
		interval = 24 * time.Hour
	}
	log.Info().Msgf("%v", interval)

	c := cron.NewCron(database)
	c.UpdateCertsAndDomains(interval)
	c.Start()

	log.Info().Msgf("Listening on port %v", env.Args.Port)
	srv.Run(fmt.Sprintf(":%s", env.Args.Port))
}
