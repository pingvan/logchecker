package english

import (
	"log/slog"

	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()

	slog.Info("starting server")     // ok
	slog.Info("запуск сервера")      // want `log message should contain only English letters, digits and spaces`
	slog.Error("ошибка подключения") // want `log message should contain only English letters, digits and spaces`

	logger.Info("starting server") // ok
	logger.Info("запуск сервера")  // want `log message should contain only English letters, digits and spaces`
}
