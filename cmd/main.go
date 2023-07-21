package main

import (
	"os"
	"time"

	"recitas/config"
	"recitas/server"
	"recitas/service"

	"golang.org/x/sync/errgroup"
	"recitas/modules/users"
)

const gracefulShutdownAwait = 1 * time.Second

func init() {
	dsn := os.Getenv("DSN")
	signKey := os.Getenv("SIGNKEY")
	if len(dsn)*len(signKey) == 0 {
		panic("отсутствуют переменные окружения")
	}

	config.SetDSN(dsn)
	users.SetSignKey(signKey)
}

func main() {
	ctx, cancelCtx := server.MakeContext()
	cfg := config.MustConfig(ctx)

	// переброска кук в контекст запросов
	handler := service.HeadersTransmitter(ctx, service.NewService(cfg))

	group, ctx := errgroup.WithContext(ctx)
	group.Go(func() error {
		server.MakeServer(ctx, cancelCtx, gracefulShutdownAwait, server.ServeOrigin([]string{"*"}, handler), "0.0.0.0:8000")
		return nil
	})

	_ = group.Wait()

	time.Sleep(gracefulShutdownAwait)
	_ = cfg.Postgres.Conn.Close()
	time.Sleep(gracefulShutdownAwait)
}
