package usecase

import (
	"context"
	"task/internal/config"
	"time"
)

type CheckRepository interface {
	Check(ctx context.Context, userID int64, TimeLimit time.Duration, MaxQuantityQuery int) (bool, error)
}

type FloodControlHandler struct {
	TimeLimit        time.Duration
	MaxQuantityQuery int

	Repository CheckRepository
}

func New(conf config.ConfigFloodControl, rep CheckRepository) FloodControlHandler {
	return FloodControlHandler{
		TimeLimit:        conf.TimeLimit,
		MaxQuantityQuery: conf.MaxQuantityQuery,

		Repository: rep,
	}
}

func (f *FloodControlHandler) Check(ctx context.Context, userID int64) (bool, error) {
	return f.Repository.Check(ctx, userID, f.TimeLimit, f.MaxQuantityQuery)
}
