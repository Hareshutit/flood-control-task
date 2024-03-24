package main

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
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
	mux := http.NewServeMux()

	mux.HandleFunc("/query", func(w http.ResponseWriter, r *http.Request) {
		idS := r.URL.Query().Get("id")
		if idS == "" {
			fmt.Fprintf(w, "Неправильное id пользовавтеля\n")
		}
		id, err := strconv.ParseInt(idS, 10, 64)
		if err != nil {
			fmt.Fprintf(w, err.Error()+"\n")
		}
		res, err := fc.Check(ctx, id)
		if err != nil {
			fmt.Fprintf(w, err.Error()+"\n")
		}
		fmt.Fprintf(w, "Разрешение: %t \n", res)
	})
	s := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	s.ListenAndServe()
}

// FloodControl интерфейс, который нужно реализовать.
// Рекомендуем создать директорию-пакет, в которой будет находиться реализация.
type FloodControl interface {
	// Check возвращает false если достигнут лимит максимально разрешенного
	// кол-ва запросов согласно заданным правилам флуд контроля.
	Check(ctx context.Context, userID int64) (bool, error)
}
