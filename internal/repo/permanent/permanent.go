package permanent

import (
	"context"
	"errors"
	"fio-service/internal/model"
	"fio-service/pkg/logger"
	"github.com/jackc/pgx/v5"
)

const (
	selectFioByIdQuery = `
        SELECT * FROM fios
        WHERE id = $1;`

	selectFioByFilterQuery = `
        SELECT * FROM fios
        WHERE (((NOT $1) OR name = $2)
            AND ((NOT $3) OR surname = $4)
            AND ((NOT $5) OR patronymic = $6)
            AND ((NOT $7) OR age = $8)
            AND ((NOT $9) OR gender = $10)
            AND ((NOT $11) OR nation = $12))
		LIMIT $13 OFFSET $14;`

	insertFioQuery = `
		INSERT INTO fios (name, surname, patronymic, age, gender, nation) 
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id;`

	updateFioQuery = `
        UPDATE fios
		SET name = $2,
		    surname = $3,
		    patronymic = $4,
		    age = $5,
		    gender = $6,
		    nation = $7
        WHERE id = $1;`

	deleteFioQuery = `
        DELETE FROM fios
        WHERE id = $1;`
)

type Repo struct {
	pgx.Conn
}

func (r *Repo) SelectFioById(ctx context.Context, id int) (model.Fio, error) {
	row := r.QueryRow(ctx, selectFioByIdQuery, id)
	var f model.Fio
	if err := row.Scan(&f.Id, &f.Name, &f.Surname, &f.Patronymic, &f.Age, &f.Gender, &f.Nation); errors.Is(err, pgx.ErrNoRows) {
		logger.Info("cannot find fio with id %d in storage: not exist", id)
		return model.Fio{}, model.ErrorFioNotFound
	} else if err != nil {
		logger.Error("cannot find fio with id %d in storage: %s", id, err.Error())
		return model.Fio{}, model.ErrorFioRepo
	} else {
		logger.Info("find fio with id %d in storage", id)
		return f, nil
	}
}

func (r *Repo) SelectFioByFilter(ctx context.Context, f model.Filter) ([]model.Fio, error) {
	rows, err := r.Query(ctx, selectFioByFilterQuery,
		f.ByName, f.Name,
		f.BySurname, f.Surname,
		f.ByPatronymic, f.Patronymic,
		f.ByAge, f.Age,
		f.ByGender, f.Gender,
		f.ByNation, f.Nation,
		f.Limit, f.Offset)
	if err != nil {
		logger.Error("cannot find fios with filter in storage: %s", err.Error())
		return nil, model.ErrorFioRepo
	}
	defer rows.Close()

	fios := make([]model.Fio, 0, f.Limit)
	for rows.Next() {
		var tempFio model.Fio
		_ = rows.Scan(&tempFio.Id, &tempFio.Name, &tempFio.Surname, &tempFio.Patronymic, &tempFio.Age, &tempFio.Gender, &tempFio.Nation)
		fios = append(fios, tempFio)
	}
	logger.Info("find fios with filter in storage")
	return fios, nil
}

func (r *Repo) InsertFio(ctx context.Context, f model.Fio) (model.Fio, error) {
	var insertedFioId int
	err := r.QueryRow(ctx, insertFioQuery,
		f.Name,
		f.Surname,
		f.Patronymic,
		f.Age,
		f.Gender,
		f.Nation).Scan(&insertedFioId)
	if err != nil {
		logger.Error("cannot insert fio %s %s in storage: %s", f.Name, f.Surname, err.Error())
		return model.Fio{}, model.ErrorFioRepo
	}
	f.Id = insertedFioId
	logger.Info("insert fio %s %s in storage with id %d", f.Name, f.Surname, insertedFioId)
	return f, nil
}

func (r *Repo) UpdateFio(ctx context.Context, id int, f model.Fio) (model.Fio, error) {
	e, err := r.Exec(ctx, updateFioQuery,
		id,
		f.Name,
		f.Surname,
		f.Patronymic,
		f.Age,
		f.Gender,
		f.Nation)
	if err != nil {
		logger.Error("cannot update fio with id %d in storage: %s", id, err.Error())
		return model.Fio{}, model.ErrorFioRepo
	} else if e.RowsAffected() == 0 {
		logger.Info("cannot update fio with id %d in storage: not exist", id)
		return model.Fio{}, model.ErrorFioNotFound
	} else {
		logger.Info("update fio with id %d in storage", id)
		return model.Fio{
			Id:         id,
			Name:       f.Name,
			Surname:    f.Surname,
			Patronymic: f.Patronymic,
			Age:        f.Age,
			Gender:     f.Gender,
			Nation:     f.Nation,
		}, nil
	}
}

func (r *Repo) DeleteFio(ctx context.Context, id int) error {
	e, err := r.Exec(ctx, deleteFioQuery, id)
	if err != nil {
		logger.Error("cannot delete fio with id %d from storage: %s", id, err.Error())
		return model.ErrorFioRepo
	} else if e.RowsAffected() == 0 {
		logger.Info("cannot delete fio with id %d from storage: not exist", id)
		return model.ErrorFioNotFound
	} else {
		logger.Info("delete fio with id %d from storage", id)
		return nil
	}
}
