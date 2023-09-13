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
	if f.Name == "" || f.Surname == "" { // sending fio to FIO_FAILED if there are no of any necessary fields
		a.SendFio(f, model.ErrorFioNoFields.Error())
		return model.Fio{}, model.ErrorFioNoFields
	} else if !valid.Name(f.Name) || !valid.Name(f.Surname) || !valid.Name(f.Patronymic) { // sending fio to FIO_FAILED if some of the fields are invalid
		a.SendFio(f, model.ErrorFioInvalidFields.Error())
		return model.Fio{}, model.ErrorFioInvalidFields
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

	return a.AddFio(ctx, f)
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
