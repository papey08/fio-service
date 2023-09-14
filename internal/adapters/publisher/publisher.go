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
	Conn *kafka.Conn
}

func NewFioFailedTopic(ctx context.Context, network string, addr string, topic string) (FioFailedTopic, error) {
	fioFailedConn, err := kafka.DialLeader(ctx, network, addr, topic, 0)
	if err != nil {
		return FioFailedTopic{}, err
	}
	return FioFailedTopic{
		Conn: fioFailedConn,
	}, nil
}

func (f *FioFailedTopic) SendFio(fio model.Fio, reason string) {
	fioToSend := failedFio{
		Name:       fio.Name,
		Surname:    fio.Surname,
		Patronymic: fio.Patronymic,
		Reason:     reason,
	}
	fioToSendData, _ := json.Marshal(fioToSend)
	_, _ = f.Conn.WriteMessages(kafka.Message{Value: fioToSendData})
}
