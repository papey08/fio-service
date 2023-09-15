package main

import (
	"context"
	"errors"
	"fio-service/internal/adapters/apis"
	"fio-service/internal/adapters/publisher"
	"fio-service/internal/app"
	"fio-service/internal/ports/consumer"
	"fio-service/internal/ports/graphql"
	"fio-service/internal/ports/rest"
	"fio-service/internal/repo"
	"fio-service/pkg/logger"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v5"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func FioRepoConfig(ctx context.Context, dbUrl string) *pgx.Conn {
	for {
		conn, err := pgx.Connect(ctx, dbUrl)
		if err != nil {
			logger.Error("unable to connect to postgres, trying again")
			time.Sleep(time.Second)
		} else {
			return conn
		}
	}
}

func ListenWithGracefulShutdown(ctx context.Context, srv *http.Server, eg *errgroup.Group) {
	eg.Go(func() error {
		defer logger.Info("server gracefully stopped")
		errCh := make(chan error)

		defer func() {
			shCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			_ = srv.Shutdown(shCtx)
			close(errCh)
		}()

		go func() {
			if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
				errCh <- err
			}
		}()

		select {
		case <-ctx.Done():
			return ctx.Err()
		case err := <-errCh:
			return fmt.Errorf("server unable to listen and serve requests: %w", err)
		}
	})

	_ = eg.Wait()
}

func main() {
	ctx := context.Background()

	// configuring logger
	logger.InitLogger(logger.DefaultLogger(os.Stdout))

	// configuring fio repo permanent storage
	pgConn := FioRepoConfig(ctx, "postgres://postgres:postgres@localhost:5432/fio_service?sslmode=disable")
	defer func() {
		_ = pgConn.Close(ctx)
	}()
	logger.Info("connected to postgres successfully")

	// configuring fio repo cache
	redisCache := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	defer func() {
		_ = redisCache.Close()
	}()
	logger.Info("connected to redis successfully")

	// configuring kafka FIO_FAILED publisher
	p := publisher.NewFioFailedTopic("localhost:9092", "FIO_FAILED")
	defer func() {
		_ = p.Writer.Close()
	}()
	logger.Info("connected to kafka FIO_FAILED successfully")

	// configuring app
	fr := repo.NewRepo(pgConn, redisCache)
	a := app.NewApp(fr, &p, &apis.Apis{})

	wg := new(sync.WaitGroup)
	wg.Add(3)

	// configuring rest server
	restServer := rest.NewRESTServer("localhost:8080", a)

	//configuring graphql server
	graphqlServer, err := graphql.NewGraphQLServer(ctx, a, "localhost:8081")
	if err != nil {
		logger.Fatal("error creating graphql server: %s", err.Error())
	}

	// configuring graceful shutdown
	sigQuit := make(chan os.Signal, 1)
	defer close(sigQuit)
	signal.Ignore(syscall.SIGHUP, syscall.SIGPIPE)
	signal.Notify(sigQuit, syscall.SIGINT, syscall.SIGTERM)

	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		select {
		case s := <-sigQuit:
			return fmt.Errorf("captured signal: %v", s)
		case <-ctx.Done():
			return nil
		}
	})

	go func() {
		defer wg.Done()
		logger.Info("started rest server successfully")
		ListenWithGracefulShutdown(ctx, restServer, eg)
	}()

	go func() {
		defer wg.Done()
		logger.Info("started graphql server successfully")
		ListenWithGracefulShutdown(ctx, graphqlServer, eg)
	}()

	// configuring kafka FIO consumer
	ft := consumer.NewFioTopic(a, "localhost:9092", "FIO")

	defer func() {
		_ = ft.Reader.Close()
	}()
	logger.Info("connected to kafka FIO successfully")

	go func() {
		defer wg.Done()
		ft.ListenFio(ctx)
	}()

	wg.Wait()
}
