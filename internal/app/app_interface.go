package app

import (
	"context"
	"fio-service/internal/model"
)

type FioRepo interface {
	AddFio(ctx context.Context, f model.Fio) (model.Fio, error)
	GetFio(ctx context.Context, f model.Filter) ([]model.Fio, error)
	UpdateFio(ctx context.Context, id uint, f model.Fio) (model.Fio, error)
	DeleteFio(ctx context.Context, id int) error
}

type App interface {
	FillFio(ctx context.Context, f model.Fio) (model.Fio, error)
	GetFio(ctx context.Context, f model.Filter) ([]model.Fio, error)
	UpdateFio(ctx context.Context, id uint, f model.Fio) (model.Fio, error)
	DeleteFio(ctx context.Context, id int) error
}
