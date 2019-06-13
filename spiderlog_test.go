package spiderlog

import (
	"context"
	"testing"
)

func TestNoUUIDAsUserID(t *testing.T) {
	logger, err := NewLogger(context.Background(), Options{})
	if err != nil {
		t.Error("Could not create logger", err)
	}
	if logger == nil {
		t.Error("Could not create logger")
	}
}
