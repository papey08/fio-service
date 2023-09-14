package graphql

import (
	"fio-service/internal/app"
	"github.com/graphql-go/graphql"
)

func rootMutation(a app.App) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "mutation",
		Fields: graphql.Fields{
			"add_fio": &graphql.Field{
				// TODO: implement
			},
			"update_fio": &graphql.Field{
				// TODO: implement
			},
			"delete_fio": &graphql.Field{
				// TODO: implement
			},
		},
	})
}
