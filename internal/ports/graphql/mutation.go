package graphql

import (
	"context"
	"fio-service/internal/app"
	"fio-service/internal/model"
	"fio-service/pkg/logger"
	"github.com/graphql-go/graphql"
)

func rootMutation(ctx context.Context, a app.App) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "mutation",
		Fields: graphql.Fields{
			"addFio": &graphql.Field{
				Type: fioType,
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"surname": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"patronymic": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"age": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"gender": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"nation": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (any, error) {
					var f model.Fio
					f.Name, _ = p.Args["name"].(string)
					f.Surname, _ = p.Args["surname"].(string)
					f.Patronymic, _ = p.Args["patronymic"].(string)
					f.Age, _ = p.Args["age"].(int)
					f.Gender, _ = p.Args["gender"].(string)
					f.Nation, _ = p.Args["nation"].(string)
					logger.Info("adding fio by graphql server: %s %s", f.Name, f.Surname)
					return a.AddFio(ctx, f)
				},
			},
			"updateFio": &graphql.Field{
				Type: fioType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"surname": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"patronymic": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"age": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"gender": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"nation": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (any, error) {
					id, _ := p.Args["id"].(int)

					var f model.Fio
					f.Name, _ = p.Args["name"].(string)
					f.Surname, _ = p.Args["surname"].(string)
					f.Patronymic, _ = p.Args["patronymic"].(string)
					f.Age, _ = p.Args["age"].(int)
					f.Gender, _ = p.Args["gender"].(string)
					f.Nation, _ = p.Args["nation"].(string)

					logger.Info("updating fio with id %d by graphql server", id)
					return a.UpdateFio(ctx, id, f)
				},
			},
			"deleteFio": &graphql.Field{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(p graphql.ResolveParams) (any, error) {
					id, _ := p.Args["id"].(int)
					logger.Info("deleting fio with id %d by graphql server", id)
					if err := a.DeleteFio(ctx, id); err != nil {
						return false, err
					} else {
						return true, nil
					}
				},
			},
		},
	})
}
