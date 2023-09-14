package apis

import (
	"encoding/json"
	"fio-service/internal/model"
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
		return 0, model.ErrorApi
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, model.ErrorApi
	}

	var ageResp ageResponse
	err = json.Unmarshal(body, &ageResp)
	if err != nil {
		return 0, model.ErrorApi
	}

	if ageResp.Age == 0 {
		return 0, model.ErrorNonExistName
	}
	return ageResp.Age, nil
}

type genderResponse struct {
	Gender string `json:"gender"`
}

func (a *Apis) GetGender(name string) (string, error) {
	url := fmt.Sprintf(genderUrl, name)
	resp, err := http.Get(url)
	if err != nil {
		return "", model.ErrorApi
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", model.ErrorApi
	}

	var genderResp genderResponse
	err = json.Unmarshal(body, &genderResp)
	if err != nil {
		return "", model.ErrorApi
	}

	if genderResp.Gender == "" {
		return "", model.ErrorNonExistName
	}
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
		return "", model.ErrorApi
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", model.ErrorApi
	}

	var nationResp nationResponse
	err = json.Unmarshal(body, &nationResp)
	if err != nil {
		return "", model.ErrorApi
	}

	if len(nationResp.Countries) == 0 {
		return "", model.ErrorNonExistName
	}
	return nationResp.Countries[0].CountryId, nil
}
