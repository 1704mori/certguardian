package env

var Args struct {
	Port                string `arg:"env:PORT" json:"port" default:"7070" help:"Port to listen on"`
	CronInterval        string `arg:"env:CRON_INTERVAL" json:"cron_interval" default:"1d" help:"Interval for the cron job. e.g: 15d (15 days) or 3h (3 hours)"`
	NearExpiryThreshold string `arg:"env:NEAR_EXPIRY_THRESHOLD" json:"near_expiry_threshold" default:"10d" help:"Threshold to determine if a certificate is near expiration. e.g: 15d (15 days) or 3h (3 hours)"`
}
