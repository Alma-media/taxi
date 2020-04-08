package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/Alma-media/taxi/api"
	"github.com/Alma-media/taxi/config"
	"github.com/Alma-media/taxi/generator"
	"github.com/Alma-media/taxi/repository"
	"github.com/Alma-media/taxi/storage/spheric"
)

const prefix = "taxi: "

func main() {
	// TODO: consider using logrus
	logger := log.New(os.Stdout, prefix, log.Lshortfile)

	config := config.New()

	generator := generator.New(config.Generator)
	if err := generator.Init(); err != nil {
		logger.Fatalf("Cannot initialize order generator: %s", err)
	}
	defer generator.Close()

	storage := spheric.NewOrderStorage()

	repository := repository.NewOrderRepository(generator, storage)

	srv := http.Server{
		Addr:    config.HTTP.Address,
		Handler: api.NewHandler(repository),
	}

	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint
		logger.Println("Shutting down the server...")

		if err := srv.Shutdown(context.Background()); err != nil {
			logger.Printf("HTTP server Shutdown: %s", err)
		}
		close(idleConnsClosed)
	}()

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		logger.Fatalf("HTTP server ListenAndServe: %s", err)
	}

	<-idleConnsClosed
	logger.Println("Done")
}
