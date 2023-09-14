package consumer

import (
	"context"
	"encoding/json"
	"fio-service/internal/app"
	"fio-service/internal/model"
	"github.com/segmentio/kafka-go"
	"log"
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
			log.Println("gracefully disconnected from kafka")
			return
		default:
			msg, _ := f.Reader.ReadMessage(ctx)

			var fio comingFio
			err := json.Unmarshal(msg.Value, &fio)
			if err != nil {
				// TODO: log
			}

			_, err = f.App.FillFio(ctx, model.Fio{
				Name:       fio.Name,
				Surname:    fio.Surname,
				Patronymic: fio.Patronymic,
			})
			if err != nil {
				// TODO
			}
		}
	}
}
