package spiderlog

import (
	"context"
	"testing"
)

func TestEmptyOpts(t *testing.T) {
	opts := Options{}
	logger, err := NewLogger(context.Background(), opts)
	if err != nil {
		t.Error("could not create logger", err)
	}
	if logger == nil {
		t.Error("could not create logger")
	}
	logger.Info("you should see this (with go test -v)")
	logger.Close()
}

func TestDisableStdout(t *testing.T) {
	stdout := false
	opts := Options{
		Stdout: &stdout}
	logger, err := NewLogger(context.Background(), opts)
	if err != nil {
		t.Error("could not create logger", err)
	}
	if logger == nil {
		t.Error("could not create logger")
	}
	logger.Info("you should not see this (with go test -v)")
	logger.Close()
}
