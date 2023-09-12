package model

type Filter struct {
	ByName bool
	Name   string

	BySurname bool
	Surname   string

	ByPatronymic bool
	Patronymic   string

	ByAge bool
	Age   uint

	ByGender bool
	Gender   string

	ByNation bool
	Nation   string
}
