package value_objects

type RouteDefinition struct {
	Method    string `json:"method"`
	Path      string `json:"path"`
	Protected bool   `json:"protected"`
	RateLimit int    `json:"rate_limit,omitempty"`
	Timeout   int    `json:"timeout,omitempty"`
}
