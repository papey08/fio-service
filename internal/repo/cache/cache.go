package cache

import (
	"context"
	"encoding/json"
	"fio-service/internal/model"
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

func cachedFioToFio(f cachedFio, id uint) model.Fio {
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

func (r *Repo) GetFioByKey(ctx context.Context, key uint) (model.Fio, error) {
	fioData, err := r.Get(ctx, strconv.Itoa(int(key))).Result()
	if err == redis.Nil {
		return model.Fio{}, model.ErrorFioNotFound
	} else if err != nil {
		return model.Fio{}, model.ErrorFioRepo
	}

	var receivedFio cachedFio
	if err = json.Unmarshal([]byte(fioData), &receivedFio); err != nil {
		return model.Fio{}, model.ErrorFioRepo
	} else {
		return cachedFioToFio(receivedFio, key), nil
	}
}

func (r *Repo) SetFioByKey(ctx context.Context, f model.Fio) (model.Fio, error) {
	cf := fioToCachedFio(f)
	data, _ := json.Marshal(cf)
	if err := r.Set(ctx, strconv.Itoa(int(f.Id)), data, expiration).Err(); err != nil {
		return model.Fio{}, model.ErrorFioRepo
	}
	return f, nil
}

func (r *Repo) DeleteFioByKey(ctx context.Context, key uint) error {
	_, err := r.Del(ctx, strconv.Itoa(int(key))).Result()
	if err == redis.Nil || err == nil {
		return nil
	} else {
		return model.ErrorFioRepo
	}
}
