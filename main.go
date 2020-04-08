package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"

	"github.com/Alma-media/taxi/api"
	"github.com/Alma-media/taxi/config"
	"github.com/Alma-media/taxi/generator"
	"github.com/Alma-media/taxi/repository"
	"github.com/Alma-media/taxi/storage/spheric"

	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

const service = "order"

func main() {
	logger := logrus.New().WithField("service", service)

	var config config.Config
	// TODO: find a better ENV parser
	if err := envconfig.Process(service, &config); err != nil {
		logger.Fatalf("Unable to read the configuration: %s", err)
	}

	logger.Infof("Starting the service with configuration: %#v", config)

	generator := generator.New(config.Generator)
	if err := generator.Init(); err != nil {
		logger.Fatalf("Cannot initialize order generator: %s", err)
	}
	defer generator.Close()

	storage := spheric.NewOrderStorage()

	repository := repository.NewOrderRepository(generator, storage)

	// TODO: consider using fasthttp.Server to increase the performance
	srv := http.Server{
		Addr:    config.HTTP.Address,
		Handler: api.NewHandler(repository),
	}

	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint
		logger.Infof("Shutting down the server...")

		if err := srv.Shutdown(context.Background()); err != nil {
			logger.Printf("Unable to shutdown HTTP server: %s", err)
		}
		close(idleConnsClosed)
	}()

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		logger.Fatalf("Failed to start HTTP server: %s", err)
	}

	<-idleConnsClosed
	logger.Infof("HTTP server has been stopped")
}
