package handlers

import (
	"context"
	"go.uber.org/zap"
)

type URLRepository interface {
	SaveURL(ctx context.Context, urlToSave, alias string) (int64, error)
	GetURL(ctx context.Context, alias string) (string, error)
	DeleteURL(ctx context.Context, alias string) error
}

type URLHandler struct {
	storage URLRepository
	log     *zap.Logger
}

func NewHandler(storage URLRepository, log *zap.Logger) *URLHandler {
	return &URLHandler{
		storage: storage,
		log:     log,
	}
}
