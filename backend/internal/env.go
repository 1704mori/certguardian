package env

var Args struct {
	Port                string `arg:"env:DOKKA_PORT" json:"port" default:"7070" help:"Port to listen on"`
	EventStreamInterval int    `arg:"env:DOKKA_EVENT_STREAM_INTERVAL" json:"event_stream_interval" default:"5" help:"Interval in seconds to send events to the frontend"`
}
