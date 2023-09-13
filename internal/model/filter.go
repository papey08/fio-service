package model

type Filter struct {
	// Offset is how many fios to skip in database
	Offset int

	// Limit is how many fios to return
	Limit int

	// ByName is true when need to filter fios by Name
	ByName bool
	Name   string

	// BySurname is true when need to filter fios by Surname
	BySurname bool
	Surname   string

	// ByPatronymic is true when need to filter fios by Patronymic
	ByPatronymic bool
	Patronymic   string

	// ByAge is true when need to filter fios by Age
	ByAge bool
	Age   uint

	// ByGender is true when need to filter fios by Gender
	ByGender bool
	Gender   string

	// ByNation is true when need to filter fios by Nation
	ByNation bool
	Nation   string
}
