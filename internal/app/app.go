package app

import (
	"context"
	"fio-service/internal/model"
)

type app struct {
	FioRepo
}

func (a *app) FillFio(ctx context.Context, f model.Fio) (model.Fio, error) {
	// TODO: implement
	return model.Fio{}, nil
}

func (a *app) GetFio(ctx context.Context, f model.Filter) ([]model.Fio, error) {
	// TODO: implement
	return nil, nil
}

func (a *app) UpdateFio(ctx context.Context, id uint, f model.Fio) (model.Fio, error) {
	// TODO: implement
	return model.Fio{}, nil
}

func (a *app) DeleteFio(ctx context.Context, id int) error {
	// TODO: implement
	return nil
}
