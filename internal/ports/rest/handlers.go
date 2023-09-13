package rest

import (
	"errors"
	"fio-service/internal/app"
	"fio-service/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func addFio(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqBody addFioRequest
		if err := c.BindJSON(&reqBody); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(err))
		}

		fio, err := a.AddFio(c, model.Fio{
			Name:       reqBody.Name,
			Surname:    reqBody.Surname,
			Patronymic: reqBody.Patronymic,
			Age:        reqBody.Age,
			Gender:     reqBody.Gender,
			Nation:     reqBody.Nation,
		})

		switch {
		case errors.Is(err, model.ErrorFioAlreadyExists):
			c.AbortWithStatusJSON(http.StatusConflict, errorResponse(err))
		case errors.Is(err, model.ErrorFioNoFields):
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(err))
		case errors.Is(err, model.ErrorFioInvalidFields):
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(err))
		case err == nil:
			c.JSON(http.StatusOK, fioSuccessResponse(fio))
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		}
	}
}

func getFioById(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(err))
		}

		fio, err := a.GetFioById(c, uint(id))

		switch {
		case errors.Is(err, model.ErrorFioNotFound):
			c.AbortWithStatusJSON(http.StatusNotFound, errorResponse(err))
		case err == nil:
			c.JSON(http.StatusOK, fioSuccessResponse(fio))
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		}
	}
}

func getFioByFilter(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqBody getFioByFilterRequest
		if err := c.BindJSON(&reqBody); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(err))
		}

		fios, err := a.GetFioByFilter(c, model.Filter{
			Offset:       reqBody.Offset,
			Limit:        reqBody.Limit,
			ByName:       reqBody.ByName,
			Name:         reqBody.Name,
			BySurname:    reqBody.BySurname,
			Surname:      reqBody.Surname,
			ByPatronymic: reqBody.ByPatronymic,
			Patronymic:   reqBody.Patronymic,
			ByAge:        reqBody.ByAge,
			Age:          reqBody.Age,
			ByGender:     reqBody.ByGender,
			Gender:       reqBody.Gender,
			ByNation:     reqBody.ByNation,
			Nation:       reqBody.Nation,
		})

		switch {
		case err == nil:
			c.JSON(http.StatusOK, fiosSuccessResponse(fios))
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		}
	}
}

func updateFio(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(err))
		}

		var reqBody updateFioRequest
		if err = c.BindJSON(&reqBody); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(err))
		}

		fio, err := a.UpdateFio(c, uint(id), model.Fio{
			Name:       reqBody.Name,
			Surname:    reqBody.Surname,
			Patronymic: reqBody.Patronymic,
			Age:        reqBody.Age,
			Gender:     reqBody.Gender,
			Nation:     reqBody.Nation,
		})

		switch {
		case errors.Is(err, model.ErrorFioNotFound):
			c.AbortWithStatusJSON(http.StatusNotFound, errorResponse(err))
		case errors.Is(err, model.ErrorFioNoFields):
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(err))
		case errors.Is(err, model.ErrorFioInvalidFields):
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(err))
		case err == nil:
			c.JSON(http.StatusOK, fioSuccessResponse(fio))
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		}
	}
}

func deleteFio(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(err))
		}

		err = a.DeleteFio(c, uint(id))

		switch {
		case errors.Is(err, model.ErrorFioNotFound):
			c.AbortWithStatusJSON(http.StatusNotFound, errorResponse(err))
		case err == nil:
			c.JSON(http.StatusOK, nil)
		default:
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(err))
		}
	}
}
