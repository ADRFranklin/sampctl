package types

// WebConfig contains configuration for the HTTP API server
type WebConfig struct {
	IP   string `json:"ip"   yaml:"ip"`   // interface to bind to
	Port string `json:"port" yaml:"port"` // port to expose
}
