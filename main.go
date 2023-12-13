/*
Copyright 2022 Chainguard, Inc.
SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/chainguard-dev/registry-redirect/pkg/redirect"
	"github.com/wlynch/slogctx"
)

func init() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})))
}

func main() {
	flag.Parse()
	logger := slogctx.FromContext(context.Background())

	http.Handle("/", redirect.New())

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	logger.Infof("Listening on port %s", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
