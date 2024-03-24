package main

import (
	"context"
	"fmt"
	"task/internal/config"
	repository "task/internal/repository/redis"
	"task/internal/usecase"
)

func main() {
	redis := repository.CreateRedisRepository()
	var conf *config.ConfigFloodControl
	conf, err := config.New("../configuratiion/config.yml")
	if err != nil {
		panic(err)
	}
	fc := usecase.New(*conf, &redis)
	ctx := context.Background()
	fmt.Println(*conf)
	fmt.Println(fc.Check(ctx, 122))
	// В идеале стоит добавить defer для очистки Redis после работы программы(если сделать серверный вариант)
}

// FloodControl интерфейс, который нужно реализовать.
// Рекомендуем создать директорию-пакет, в которой будет находиться реализация.
type FloodControl interface {
	// Check возвращает false если достигнут лимит максимально разрешенного
	// кол-ва запросов согласно заданным правилам флуд контроля.
	Check(ctx context.Context, userID int64) (bool, error)
}
