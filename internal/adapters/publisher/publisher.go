package publisher

import (
	"context"
	"encoding/json"
	"fio-service/internal/model"
	"github.com/segmentio/kafka-go"
)

type failedFio struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Reason     string `json:"reason"`
}

type FioFailedTopic struct {
	Writer *kafka.Writer
}

func NewFioFailedTopic(addr string, topic string) FioFailedTopic {
	return FioFailedTopic{
		Writer: &kafka.Writer{
			Addr:  kafka.TCP(addr),
			Topic: topic,
		},
	}
}

func (f *FioFailedTopic) SendFio(ctx context.Context, fio model.Fio, reason string) error {
	fioToSend := failedFio{
		Name:       fio.Name,
		Surname:    fio.Surname,
		Patronymic: fio.Patronymic,
		Reason:     reason,
	}
	fioToSendData, _ := json.Marshal(fioToSend)
	return f.Writer.WriteMessages(ctx, kafka.Message{Value: fioToSendData})
}
