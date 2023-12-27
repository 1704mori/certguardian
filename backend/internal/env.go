package env

var Args struct {
	Port                string `arg:"env:PORT" default:"7070" help:"Port to listen on"`
	CronInterval        string `arg:"env:CRON_INTERVAL" default:"1d" help:"Interval for the cron job. e.g: 15d (15 days) or 3h (3 hours)"`
	NearExpiryThreshold string `arg:"env:NEAR_EXPIRY_THRESHOLD" default:"10d" help:"Threshold to determine if a certificate is near expiration. e.g: 15d (15 days) or 3h (3 hour"`
}
