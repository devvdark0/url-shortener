package main

import (
	"flag"
	"github.com/devvdark0/url-shortener/internal/config"
	"github.com/devvdark0/url-shortener/internal/storage/sqlite"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
	"os"
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

	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("failed to init storage:", zap.Error(err))
		os.Exit(1)
	}
	_ = storage
	//TODO: init router
	router := configureRouter()
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

func configureRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RealIP)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)

	return r
}
