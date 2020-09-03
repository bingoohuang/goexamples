package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"fx-example/cache"
	"fx-example/handlers"

	"github.com/go-redis/redis"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttprouter"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	kill := make(chan os.Signal, 1)
	signal.Notify(kill, os.Interrupt)

	go func() {
		sig := <-kill
		fmt.Println("signal", sig)
		cancel()
	}()

	app := fx.New(
		fx.Provide(newZapLogger, newRedisClient),
		fx.Provide(cache.NewMeaningOfLifeCacheRedis),
		fx.Provide(handlers.NewMeaningOfLifeHandler),
		fx.Invoke(runHttpServer),
	)
	if err := app.Start(ctx); err != nil {
		fmt.Println(err)
	}
}

func runHttpServer(lifecycle fx.Lifecycle, molHandler *handlers.MeaningOfLife) {
	lifecycle.Append(fx.Hook{OnStart: func(context.Context) error {
		r := fasthttprouter.New()
		r.Handle(http.MethodGet, "/what-is-the-meaning-of-life", molHandler.Handle)
		return fasthttp.ListenAndServe("localhost:8080", r.Handler)
	}})
}

func newRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{Addr: "localhost:6379"})
}
func newZapLogger() *zap.Logger {
	logger, _ := zap.NewProduction()
	return logger
}
