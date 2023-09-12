package repo

import (
	"context"
	"fio-service/internal/model"
)

type permanentRepo interface {
	// SelectFioById selects fio by given id
	SelectFioById(ctx context.Context, id uint) (model.Fio, error)

	// SelectFioByFilter selects all fios by given filter
	SelectFioByFilter(ctx context.Context, f model.Filter) ([]model.Fio, error)

	// InsertFio inserts new fio to database
	InsertFio(ctx context.Context, f model.Fio) (model.Fio, error)

	// UpdateFio updates fields of existing fio by id
	UpdateFio(ctx context.Context, id uint, f model.Fio) (model.Fio, error)

	// DeleteFio deletes fio by id
	DeleteFio(ctx context.Context, id uint) error
}

type cacheRepo interface {
	// GetFioByKey searches fio with given key
	GetFioByKey(ctx context.Context, key uint) (model.Fio, error)

	// SetFioByKey sets fio with its id as key
	SetFioByKey(ctx context.Context, f model.Fio) (model.Fio, error)

	// DeleteFio deletes fio by key
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

func (r *Repo) GetFioById(ctx context.Context, id uint) (model.Fio, error) {
	// TODO: implement
	return model.Fio{}, nil
}

func (r *Repo) GetFioByFilter(ctx context.Context, f model.Filter) ([]model.Fio, error) {
	// TODO: implement
	return nil, nil
}

func (r *Repo) UpdateFio(ctx context.Context, id uint, f model.Fio) (model.Fio, error) {
	// TODO: implement
	return model.Fio{}, nil
}

func (r *Repo) DeleteFio(ctx context.Context, id uint) error {
	// TODO: implement
	return nil
}
