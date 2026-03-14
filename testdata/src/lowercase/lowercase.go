package lowercase

import (
	"log/slog"

	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()

	slog.Info("starting server")          // ok
	slog.Error("failed to connect")       // ok
	slog.Info("Starting server on port")  // want `log message must start with lowercase`
	slog.Error("Failed to connect to db") // want `log message must start with lowercase`

	logger.Info("starting server")    // ok
	logger.Error("failed to connect") // ok
	logger.Info("Starting server")    // want `log message must start with lowercase`
	logger.Error("Failed to connect") // want `log message must start with lowercase`
}
