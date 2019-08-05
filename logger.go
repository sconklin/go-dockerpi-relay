package relay

import logger "github.com/sconklin/go-logger"

// You can manage verbosity of log output
// in the package by changing last parameter value.
var logrelay = logger.NewPackageLogger("relay",
	// logger.DebugLevel,
	logger.InfoLevel,
)
