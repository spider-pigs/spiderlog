package spiderlog

import (
	"context"
	"log"

	"cloud.google.com/go/logging"
)

// Logger type
type Logger struct {
	stackdriverEnabled bool
	stdoutEnabled      bool
	client             *logging.Client
	d                  *log.Logger
	i                  *log.Logger
	w                  *log.Logger
	e                  *log.Logger
}

// Options contains options for configuring the logger.
type Options struct {
	// ProjectID is the identifier of the Stackdriver
	// project the user is uploading the logs to.
	ProjectID string

	// A log ID must be less than 512 characters long and can only
	// include the following characters: upper and lower case alphanumeric
	// characters: [A-Za-z0-9]; and punctuation characters: forward-slash,
	// underscore, hyphen, and period.
	LogID string

	// Should logs be written to stdout? Defaults to true.
	Stdout *bool
}

// NewLogger constructs a new logger
// It's recommended to make this a global instance called `log`.
func NewLogger(ctx context.Context, opts Options) (*Logger, error) {
	stdout := true
	if opts.Stdout != nil {
		stdout = *opts.Stdout
	}

	projectID := opts.ProjectID
	if len(projectID) == 0 {
		logger := &Logger{
			stdoutEnabled:      stdout,
			stackdriverEnabled: false}
		return logger, nil
	}

	project := "projects/" + projectID
	client, err := logging.NewClient(ctx, project)
	if err != nil {
		return nil, err
	}

	logID := opts.LogID

	d := client.Logger(logID).StandardLogger(logging.Debug)
	i := client.Logger(logID).StandardLogger(logging.Info)
	w := client.Logger(logID).StandardLogger(logging.Warning)
	e := client.Logger(logID).StandardLogger(logging.Error)

	logger := &Logger{
		client:             client,
		stackdriverEnabled: true,
		stdoutEnabled:      stdout,
		d:                  d,
		i:                  i,
		w:                  w,
		e:                  e}

	return logger, nil
}

// Debug prints debug messages to the logger
func (logger *Logger) Debug(v ...interface{}) {
	if logger.stackdriverEnabled {
		logger.d.Println(v...)
	}
	if logger.stdoutEnabled {
		log.Println(v...)
	}
}

// Debugf prints debug messages to the logger
func (logger *Logger) Debugf(format string, v ...interface{}) {
	if logger.stackdriverEnabled {
		logger.d.Printf(format, v...)
	}
	if logger.stdoutEnabled {
		log.Printf(format, v...)
	}
}

// Error prints error messages to the logger
func (logger *Logger) Error(v ...interface{}) {
	if logger.stackdriverEnabled {
		logger.e.Println(v...)
	}
	if logger.stdoutEnabled {
		log.Println(v...)
	}
}

// Errorf prints error messages to the logger
func (logger *Logger) Errorf(format string, v ...interface{}) {
	if logger.stackdriverEnabled {
		logger.e.Printf(format, v...)
	}
	if logger.stdoutEnabled {
		log.Printf(format, v...)
	}
}

// Fatal prints error message followed by a call to os.Exit(1).
func (logger *Logger) Fatal(v ...interface{}) {
	if logger.stackdriverEnabled {
		logger.e.Fatal(v...)
	}
	if logger.stdoutEnabled {
		log.Fatal(v...)
	}
}

// Fatalf prints error message followed by a call to os.Exit(1).
func (logger *Logger) Fatalf(format string, v ...interface{}) {
	if logger.stackdriverEnabled {
		logger.e.Fatalf(format, v...)
	}
	if logger.stdoutEnabled {
		log.Fatalf(format, v...)
	}
}

// Info prints info messages to the logger
func (logger *Logger) Info(v ...interface{}) {
	if logger.stackdriverEnabled {
		logger.i.Println(v...)
	}
	if logger.stdoutEnabled {
		log.Println(v...)
	}
}

// Infof prints info messages to the logger
func (logger *Logger) Infof(format string, v ...interface{}) {
	if logger.stackdriverEnabled {
		logger.i.Printf(format, v...)
	}
	if logger.stdoutEnabled {
		log.Printf(format, v...)
	}
}

// Warning prints warning messages to the logger
func (logger *Logger) Warning(v ...interface{}) {
	if logger.stackdriverEnabled {
		logger.w.Println(v...)
	}
	if logger.stdoutEnabled {
		log.Println(v...)
	}
}

// Warningf prints warning messages to the logger
func (logger *Logger) Warningf(format string, v ...interface{}) {
	if logger.stackdriverEnabled {
		logger.w.Printf(format, v...)
	}
	if logger.stdoutEnabled {
		log.Printf(format, v...)
	}
}

// Close waits for all opened loggers to be flushed and closes the client.
func (logger *Logger) Close() {
	if logger.client != nil {
		logger.client.Close()
	}
}
