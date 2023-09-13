package valid

import "fio-service/internal/model"

const allowedNameSymbols = " abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-'"

func isAllowed(c rune) bool {
	for _, s := range allowedNameSymbols {
		if s == c {
			return true
		}
	}
	return false
}

func nameSurname(n string) error {
	if len(n) == 0 {
		return model.ErrorFioNoFields
	}
	for _, c := range n {
		if !isAllowed(c) {
			return model.ErrorFioInvalidFields
		}
	}
	return nil
}

func patronymic(p string) error {
	for _, c := range p {
		if !isAllowed(c) {
			return model.ErrorFioInvalidFields
		}
	}
	return nil
}

func age(a int) bool {
	return a >= 0
}

func gender(g string) bool {
	return g == "male" || g == "female"
}

func nation(n string) bool {
	return len(n) == 2
}

func NonFilledFio(f model.Fio) error {
	if err := nameSurname(f.Name); err != nil {
		return err
	}
	if err := nameSurname(f.Surname); err != nil {
		return err
	}
	if err := patronymic(f.Patronymic); err != nil {
		return err
	}
	return nil
}

func FilledFio(f model.Fio) error {
	if err := NonFilledFio(f); err != nil {
		return err
	}
	if !age(f.Age) || !gender(f.Gender) || !nation(f.Nation) {
		return model.ErrorFioInvalidFields
	}
	return nil
}
