package env

import (
	"fmt"
	"os"
	"strings"
)

var Args struct {
	Port                string `arg:"env:PORT" default:"7070" help:"Port to listen on"`
	CronInterval        string `arg:"env:CRON_INTERVAL" default:"1d" help:"Interval for the cron job. e.g: 15d (15 days) or 3h (3 hours)"`
	NearExpiryThreshold string `arg:"env:NEAR_EXPIRY_THRESHOLD" default:"10d" help:"Threshold to determine if a certificate is near expiration. e.g: 15d (15 days) or 3h (3 hour"`
}

func BuildEnv() {
	addedVars := make(map[string]bool)

	var envContent strings.Builder

	addVar := func(key, value string) {
		if _, exists := addedVars[key]; !exists {
			addedVars[key] = true
			envContent.WriteString(fmt.Sprintf("%s=%s\n", key, value))
		}
	}

	addVar("PUBLIC_PORT", Args.Port)
	addVar("PUBLIC_CRON_INTERVAL", Args.CronInterval)
	addVar("PUBLIC_NEAR_EXPIRY_THRESHOLD", Args.NearExpiryThreshold)

	fmt.Println("----------- content", envContent.String())

	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		if len(pair) == 2 {
			addVar("PUBLIC_"+pair[0], pair[1])
		}
	}
	fmt.Println("----------- content after", envContent.String())

	err := os.WriteFile("/build/frontend/.env", []byte(envContent.String()), 0644)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}
