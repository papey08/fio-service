package graphql

import (
	"context"
	"fio-service/internal/app"
	"fio-service/internal/model"
	"fio-service/pkg/logger"
	"github.com/graphql-go/graphql"
)

func rootQuery(ctx context.Context, a app.App) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "query",
		Fields: graphql.Fields{
			"getFioById": &graphql.Field{
				Type: fioType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(p graphql.ResolveParams) (any, error) {
					if id, ok := p.Args["id"].(int); ok {
						logger.Info("getting fio with id %d by graphql server", id)
						return a.GetFioById(ctx, id)
					}
					return nil, model.ErrorInvalidInput
				},
			},
			"getFioByFilter": &graphql.Field{
				Type: graphql.NewList(fioType),
				Args: graphql.FieldConfigArgument{
					"offset": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"limit": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"surname": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"patronymic": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"age": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"gender": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"nation": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (any, error) {
					var f model.Filter
					if offset, ok := p.Args["offset"].(int); ok {
						f.Offset = offset
					}
					if limit, ok := p.Args["limit"].(int); ok {
						f.Limit = limit
					}
					if name, ok := p.Args["name"].(string); ok {
						f.ByName = true
						f.Name = name
					}
					if surname, ok := p.Args["surname"].(string); ok {
						f.BySurname = true
						f.Surname = surname
					}
					if patronymic, ok := p.Args["patronymic"].(string); ok {
						f.ByPatronymic = true
						f.Patronymic = patronymic
					}
					if age, ok := p.Args["age"].(int); ok {
						f.ByAge = true
						f.Age = age
					}
					if gender, ok := p.Args["gender"].(string); ok {
						f.ByGender = true
						f.Gender = gender
					}
					if nation, ok := p.Args["nation"].(string); ok {
						f.ByNation = true
						f.Nation = nation
					}
					logger.Info("getting fios with filter by graphql server")
					return a.GetFioByFilter(ctx, f)
				},
			},
		},
	})
}
