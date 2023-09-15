package apis

import (
	"encoding/json"
	"fio-service/internal/model"
	"fio-service/pkg/logger"
	"fmt"
	"io"
	"net/http"
)

const (
	ageUrl    = "https://api.agify.io/?name=%s"
	genderUrl = "https://api.genderize.io/?name=%s"
	nationUrl = "https://api.nationalize.io/?name=%s"
)

type Apis struct{}

type ageResponse struct {
	Age int `json:"age"`
}

func (a *Apis) GetAge(name string) (int, error) {
	url := fmt.Sprintf(ageUrl, name)
	resp, err := http.Get(url)
	if err != nil {
		logger.Error("cannot fill age of name %s: %s", name, err.Error())
		return 0, model.ErrorApi
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error("cannot fill age of name %s: %s", name, err.Error())
		return 0, model.ErrorApi
	}

	var ageResp ageResponse
	err = json.Unmarshal(body, &ageResp)
	if err != nil {
		logger.Error("cannot fill age of name %s: %s", name, err.Error())
		return 0, model.ErrorApi
	}

	if ageResp.Age == 0 {
		logger.Info("cannot fill age of name %s: name not exist", name)
		return 0, model.ErrorNonExistName
	}
	logger.Info("fill name %s with age %d", name, ageResp.Age)
	return ageResp.Age, nil
}

type genderResponse struct {
	Gender string `json:"gender"`
}

func (a *Apis) GetGender(name string) (string, error) {
	url := fmt.Sprintf(genderUrl, name)
	resp, err := http.Get(url)
	if err != nil {
		logger.Error("cannot fill gender of name %s: %s", name, err.Error())
		return "", model.ErrorApi
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error("cannot fill gender of name %s: %s", name, err.Error())
		return "", model.ErrorApi
	}

	var genderResp genderResponse
	err = json.Unmarshal(body, &genderResp)
	if err != nil {
		logger.Error("cannot fill gender of name %s: %s", name, err.Error())
		return "", model.ErrorApi
	}

	if genderResp.Gender == "" {
		logger.Info("cannot fill gender of name %s: name not exist", name)
		return "", model.ErrorNonExistName
	}
	logger.Info("fill name %s with gender %s", name, genderResp.Gender)
	return genderResp.Gender, nil
}

type country struct {
	CountryId string `json:"country_id"`
}

type nationResponse struct {
	Countries []country `json:"country"`
}

func (a *Apis) GetNation(name string) (string, error) {
	url := fmt.Sprintf(nationUrl, name)
	resp, err := http.Get(url)
	if err != nil {
		logger.Error("cannot fill nation of name %s: %s", name, err.Error())
		return "", model.ErrorApi
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error("cannot fill nation of name %s: %s", name, err.Error())
		return "", model.ErrorApi
	}

	var nationResp nationResponse
	err = json.Unmarshal(body, &nationResp)
	if err != nil {
		logger.Error("cannot fill nation of name %s: %s", name, err.Error())
		return "", model.ErrorApi
	}

	if len(nationResp.Countries) == 0 {
		logger.Info("cannot fill nation of name %s: name not exist", name)
		return "", model.ErrorNonExistName
	}
	logger.Info("fill name %s with nation %s", name, nationResp.Countries[0].CountryId)
	return nationResp.Countries[0].CountryId, nil
}
