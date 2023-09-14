package consumer

import (
	"context"
	"encoding/json"
	"fio-service/internal/app"
	"fio-service/internal/model"
	"github.com/segmentio/kafka-go"
)

// maxBytes is a longest length of kafka message
const maxBytes = 1000

type comingFio struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
}

type FioTopic struct {
	Conn *kafka.Conn
	app.App
}

func NewFioTopic(ctx context.Context, a app.App, network string, addr string, topic string) (FioTopic, error) {
	fioConn, err := kafka.DialLeader(ctx, network, addr, topic, 0)
	if err != nil {
		return FioTopic{}, err
	}
	return FioTopic{
		Conn: fioConn,
		App:  a,
	}, err
}

func (f *FioTopic) ListenFio(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			msg, _ := f.Conn.ReadMessage(maxBytes)

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
