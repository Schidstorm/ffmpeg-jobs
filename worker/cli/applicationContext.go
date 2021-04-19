package cli

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

func ApplicationContext() context.Context {
	ctx, cancel := context.WithCancel(context.Background())

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Kill, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		<-signals
		cancel()
	}()

	return ctx
}
