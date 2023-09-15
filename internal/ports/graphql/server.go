package graphql

import (
	"context"
	"fio-service/internal/app"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"net/http"
)

// NewGraphQLServer creates graphql server with queries and mutations
func NewGraphQLServer(ctx context.Context, addr string, a app.App) (*http.Server, error) {
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
