package graphql

import (
	"fio-service/internal/app"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"net/http"
)

func NewGraphQLServer(a app.App, addr string) (*http.Server, error) {
	query := rootQuery(a)
	mutation := rootMutation(a)

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
