package graphql

import (
	"context"
	"fio-service/internal/app"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"net/http"
)

func NewGraphQLServer(ctx context.Context, a app.App, addr string) (*http.Server, error) {
	query := rootQuery(ctx, a)
	mutation := rootMutation(ctx, a)

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    query,
		Mutation: mutation,
	})
	if err != nil {
		return nil, err
	}

	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})

	server := &http.Server{
		Addr:    addr,
		Handler: h,
	}
	return server, nil
}
