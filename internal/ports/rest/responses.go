package rest

import (
	"fio-service/internal/model"
	"github.com/gin-gonic/gin"
)

type fioResponse struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Age        int    `json:"age"`
	Gender     string `json:"gender"`
	Nation     string `json:"nation"`
}

type fiosResponse []fioResponse

func fioSuccessResponse(f model.Fio) *gin.H {
	return &gin.H{
		"data": fioResponse{
			Id:         f.Id,
			Name:       f.Name,
			Surname:    f.Surname,
			Patronymic: f.Patronymic,
			Age:        f.Age,
			Gender:     f.Gender,
			Nation:     f.Nation,
		},
		"error": nil,
	}
}

func fiosSuccessResponse(fios []model.Fio) *gin.H {
	resp := make(fiosResponse, 0, len(fios))
	for _, f := range fios {
		resp = append(resp, fioResponse{
			Id:         f.Id,
			Name:       f.Name,
			Surname:    f.Surname,
			Patronymic: f.Patronymic,
			Age:        f.Age,
			Gender:     f.Gender,
			Nation:     f.Nation,
		})
	}
	return &gin.H{
		"data":  resp,
		"error": nil,
	}
}

func errorResponse(err error) *gin.H {
	return &gin.H{
		"data":  nil,
		"error": err.Error(),
	}
}
