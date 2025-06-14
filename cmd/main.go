package main

import (
	"context"
	"fmt"
	"github.com/ne4chelovek/base-rest-todo/internal/app"
	"github.com/sirupsen/logrus"
	"net/http"
)

// @title Todo App API
// @version         1.0
// @description     API Server for TodoList Application.

// @host localhost:8000
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	server, err := app.SetupServer(ctx)
	if err != nil {
		logrus.Fatal(err)
	}

	errChan := make(chan error, 1)

	go runHTTPServer(server.HTTPServer, errChan)

	app.WaitForShutdown(ctx, cancel, errChan, server)

}

func runHTTPServer(s *http.Server, errChan chan<- error) {
	logrus.Printf("Listening and serving HTTP on %s", s.Addr)
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		errChan <- fmt.Errorf("Error starting server: %s", err)
	}
}
