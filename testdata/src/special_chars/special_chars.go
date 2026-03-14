package special_chars

import (
	"log/slog"

	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()

	slog.Info("server started")          // ok
	slog.Info("server started!")         // want `log message should not contain special characters`
	slog.Info("server started🚀")         // want `log message should not contain special characters`
	slog.Error("connection failed!!!")   // want `log message should not contain special characters`
	slog.Warn("something went wrong...") // want `log message should not contain special characters`

	logger.Info("server started")  // ok
	logger.Info("server started!") // want `log message should not contain special characters`
	logger.Info("server started🚀") // want `log message should not contain special characters`

	// non-logging methods must be ignored
	logger.With(zap.String("key!", "value"))
	logger.Named("service.main")
	slog.With("key!", "value")
}
