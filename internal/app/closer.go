package app

import (
	"context"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func WaitForShutdown(ctx context.Context, cancel context.CancelFunc, errChan <-chan error, s *Server) {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGTERM, syscall.SIGINT)

	select {
	case <-sigchan:
		logrus.Println("Received shutdown signal")
	case err := <-errChan:
		logrus.Printf("Critical error: %v", err)
	case <-ctx.Done():
	}

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	logrus.Println("Stopping HTTP server...")
	s.HTTPServer.Shutdown(shutdownCtx)
	logrus.Println("Closing database connections...")
	s.DB.Close()

	cancel()
}
