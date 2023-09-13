package rest

type addFioRequest struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Age        int    `json:"age"`
	Gender     string `json:"gender"`
	Nation     string `json:"nation"`
}

type getFioByFilterRequest struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`

	ByName bool   `json:"by_name"`
	Name   string `json:"name"`

	BySurname bool   `json:"by_surname"`
	Surname   string `json:"surname"`

	ByPatronymic bool   `json:"by_patronymic"`
	Patronymic   string `json:"patronymic"`

	ByAge bool `json:"by_age"`
	Age   uint `json:"age"`

	ByGender bool   `json:"by_gender"`
	Gender   string `json:"gender"`

	ByNation bool   `json:"by_nation"`
	Nation   string `json:"nation"`
}

type updateFioRequest struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Age        int    `json:"age"`
	Gender     string `json:"gender"`
	Nation     string `json:"nation"`
}
