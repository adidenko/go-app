package cfg

import (
	"github.com/adidenko/go-app/internal/metrics"
)

// Config is the main app config struct
type AppConfig struct {
	Metrics metrics.AppMetrics
}
