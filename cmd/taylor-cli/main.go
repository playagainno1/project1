package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"taylor-ai-server/internal/cmd"

	"github.com/sirupsen/logrus"
)

func main() {
	initLogger()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		defer cancel()
		waitSignal(ctx, syscall.SIGINT, syscall.SIGTERM)
	}()

	cmd := cmd.NewCommand("taylor-cli")
	err := cmd.ExecuteContext(ctx)
	if err != nil {
		logrus.WithError(err).Fatal("execute command")
	}
}

func waitSignal(ctx context.Context, sig ...os.Signal) error {
	sigC := make(chan os.Signal, 1)
	signal.Notify(sigC, sig...)
	select {
	case <-ctx.Done():
		return ctx.Err()
	case s := <-sigC:
		logrus.WithField("signal", s.String()).Info("stop command")
		return nil
	}
}

func initLogger() {
	formatter := &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02T15:04:05.000-07:00",
	}
	logrus.SetFormatter(formatter)
	logrus.SetLevel(logrus.DebugLevel)
}
