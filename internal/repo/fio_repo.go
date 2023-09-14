package repo

import (
	"context"
	"errors"
	"fio-service/internal/app"
	"fio-service/internal/model"
	"fio-service/internal/repo/cache"
	"fio-service/internal/repo/permanent"
	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v5"
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

	// DeleteFioByKey deletes fio by key
	DeleteFioByKey(ctx context.Context, key uint) error
}

type Repo struct {
	permanentRepo
	cacheRepo
}

func NewRepo(conn *pgx.Conn, rc *redis.Client) app.FioRepo {
	return &Repo{
		permanentRepo: &permanent.Repo{
			Conn: *conn,
		},
		cacheRepo: &cache.Repo{
			Client: *rc,
		},
	}
}

func (r *Repo) AddFio(ctx context.Context, f model.Fio) (model.Fio, error) {
	fio, err := r.InsertFio(ctx, f) // add fio to permanent db
	if err != nil {
		return model.Fio{}, err
	}

	_, err = r.SetFioByKey(ctx, fio) // add fio to cache
	if err != nil {
		return model.Fio{}, err
	}
	return fio, nil
}

func (r *Repo) GetFioById(ctx context.Context, id uint) (model.Fio, error) {
	if fio, err := r.GetFioByKey(ctx, id); errors.Is(err, model.ErrorFioRepo) { // case when something wrong with cache
		return model.Fio{}, err
	} else if err == nil { // case when fio was found in cache
		return fio, nil
	}

	// case when fio not in cache
	if fio, err := r.SelectFioById(ctx, id); err != nil { // case when fio not in cache and not in db
		return model.Fio{}, err
	} else { // case when fio in db and not in cache
		if _, err = r.SetFioByKey(ctx, fio); err != nil {
			return model.Fio{}, err
		}
		return fio, nil
	}
}

func (r *Repo) GetFioByFilter(ctx context.Context, f model.Filter) ([]model.Fio, error) {
	return r.SelectFioByFilter(ctx, f)
}

func (r *Repo) UpdateFio(ctx context.Context, id uint, f model.Fio) (model.Fio, error) {
	fio, err := r.permanentRepo.UpdateFio(ctx, id, f) // update fio in permanent db
	if err != nil {
		return model.Fio{}, err
	}

	_, err = r.SetFioByKey(ctx, fio) // update fio in cache
	if err != nil {
		return model.Fio{}, err
	}
	return fio, nil
}

func (r *Repo) DeleteFio(ctx context.Context, id uint) error {
	err := r.permanentRepo.DeleteFio(ctx, id) // delete fio from permanent db
	if err != nil {
		return err
	}

	err = r.DeleteFioByKey(ctx, id) // delete fio from cache
	if err != nil {
		return err
	}
	return nil
}
