package spiderlog

import (
	"log"
)

// Config type
type Config struct {
	DebugLogger   *log.Logger
	ErrorLogger   *log.Logger
	InfoLogger    *log.Logger
	WarningLogger *log.Logger
	StdoutEnabled bool
}

// Option type
type Option func(*Config)

// DebugLogger sets debug logger
func DebugLogger(logger *log.Logger) Option {
	return func(cfg *Config) {
		cfg.DebugLogger = logger
	}
}

// ErrorLogger sets error logger
func ErrorLogger(logger *log.Logger) Option {
	return func(cfg *Config) {
		cfg.ErrorLogger = logger
	}
}

// InfoLogger sets info logger
func InfoLogger(logger *log.Logger) Option {
	return func(cfg *Config) {
		cfg.InfoLogger = logger
	}
}

// WarningLogger sets warning logger
func WarningLogger(logger *log.Logger) Option {
	return func(cfg *Config) {
		cfg.WarningLogger = logger
	}
}

// StdoutEnabled sets if logging to stdout should be enabled or not. Defaults
// to true.
func StdoutEnabled(b bool) Option {
	return func(cfg *Config) {
		cfg.StdoutEnabled = b
	}
}

// Logger type
type Logger struct {
	cfg *Config
}

// New constructs a new logger
// It's recommended to make this a global instance called `log`.
func New(options ...Option) *Logger {
	cfg := &Config{StdoutEnabled: true}

	for _, option := range options {
		option(cfg)
	}

	return &Logger{cfg: cfg}
}

// DebugLogger returns the logger used for debug
func (logger *Logger) DebugLogger() *log.Logger {
	return logger.cfg.DebugLogger
}

// InfoLogger returns the logger used for info
func (logger *Logger) InfoLogger() *log.Logger {
	return logger.cfg.InfoLogger
}

// WarningLogger returns the logger used for warnings
func (logger *Logger) WarningLogger() *log.Logger {
	return logger.cfg.WarningLogger
}

// ErrorLogger returns the logger used for errors
func (logger *Logger) ErrorLogger() *log.Logger {
	return logger.cfg.ErrorLogger
}

// Debug prints debug messages to the logger
func (logger *Logger) Debug(v ...interface{}) {
	println(logger.cfg.DebugLogger, v...)
	logger.println(v...)
}

// Debugf prints debug messages to the logger
func (logger *Logger) Debugf(format string, v ...interface{}) {
	printf(logger.cfg.DebugLogger, format, v...)
	logger.printf(format, v...)
}

// Error prints error messages to the logger
func (logger *Logger) Error(v ...interface{}) {
	println(logger.cfg.ErrorLogger, v...)
	logger.println(v...)
}

// Errorf prints error messages to the logger
func (logger *Logger) Errorf(format string, v ...interface{}) {
	printf(logger.cfg.ErrorLogger, format, v...)
	logger.printf(format, v...)
}

// Fatal prints error message followed by a call to os.Exit(1).
func (logger *Logger) Fatal(v ...interface{}) {
	println(logger.cfg.ErrorLogger, v...)
	log.Fatal(v...)
}

// Fatalf prints error message followed by a call to os.Exit(1).
func (logger *Logger) Fatalf(format string, v ...interface{}) {
	printf(logger.cfg.ErrorLogger, format, v...)
	log.Fatalf(format, v...)
}

// Info prints info messages to the logger
func (logger *Logger) Info(v ...interface{}) {
	println(logger.cfg.InfoLogger, v...)
	logger.println(v...)
}

// Infof prints info messages to the logger
func (logger *Logger) Infof(format string, v ...interface{}) {
	printf(logger.cfg.InfoLogger, format, v...)
	logger.printf(format, v...)
}

// Warning prints warning messages to the logger
func (logger *Logger) Warning(v ...interface{}) {
	println(logger.cfg.WarningLogger, v...)
	logger.println(v...)
}

// Warningf prints warning messages to the logger
func (logger *Logger) Warningf(format string, v ...interface{}) {
	printf(logger.cfg.WarningLogger, format, v...)
	logger.printf(format, v...)
}

func (logger *Logger) printf(format string, v ...interface{}) {
	if logger.cfg.StdoutEnabled {
		log.Printf(format, v...)
	}
}

func (logger *Logger) println(v ...interface{}) {
	if logger.cfg.StdoutEnabled {
		log.Println(v...)
	}
}

func printf(logger *log.Logger, format string, v ...interface{}) {
	if logger != nil {
		logger.Printf(format, v...)
	}
}

func println(logger *log.Logger, v ...interface{}) {
	if logger != nil {
		logger.Println(v...)
	}
}
