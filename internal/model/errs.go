package model

import "errors"

var ErrorFioNotFound = errors.New("could not find required fio")
var ErrorFioRepo = errors.New("something wrong with fio repo")
var ErrorFioAlreadyExists = errors.New("required fio already exists")
var ErrorFioNoFields = errors.New("some necessary fields of fio are missing")
var ErrorFioInvalidFields = errors.New("some necessary fields are invalid")
var ErrorNonExistName = errors.New("name from fio not exist in api")
var ErrorApi = errors.New("something wrong with api")
var ErrorInvalidInput = errors.New("invalid input in request")
var ErrorSendingFio = errors.New("unable to send fio")
