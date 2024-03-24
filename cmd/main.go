package main

import (
	"context"
	"fmt"
	"task/internal/config"
	repository "task/internal/repository/redis"
	"task/internal/usecase"
	"time"
)

func main() {
	redis := repository.CreateRedisRepository()
	var conf config.ConfigFloodControl
	conf.MaxQuantityQuery = 2
	conf.TimeLimit = 10 * time.Second
	fc := usecase.New(conf, &redis)
	ctx := context.Background()
	fmt.Println(fc.Check(ctx, 11))
	// В идеале стоит добавить defer для очистки Redis после работы программы(если сделать серверный вариант)
}

// FloodControl интерфейс, который нужно реализовать.
// Рекомендуем создать директорию-пакет, в которой будет находиться реализация.
type FloodControl interface {
	// Check возвращает false если достигнут лимит максимально разрешенного
	// кол-ва запросов согласно заданным правилам флуд контроля.
	Check(ctx context.Context, userID int64) (bool, error)
}
