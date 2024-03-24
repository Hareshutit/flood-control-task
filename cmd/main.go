package main

import (
	"context"
	"task/internal/config"
	repository "task/internal/repository/redis"
	"task/internal/usecase"
)

func main() {
	redis := repository.CreateRedisRepository()
	var conf config.ConfigFloodControl
	conf.MaxQuantityQuery = 20
	conf.TimeLimit = 1111111
	fc := usecase.New(conf, &redis)
	ctx := context.Background()
	fc.Check(ctx, 11)
}

// FloodControl интерфейс, который нужно реализовать.
// Рекомендуем создать директорию-пакет, в которой будет находиться реализация.
type FloodControl interface {
	// Check возвращает false если достигнут лимит максимально разрешенного
	// кол-ва запросов согласно заданным правилам флуд контроля.
	Check(ctx context.Context, userID int64) (bool, error)
}
