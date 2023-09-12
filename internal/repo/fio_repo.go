package repo

import (
	"context"
	"fio-service/internal/model"
)

type permanentRepo interface {
	SelectFio(ctx context.Context, f model.Filter) ([]model.Fio, error)
	InsertFio(ctx context.Context, f model.Fio) (model.Fio, error)
	UpdateFio(ctx context.Context, id uint, f model.Fio) (model.Fio, error)
	DeleteFio(ctx context.Context, id uint) error
}

type cacheRepo interface {
	GetFioByKey(ctx context.Context, key string) (model.Fio, error)
	SetFioByKey(ctx context.Context, key string, f model.Fio) (model.Fio, error)
	DeleteFio(ctx context.Context, key string) error
}

type Repo struct {
	permanentRepo
	cacheRepo
}

func (r *Repo) AddFio(ctx context.Context, f model.Fio) (model.Fio, error) {
	// TODO: implement
	return model.Fio{}, nil
}

func (r *Repo) GetFio(ctx context.Context, f model.Filter) ([]model.Fio, error) {
	// TODO: implement
	return nil, nil
}

func (r *Repo) UpdateFio(ctx context.Context, id uint, f model.Fio) (model.Fio, error) {
	// TODO: implement
	return model.Fio{}, nil
}

func (r *Repo) DeleteFio(ctx context.Context, id int) error {
	// TODO: implement
	return nil
}
