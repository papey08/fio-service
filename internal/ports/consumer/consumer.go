package consumer

import (
	"context"
	"encoding/json"
	"fio-service/internal/app"
	"fio-service/internal/model"
	"fio-service/pkg/logger"
	"github.com/segmentio/kafka-go"
)

type comingFio struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
}

type FioTopic struct {
	Reader *kafka.Reader
	app.App
}

func NewFioTopic(a app.App, addr string, topic string) FioTopic {
	return FioTopic{
		Reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers: []string{addr},
			Topic:   topic,
		}),
		App: a,
	}
}

func (f *FioTopic) ListenFio(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			logger.Info("disconnected from FIO")
			return
		default:
			msg, err := f.Reader.ReadMessage(ctx)
			if err != nil {
				logger.Error("cannot get message from FIO: %s", err.Error())
			}

			var fio comingFio
			err = json.Unmarshal(msg.Value, &fio)
			if err != nil {
				logger.Error("cannot get message from FIO: %s", err.Error())
			}

			logger.Info("adding fio by FIO topic: %s %s", fio.Name, fio.Surname)

			_, _ = f.App.FillFio(ctx, model.Fio{
				Name:       fio.Name,
				Surname:    fio.Surname,
				Patronymic: fio.Patronymic,
			})
		}
	}
}
