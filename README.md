# spiderlog
[![Build Status](https://travis-ci.org/spider-pigs/spiderlog.svg?branch=master)](https://travis-ci.org/spider-pigs/spiderlog) [![Go Report Card](https://goreportcard.com/badge/github.com/spider-pigs/spiderlog)](https://goreportcard.com/report/github.com/spider-pigs/spiderlog) [![Codacy Badge](https://api.codacy.com/project/badge/Grade/ca5c6fceda1547ba92ba4c4673db461f)](https://www.codacy.com/app/spider-pigs/spiderlog?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=spider-pigs/spiderlog&amp;utm_campaign=Badge_Grade) [![GoDoc](https://godoc.org/github.com/spider-pigs/spiderlog?status.svg)](https://godoc.org/github.com/spider-pigs/spiderlog)

spiderlog is a simple logging client.

## Install

```Golang
import "github.com/spider-pigs/spiderlog"
```

## Usage

There are four logging levels supported: `Debug`, `Info`, `Warning`, `Error`

Here is an example of how to setup logging for a Google logging client.

```Golang
package main

import (
	"context"

	"cloud.google.com/go/logging"
	"github.com/spider-pigs/spiderlog"
)

// Initialize log to guard from nil pointer dereferences
var log = spiderlog.New()

func main() {
	project := "projects/some_project_id"
	client, err := logging.NewClient(context.Background(), project)
	if err != nil {
		log.Fatal(err)
	}

	log = spiderlog.New(
		spiderlog.DebugLogger(client.Logger("a_log_id").StandardLogger(logging.Debug)),
		spiderlog.InfoLogger(client.Logger("a_log_id").StandardLogger(logging.Info)),
		spiderlog.WarningLogger(client.Logger("a_log_id").StandardLogger(logging.Warning)),
		spiderlog.ErrorLogger(client.Logger("a_log_id").StandardLogger(logging.Error)),
		spiderlog.StdoutEnabled(false), // Defaults to true
	)

	// Write an info log
	log.Info("Hello World!")

	// Or log with formatting
	log.Infof("Hello %s!\n", "World")
}
```
