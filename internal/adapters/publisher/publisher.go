package publisher

import (
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
	kafka.Conn
}

func (f *FioFailedTopic) SendFio(fio model.Fio, reason string) {
	fioToSend := failedFio{
		Name:       fio.Name,
		Surname:    fio.Surname,
		Patronymic: fio.Patronymic,
		Reason:     reason,
	}
	fioToSendData, _ := json.Marshal(fioToSend)
	_, _ = f.WriteMessages(kafka.Message{Value: fioToSendData})
}
