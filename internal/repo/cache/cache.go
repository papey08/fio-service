package cache

import (
	"context"
	"encoding/json"
	"fio-service/internal/model"
	"fio-service/pkg/logger"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

// expiration is how long fio would stay in cache
const expiration = time.Minute * 30

type cachedFio struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Age        int    `json:"age"`
	Gender     string `json:"gender"`
	Nation     string `json:"nation"`
}

func fioToCachedFio(f model.Fio) cachedFio {
	return cachedFio{
		Name:       f.Name,
		Surname:    f.Surname,
		Patronymic: f.Patronymic,
		Age:        f.Age,
		Gender:     f.Gender,
		Nation:     f.Nation,
	}
}

func cachedFioToFio(f cachedFio, id int) model.Fio {
	return model.Fio{
		Id:         id,
		Name:       f.Name,
		Surname:    f.Surname,
		Patronymic: f.Patronymic,
		Age:        f.Age,
		Gender:     f.Gender,
		Nation:     f.Nation,
	}
}

type Repo struct {
	redis.Client
}

func (r *Repo) GetFioByKey(ctx context.Context, key int) (model.Fio, error) {
	fioData, err := r.Get(ctx, strconv.Itoa(key)).Result()
	if err == redis.Nil {
		logger.Info("cannot find fio with id %d in cache: not exist", key)
		return model.Fio{}, model.ErrorFioNotFound
	} else if err != nil {
		logger.Error("cannot find fio with id %d in cache: %s", key, err.Error())
		return model.Fio{}, model.ErrorFioRepo
	}

	var receivedFio cachedFio
	if err = json.Unmarshal([]byte(fioData), &receivedFio); err != nil {
		logger.Error("cannot find fio with id %d in cache: %s", key, err.Error())
		return model.Fio{}, model.ErrorFioRepo
	} else {
		logger.Info("find fio with id %d in cache", key)
		return cachedFioToFio(receivedFio, key), nil
	}
}

func (r *Repo) SetFioByKey(ctx context.Context, f model.Fio) (model.Fio, error) {
	cf := fioToCachedFio(f)
	data, _ := json.Marshal(cf)
	if err := r.Set(ctx, strconv.Itoa(f.Id), data, expiration).Err(); err != nil {
		logger.Error("cannot insert fio %s %s in cache: %s", f.Name, f.Surname, err.Error())
		return model.Fio{}, model.ErrorFioRepo
	}
	logger.Info("insert fio %s %s in cache", f.Name, f.Surname)
	return f, nil
}

func (r *Repo) DeleteFioByKey(ctx context.Context, key int) error {
	_, err := r.Del(ctx, strconv.Itoa(key)).Result()
	if err == redis.Nil || err == nil {
		return nil
	} else {
		logger.Error("cannot delete fio with id %d: %s", key, err.Error())
		return model.ErrorFioRepo
	}
}
