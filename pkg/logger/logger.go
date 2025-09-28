package logger

import (
	"log/slog"
	"os"
	"sync"

	"github.com/chandan167/pharmacy-backend/pkg/slogmulti"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	once   sync.Once
	Logger *slog.Logger
)

// Init initializes the global logger
func Init(env string, filePath string) {
	once.Do(func() {
		var handlers []slog.Handler

		// File handler with log rotation
		if filePath != "" {
			fileHandler := slog.NewTextHandler(&lumberjack.Logger{
				Filename:   filePath,
				MaxSize:    10, // MB
				MaxBackups: 3,
				MaxAge:     28,   // days
				Compress:   true, // gzip
			}, &slog.HandlerOptions{Level: slog.LevelInfo})
			handlers = append(handlers, fileHandler)
		}

		// Console handler
		var consoleHandler slog.Handler
		if env == "development" {
			// consoleHandler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})
			consoleHandler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})
		} else {

		}
		handlers = append(handlers, consoleHandler)

		// Combine handlers
		multiHandler := slogmulti.NewMultiHandler(handlers...)
		Logger = slog.New(multiHandler)
	})
}
