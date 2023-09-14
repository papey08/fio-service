package graphql

import (
	"fio-service/internal/app"
	"github.com/graphql-go/graphql"
)

func rootQuery(a app.App) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "query",
		Fields: graphql.Fields{
			"get_fio_by_id": &graphql.Field{
				// TODO: implement
			},
			"get_fio_by_filter": &graphql.Field{
				// TODO: implement
			},
		},
	})
}
