package model

import "errors"

var FioNotFound = errors.New("could not find required fio")
var FioRepoError = errors.New("something wrong with fio repo")
var FioAlreadyExists = errors.New("required fio already exists")
var FioNoFields = errors.New("some necessary fields of fio are missing")
var FioInvalidFields = errors.New("some necessary fields are invalid")
