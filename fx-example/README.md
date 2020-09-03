# Fx Medium Example

Full blog post can be found here https://medium.com/@erez.levi/using-uber-fx-to-simplify-dependency-injection-875363245c4c

## To Install

    go get fx-example

## To Run

```bash
🕙[2020-09-03 15:33:14.315] ❯ go run main.go
2020/09/03 15:33:31 [Fx] PROVIDE        *zap.Logger <= main.newZapLogger()
2020/09/03 15:33:31 [Fx] PROVIDE        *redis.Client <= main.newRedisClient()
2020/09/03 15:33:31 [Fx] PROVIDE        cache.MeaningOfLifeCache <= fx-example/cache.NewMeaningOfLifeCacheRedis()
2020/09/03 15:33:31 [Fx] PROVIDE        *handlers.MeaningOfLife <= fx-example/handlers.NewMeaningOfLifeHandler()
2020/09/03 15:33:31 [Fx] PROVIDE        fx.Lifecycle <= go.uber.org/fx.New.func1()
2020/09/03 15:33:31 [Fx] PROVIDE        fx.Shutdowner <= go.uber.org/fx.(*App).shutdowner-fm()
2020/09/03 15:33:31 [Fx] PROVIDE        fx.DotGraph <= go.uber.org/fx.(*App).dotGraph-fm()
2020/09/03 15:33:31 [Fx] INVOKE         main.runHttpServer()
2020/09/03 15:33:31 [Fx] START          main.runHttpServer()
^Csignal interrupt
context canceled
```

access http://localhost:8080/what-is-the-meaning-of-life you will get

> 这个概念通过许多相关问题体现出来，例如：“我为何在此”“什么是生命？”“生命的真谛是什么？”。
> 在历史长河中，它也是哲学，科学以及神学一直所思索的主题。前人在不同的文化环境与意识形态背景下也给出了很多的多元化答案。
> 阿尔贝·加缪指出，作为一个存在的人，人类用生命的价值和意义来说服自己：人的存在不是荒诞的。

## TL:DR

The main looks like this:

```go
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	kill := make(chan os.Signal, 1)
	signal.Notify(kill)

	go func() {
		<-kill
		cancel()
	}()

	app := fx.New(
		fx.Provide(newZapLogger),
		fx.Provide(newRedisClient),
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
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}
func newZapLogger() *zap.Logger {
	logger, _ := zap.NewProduction()
	return logger
}
```
