package main

import (
	"flag"
	"github.com/devvdark0/url-shortener/internal/config"
	"go.uber.org/zap"
)

const (
	envLocal = "local"
	envProd  = "prod"
)

func main() {
	configPath := flag.String("config-path", "", "path to config file")
	flag.Parse()

	cfg := config.MustLoad(*configPath)

	log := configureLogger(cfg.Env)
	log.Info("logger successfully set up")
	//TODO: init storage

	//TODO: init router

	//TODO: run the server

}

func configureLogger(env string) *zap.Logger {
	var (
		log *zap.Logger
		err error
	)

	switch env {
	case envLocal:
		log, err = zap.NewDevelopment()
	case envProd:
		log, err = zap.NewProduction()
	}

	if err != nil {
		panic("failed to set up logger: " + err.Error())
	}

	return log
}
