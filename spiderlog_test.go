package spiderlog_test

import (
	"testing"

	"github.com/spider-pigs/spiderlog"
)

func TestLogger(t *testing.T) {
	logger := spiderlog.New(spiderlog.StdoutEnabled(false))
	logger.Info("you should not see this (with go test -v)")
}
