package app

import (
	"context"
	"fio-service/internal/model"
)

type FioRepo interface {
	// AddFio adds given fio to database
	AddFio(ctx context.Context, f model.Fio) (model.Fio, error)

	// GetFioById searches fio in database with given id
	GetFioById(ctx context.Context, id int) (model.Fio, error)

	// GetFioByFilter searches fios in database by given filter
	GetFioByFilter(ctx context.Context, f model.Filter) ([]model.Fio, error)

	// UpdateFio updates fio fields with given id
	UpdateFio(ctx context.Context, id int, f model.Fio) (model.Fio, error)

	// DeleteFio deletes
	DeleteFio(ctx context.Context, id int) error
}

type Publisher interface {
	// SendFio sends invalid fio to message broker
	SendFio(ctx context.Context, fio model.Fio, reason string) error
}

type Apis interface {
	// GetAge makes a http request to get age by name
	GetAge(name string) (int, error)

	// GetGender makes a http request to get gender by name
	GetGender(name string) (string, error)

	// GetNation makes a http request to get nation by name
	GetNation(name string) (string, error)
}

type App interface {
	// FillFio fills fields Age, Gender and Nation of given fio and adds it to database
	FillFio(ctx context.Context, f model.Fio) (model.Fio, error)

	FioRepo
}
