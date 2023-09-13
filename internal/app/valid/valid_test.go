package valid

import (
	"fio-service/internal/model"
	"github.com/go-playground/assert/v2"
	"testing"
)

type nonFilledFioTest struct {
	description string
	fio         model.Fio
	expected    error
}

func TestNonFilledFio(t *testing.T) {
	tests := []nonFilledFioTest{
		{
			description: "valid fio",
			fio: model.Fio{
				Name:       "Dmitriy",
				Surname:    "Ushakov",
				Patronymic: "Vasilevich",
			},
			expected: nil,
		},
		{
			description: "invalid fio without necessary field",
			fio: model.Fio{
				Name: "Dmitriy",
			},
			expected: model.ErrorFioNoFields,
		},
		{
			description: "invalid fio with invalid name",
			fio: model.Fio{
				Name:    "X Æ A-12",
				Surname: "Musk",
			},
			expected: model.ErrorFioInvalidFields,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			assert.Equal(t, NonFilledFio(test.fio), test.expected)
		})
	}
}

type filledFioTest struct {
	description string
	fio         model.Fio
	expected    error
}

func TestFilledFio(t *testing.T) {
	tests := []filledFioTest{
		{
			description: "valid fio",
			fio: model.Fio{
				Name:       "Dmitriy",
				Surname:    "Ushakov",
				Patronymic: "Vasilevich",
				Age:        42,
				Gender:     "male",
				Nation:     "UA",
			},
			expected: nil,
		},
		{
			description: "invalid fio",
			fio: model.Fio{
				Name:       "Dmitriy",
				Surname:    "Ushakov",
				Patronymic: "Vasilevich",
				Age:        -1,
				Gender:     "other",
				Nation:     "other",
			},
			expected: model.ErrorFioInvalidFields,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			assert.Equal(t, FilledFio(test.fio), test.expected)
		})
	}
}
