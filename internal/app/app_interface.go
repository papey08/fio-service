package app

import (
	"context"
	"fio-service/internal/model"
)

type FioRepo interface {
	// AddFio adds given fio to database
	AddFio(ctx context.Context, f model.Fio) (model.Fio, error)

	// GetFioById searches fio in database with given id
	GetFioById(ctx context.Context, id uint) (model.Fio, error)

	// GetFioByFilter searches fios in database by given filter
	GetFioByFilter(ctx context.Context, f model.Filter) ([]model.Fio, error)

	// UpdateFio updates fio fields with given id
	UpdateFio(ctx context.Context, id uint, f model.Fio) (model.Fio, error)

	// DeleteFio deletes
	DeleteFio(ctx context.Context, id uint) error
}

type Publisher interface {
	// SendFio sends invalid fio to message broker
	SendFio(fio model.Fio, reason string)
}

type App interface {
	// FillFio fills fields Age, Gender and Nation of given fio and adds it to database
	FillFio(ctx context.Context, f model.Fio) (model.Fio, error)

	/*// GetFioById gets fio with given id
	GetFioById(ctx context.Context, id uint) (model.Fio, error)

	// GetFioByFilter gets all fios from database by given filter
	GetFioByFilter(ctx context.Context, f model.Filter) ([]model.Fio, error)

	// UpdateFio updates fields of existing fio by id
	UpdateFio(ctx context.Context, id uint, f model.Fio) (model.Fio, error)

	// DeleteFio deletes fio by id
	DeleteFio(ctx context.Context, id uint) error*/

	FioRepo
}
