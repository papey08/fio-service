package app

import (
	"context"
	"fio-service/internal/app/mocks"
	"fio-service/internal/model"
	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type appTestSuite struct {
	suite.Suite
	fioRepo   *mocks.FioRepo
	publisher *mocks.Publisher
	apis      *mocks.Apis
	service   App
}

func (s *appTestSuite) SetupSuite() {
	s.fioRepo = new(mocks.FioRepo)
	s.publisher = new(mocks.Publisher)
	s.apis = new(mocks.Apis)
	s.service = NewApp(s.fioRepo, s.publisher, s.apis)
}

type fillFioMocks struct {
	name string

	age    int
	ageErr error

	gender    string
	genderErr error

	nation    string
	nationErr error

	sendFio    model.Fio
	sendReason string
	sendErr    error

	fio    model.Fio
	gotFio model.Fio
	err    error
}

type fillFioTest struct {
	description string
	givenFio    model.Fio
	expectedFio model.Fio
	expectedErr error
}

func (s *appTestSuite) TestFillFio() {
	fillMocks := []fillFioMocks{
		{
			name: "Dmitriy",

			age:    42,
			ageErr: nil,

			gender:    "male",
			genderErr: nil,

			nation:    "UA",
			nationErr: nil,

			fio: model.Fio{
				Name:       "Dmitriy",
				Surname:    "Ushakov",
				Patronymic: "Vasilevich",
				Age:        42,
				Gender:     "male",
				Nation:     "UA",
			},

			gotFio: model.Fio{
				Id:         0,
				Name:       "Dmitriy",
				Surname:    "Ushakov",
				Patronymic: "Vasilevich",
				Age:        42,
				Gender:     "male",
				Nation:     "UA",
			},
			err: nil,
		},
		{
			sendFio: model.Fio{
				Name: "Dmitriy",
			},
			sendReason: model.ErrorFioNoFields.Error(),
			sendErr:    nil,
		},
		{
			sendFio: model.Fio{
				Name:    "X Æ A-12",
				Surname: "Musk",
			},
			sendReason: model.ErrorFioInvalidFields.Error(),
			sendErr:    nil,
		},
		{
			name:   "qwerty",
			age:    0,
			ageErr: model.ErrorNonExistName,
			sendFio: model.Fio{
				Name:    "qwerty",
				Surname: "asdfgh",
			},
			sendReason: model.ErrorNonExistName.Error(),
			sendErr:    nil,
		},
	}

	tests := []fillFioTest{
		{
			description: "correct filling of fio and adding it to database",
			givenFio: model.Fio{
				Name:       "Dmitriy",
				Surname:    "Ushakov",
				Patronymic: "Vasilevich",
			},
			expectedFio: model.Fio{
				Id:         0,
				Name:       "Dmitriy",
				Surname:    "Ushakov",
				Patronymic: "Vasilevich",
				Age:        42,
				Gender:     "male",
				Nation:     "UA",
			},
			expectedErr: nil,
		},
		{
			description: "sending fio without surname to FIO_FAILED",
			givenFio: model.Fio{
				Name:       "Dmitriy",
				Surname:    "",
				Patronymic: "",
			},
			expectedFio: model.Fio{},
			expectedErr: model.ErrorFioNoFields,
		},
		{
			description: "sending fio with invalid name to FIO_FAILED",
			givenFio: model.Fio{
				Name:    "X Æ A-12",
				Surname: "Musk",
			},
			expectedFio: model.Fio{},
			expectedErr: model.ErrorFioInvalidFields,
		},
		{
			description: "sending fio with non existing name and surname to FIO_FAILED",
			givenFio: model.Fio{
				Name:    "qwerty",
				Surname: "asdfgh",
			},
			expectedFio: model.Fio{},
			expectedErr: model.ErrorNonExistName,
		},
	}

	for _, m := range fillMocks {
		s.apis.On("GetAge", m.name).Return(m.age, m.ageErr).Once()
		s.apis.On("GetGender", m.name).Return(m.gender, m.genderErr).Once()
		s.apis.On("GetNation", m.name).Return(m.nation, m.nationErr).Once()
		s.publisher.On("SendFio", mock.Anything, m.sendFio, m.sendReason).Return(m.sendErr).Once()
		s.fioRepo.On("AddFio", mock.Anything, m.fio).Return(m.gotFio, m.err).Once()
	}

	for _, test := range tests {
		s.T().Run(test.description, func(t *testing.T) {
			fio, err := s.service.FillFio(context.Background(), test.givenFio)
			assert.Equal(s.T(), fio, test.expectedFio)
			assert.Equal(s.T(), err, test.expectedErr)
		})
	}
}

type addFioMock struct {
	fio    model.Fio
	gotFio model.Fio
	err    error
}

type addFioTest struct {
	description string
	givenFio    model.Fio
	expectedFio model.Fio
	expectedErr error
}

func (s *appTestSuite) TestAddFio() {
	addMocks := []addFioMock{
		{
			fio: model.Fio{
				Name:       "Dmitriy",
				Surname:    "Ushakov",
				Patronymic: "Vasilevich",
				Age:        42,
				Gender:     "male",
				Nation:     "UA",
			},
			gotFio: model.Fio{
				Id:         0,
				Name:       "Dmitriy",
				Surname:    "Ushakov",
				Patronymic: "Vasilevich",
				Age:        42,
				Gender:     "male",
				Nation:     "UA",
			},
			err: nil,
		},
		{
			fio: model.Fio{
				Name: "Dmitriy",
			},
			gotFio: model.Fio{},
			err:    model.ErrorFioNoFields,
		},
	}

	tests := []addFioTest{
		{
			description: "correct adding of the fio",
			givenFio: model.Fio{
				Name:       "Dmitriy",
				Surname:    "Ushakov",
				Patronymic: "Vasilevich",
				Age:        42,
				Gender:     "male",
				Nation:     "UA",
			},
			expectedFio: model.Fio{
				Id:         0,
				Name:       "Dmitriy",
				Surname:    "Ushakov",
				Patronymic: "Vasilevich",
				Age:        42,
				Gender:     "male",
				Nation:     "UA",
			},
			expectedErr: nil,
		},
		{
			description: "attempt of adding invalid fio",
			givenFio: model.Fio{
				Name: "Dmitriy",
			},
			expectedFio: model.Fio{},
			expectedErr: model.ErrorFioNoFields,
		},
	}

	for _, m := range addMocks {
		s.fioRepo.On("AddFio", mock.Anything, m.fio).Return(m.gotFio, m.err).Once()
	}

	for _, test := range tests {
		s.T().Run(test.description, func(t *testing.T) {
			fio, err := s.service.AddFio(context.Background(), test.givenFio)
			assert.Equal(s.T(), fio, test.expectedFio)
			assert.Equal(s.T(), err, test.expectedErr)
		})
	}
}

type updateFioMock struct {
	id     int
	fio    model.Fio
	gotFio model.Fio
	err    error
}

type updateFioTest struct {
	description string
	givenId     int
	givenFio    model.Fio
	expectedFio model.Fio
	expectedErr error
}

func (s *appTestSuite) TestUpdateFio() {
	updateMocks := []updateFioMock{
		{
			id: 0,
			fio: model.Fio{
				Name:       "Dmitriy",
				Surname:    "Ushakov",
				Patronymic: "Vasilevich",
				Age:        42,
				Gender:     "male",
				Nation:     "UA",
			},
			gotFio: model.Fio{
				Id:         0,
				Name:       "Dmitriy",
				Surname:    "Ushakov",
				Patronymic: "Vasilevich",
				Age:        42,
				Gender:     "male",
				Nation:     "UA",
			},
			err: nil,
		},
		{
			id: 0,
			fio: model.Fio{
				Name: "Dmitriy",
			},
			gotFio: model.Fio{},
			err:    model.ErrorFioNoFields,
		},
	}

	tests := []updateFioTest{
		{
			description: "correct updating of fio",
			givenId:     0,
			givenFio: model.Fio{
				Name:       "Dmitriy",
				Surname:    "Ushakov",
				Patronymic: "Vasilevich",
				Age:        42,
				Gender:     "male",
				Nation:     "UA",
			},
			expectedFio: model.Fio{
				Id:         0,
				Name:       "Dmitriy",
				Surname:    "Ushakov",
				Patronymic: "Vasilevich",
				Age:        42,
				Gender:     "male",
				Nation:     "UA",
			},
			expectedErr: nil,
		},
		{
			description: "attempt of updating fio name to invalid",
			givenId:     0,
			givenFio: model.Fio{
				Name: "Dmitriy",
			},
			expectedFio: model.Fio{},
			expectedErr: model.ErrorFioNoFields,
		},
	}

	for _, m := range updateMocks {
		s.fioRepo.On("UpdateFio", mock.Anything, m.id, m.fio).Return(m.gotFio, m.err).Once()
	}

	for _, test := range tests {
		s.T().Run(test.description, func(t *testing.T) {
			fio, err := s.service.UpdateFio(context.Background(), test.givenId, test.givenFio)
			assert.Equal(s.T(), fio, test.expectedFio)
			assert.Equal(s.T(), err, test.expectedErr)
		})
	}
}

func TestAppTestSuite(t *testing.T) {
	suite.Run(t, new(appTestSuite))
}
