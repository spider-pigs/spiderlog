package spiderlog_test

import (
	"context"
	"testing"

	"github.com/spider-pigs/spiderlog"
)

func TestEmptyOpts(t *testing.T) {
	opts := spiderlog.Options{}
	logger, err := spiderlog.NewLogger(context.Background(), opts)
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
	opts := spiderlog.Options{
		Stdout: &stdout}
	logger, err := spiderlog.NewLogger(context.Background(), opts)
	if err != nil {
		t.Error("could not create logger", err)
	}
	if logger == nil {
		t.Error("could not create logger")
	}
	logger.Info("you should not see this (with go test -v)")
	logger.Close()
}
