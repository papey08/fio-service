package app

import (
	"context"
	"errors"
	"fio-service/internal/app/valid"
	"fio-service/internal/model"
)

type app struct {
	FioRepo
	Publisher
	Apis
}

func NewApp(fr FioRepo, p Publisher, a Apis) App {
	return &app{
		FioRepo:   fr,
		Publisher: p,
		Apis:      a,
	}
}

func (a *app) FillFio(ctx context.Context, f model.Fio) (model.Fio, error) {
	if err := valid.NonFilledFio(f); err != nil { // check if name, surname or patronymic are valid
		a.SendFio(f, err.Error())
		return model.Fio{}, err
	}

	// fill age field
	if age, err := a.GetAge(f.Name); errors.Is(err, model.ErrorNonExistName) {
		a.SendFio(f, err.Error())
		return model.Fio{}, model.ErrorNonExistName
	} else if err != nil {
		return model.Fio{}, model.ErrorApi
	} else {
		f.Age = age
	}

	// fill gender field
	if gender, err := a.GetGender(f.Name); errors.Is(err, model.ErrorNonExistName) {
		a.SendFio(f, err.Error())
		return model.Fio{}, model.ErrorNonExistName
	} else if err != nil {
		return model.Fio{}, model.ErrorApi
	} else {
		f.Gender = gender
	}

	// fill nation field
	if nation, err := a.GetNation(f.Name); errors.Is(err, model.ErrorNonExistName) {
		a.SendFio(f, err.Error())
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

func (a *app) UpdateFio(ctx context.Context, id int, f model.Fio) (model.Fio, error) {
	if err := valid.FilledFio(f); err != nil { // check if fio is valid
		return model.Fio{}, err
	}

	return a.FioRepo.UpdateFio(ctx, id, f)
}
