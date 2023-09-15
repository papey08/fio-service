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
	"github.com/joho/godotenv"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

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

	// loading env variables from .env
	if err := godotenv.Load("config/.env"); err != nil {
		logger.Fatal("unable to load .env file: %s", err.Error())
	}

	// configuring fio repo permanent storage
	postgresUser := os.Getenv("POSTGRESQL_USER")
	postgresPass := os.Getenv("POSTGRESQL_PASSWORD")
	postgresHost := os.Getenv("POSTGRESQL_HOST")
	postgresPort := os.Getenv("POSTGRESQL_PORT")
	postgresDb := os.Getenv("POSTGRESQL_DB")
	postgresSSL := os.Getenv("POSTGRESQL_SSLMODE")

	pgConn, err := pgx.Connect(ctx, fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		postgresUser, postgresPass,
		postgresHost, postgresPort,
		postgresDb, postgresSSL))
	if err != nil {
		logger.Fatal("unable to connect to postgresql: %s", err.Error())
	}
	defer func() {
		_ = pgConn.Close(ctx)
	}()
	logger.Info("connected to postgres successfully")

	// configuring fio repo cache
	redisDB, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		logger.Fatal("cannot get redis db: %s", err.Error())
	}

	redisCache := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       redisDB,
	})
	defer func() {
		_ = redisCache.Close()
	}()
	logger.Info("connected to redis successfully")

	// configuring kafka FIO_FAILED publisher
	fioFailedAddr := fmt.Sprintf("%s:%s", os.Getenv("KAFKA_PUBLISHER_HOST"), os.Getenv("KAFKA_PUBLISHER_PORT"))
	p := publisher.NewFioFailedTopic(fioFailedAddr, os.Getenv("KAFKA_PUBLISHER_TOPIC"))
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
	restServerAddr := fmt.Sprintf("%s:%s", os.Getenv("REST_SERVER_HOST"), os.Getenv("REST_SERVER_PORT"))
	restServer := rest.NewRESTServer(restServerAddr, a)

	//configuring graphql server
	graphqlServerAddr := fmt.Sprintf("%s:%s", os.Getenv("GRAPHQL_SERVER_HOST"), os.Getenv("GRAPHQL_SERVER_PORT"))
	graphqlServer, err := graphql.NewGraphQLServer(ctx, graphqlServerAddr, a)
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
	fioAddr := fmt.Sprintf("%s:%s", os.Getenv("KAFKA_CONSUMER_HOST"), os.Getenv("KAFKA_CONSUMER_PORT"))
	ft := consumer.NewFioTopic(a, fioAddr, os.Getenv("KAFKA_CONSUMER_TOPIC"))

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
