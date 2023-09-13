package app

import (
	"context"
	"errors"
	"fio-service/internal/adapters/apis"
	"fio-service/internal/app/valid"
	"fio-service/internal/model"
)

type app struct {
	FioRepo
	Publisher
}

func (a *app) FillFio(ctx context.Context, f model.Fio) (model.Fio, error) {
	if err := valid.NonFilledFio(f); err != nil { // check if name, surname or patronymic are valid
		return model.Fio{}, err
	}

	// fill age field
	if age, err := apis.GetAge(f.Name); errors.Is(err, model.ErrorNonExistName) {
		return model.Fio{}, model.ErrorNonExistName
	} else if err != nil {
		return model.Fio{}, model.ErrorApi
	} else {
		f.Age = age
	}

	// fill gender field
	if gender, err := apis.GetGender(f.Name); errors.Is(err, model.ErrorNonExistName) {
		return model.Fio{}, model.ErrorNonExistName
	} else if err != nil {
		return model.Fio{}, model.ErrorApi
	} else {
		f.Gender = gender
	}

	// fill nation field
	if nation, err := apis.GetNation(f.Name); errors.Is(err, model.ErrorNonExistName) {
		return model.Fio{}, model.ErrorNonExistName
	} else if err != nil {
		return model.Fio{}, model.ErrorApi
	} else {
		f.Nation = nation
	}

	return a.FioRepo.AddFio(ctx, f)
}

func (a *app) AddFio(ctx context.Context, f model.Fio) (model.Fio, error) {
	if err := valid.FilledFio(f); err != nil { // check if fio is valid
		return model.Fio{}, err
	}

	return a.FioRepo.AddFio(ctx, f)
}

func (a *app) UpdateFio(ctx context.Context, id uint, f model.Fio) (model.Fio, error) {
	if err := valid.FilledFio(f); err != nil { // check if fio is valid
		return model.Fio{}, err
	}

	return a.FioRepo.UpdateFio(ctx, id, f)
}

/*func (a *app) GetFioById(ctx context.Context, id uint) (model.Fio, error) {
	// TODO: implement
	return a.FioRepo.GetFioById()
}

func (a *app) GetFioByFilter(ctx context.Context, f model.Filter) ([]model.Fio, error) {
	// TODO: implement
	return nil, nil
}

func (a *app) UpdateFio(ctx context.Context, id uint, f model.Fio) (model.Fio, error) {
	// TODO: implement
	return model.Fio{}, nil
}

func (a *app) DeleteFio(ctx context.Context, id uint) error {
	// TODO: implement
	return nil
}*/
