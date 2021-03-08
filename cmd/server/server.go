package main

import (
	"context"
	"os"
	"os/signal"

	"go.uber.org/zap/zapcore"

	"github.com/kic/media/internal/setup"
	"github.com/kic/media/pkg/logging"
)

func main() {
	logger := logging.CreateLogger(zapcore.InfoLevel)

	repo, mongoClient := setup.DBRepositorySetup(logger, "media")

	serv := setup.GRPCSetup(logger, repo)

	defer serv.Stop()
	defer mongoClient.Disconnect(context.Background())

	// the server is listening in a goroutine so hang until we get an interrupt signal
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	<-c
}
