package main

import (
	"context"
	"fio-service/internal/adapters/apis"
	"fio-service/internal/adapters/publisher"
	"fio-service/internal/app"
	"fio-service/internal/ports/consumer"
	"fio-service/internal/ports/graphql"
	"fio-service/internal/ports/rest"
	"fio-service/internal/repo"
	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v5"
	"log"
	"sync"
	"time"
)

func FioRepoConfig(ctx context.Context, dbUrl string) *pgx.Conn {
	for {
		conn, err := pgx.Connect(ctx, dbUrl)
		if err != nil {
			// TODO: log
			time.Sleep(time.Second)
		} else {
			return conn
		}
	}
}

func main() {
	ctx := context.Background()

	// configuring fio repo permanent storage
	pgConn := FioRepoConfig(ctx, "postgres://postgres:postgres@localhost:5432/fio_service?sslmode=disable")
	defer func() {
		_ = pgConn.Close(ctx)
	}()
	log.Println("connected to postgres successfully")

	// configuring fio repo cache
	redisCache := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	defer func() {
		_ = redisCache.Close()
	}()
	log.Println("connected to redis successfully")

	// configuring kafka FIO_FAILED
	p, err := publisher.NewFioFailedTopic(ctx, "tcp", "localhost:9092", "FIO_FAILED")
	if err != nil {
		// TODO: log
	}
	defer func() {
		_ = p.Conn.Close()
	}()
	log.Println("connected to kafka FIO_FAILED successfully")

	// configuring app
	fr := repo.NewRepo(pgConn, redisCache)
	a := app.NewApp(fr, &p, &apis.Apis{})

	wg := new(sync.WaitGroup)
	wg.Add(3)

	// configuring rest server
	restServer := rest.NewRESTServer("localhost:8080", a)
	go func() {
		defer wg.Done()
		_ = restServer.ListenAndServe()
	}()

	//configuring graphql server
	graphqlServer, err := graphql.NewGraphQLServer(ctx, a, "localhost:8081")
	if err != nil {
		log.Fatal("error creating graphql server:", err.Error())
	}
	go func() {
		defer wg.Done()
		_ = graphqlServer.ListenAndServe()
	}()

	// configuring kafka FIO
	ft, err := consumer.NewFioTopic(ctx, a, "tcp", "localhost:9092", "FIO")
	if err != nil {
		// TODO
	}
	defer func() {
		_ = ft.Conn.Close()
	}()
	log.Println("connected to kafka FIO successfully")

	go func() {
		defer wg.Done()
		ft.ListenFio(ctx)
	}()

	wg.Wait()

}
