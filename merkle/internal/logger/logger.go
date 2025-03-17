package logger

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/Bedrock-Technology/VeMerkle/internal/config"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

// InitLogger initializes the global logger
func InitLogger(config config.LogConfig) error {
	// Set log level
	level, err := logrus.ParseLevel(config.Level)
	if err != nil {
		return fmt.Errorf("invalid log level: %v", err)
	}
	logrus.SetLevel(level)

	// Set log format
	if config.Format == "json" {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	} else {
		logrus.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
		})
	}

	// Setup writers
	var writers []io.Writer

	// Add stdout if enabled
	if config.Stdout {
		writers = append(writers, os.Stdout)
	}

	// Setup file output if configured
	if config.File.Path != "" {
		// Ensure log directory exists
		logDir := filepath.Dir(config.File.Path)
		if err := os.MkdirAll(logDir, 0755); err != nil {
			return fmt.Errorf("failed to create log directory: %v", err)
		}

		// Configure lumberjack logger
		fileLogger := &lumberjack.Logger{
			Filename:   config.File.Path,
			MaxSize:    config.File.MaxSize,
			MaxBackups: config.File.MaxBackups,
			MaxAge:     config.File.MaxAge,
			Compress:   config.File.Compress,
		}

		writers = append(writers, fileLogger)
	}

	// Set output to multi-writer if we have multiple writers
	if len(writers) > 0 {
		logrus.SetOutput(io.MultiWriter(writers...))
	}

	return nil
}
