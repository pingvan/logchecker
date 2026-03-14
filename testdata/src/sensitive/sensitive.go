package sensitive

import (
	"log/slog"

	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	password := "secret"
	apiKey := "key"
	token := "tok"

	slog.Info("user authenticated")         // ok
	slog.Info("user password: " + password) // want `log message contains potentially sensitive data`
	slog.Info("api_key=" + apiKey)          // want `log message contains potentially sensitive data`

	logger.Info("user authenticated") // ok
	logger.Info("request completed",
		zap.String("username", "john")) // ok
	logger.Info("user logged in",
		zap.String("password", password)) // want `log field key contains potentially sensitive data`
	logger.Info("api request",
		zap.String("api_key", apiKey)) // want `log field key contains potentially sensitive data`
	logger.Debug("auth completed",
		zap.String("token", token)) // want `log field key contains potentially sensitive data`

	// non-logging methods must be ignored
	logger.With(zap.String("password", password))
	logger.Named("token-service")
	slog.With("api_key", apiKey)

	_ = password
	_ = apiKey
	_ = token
}
