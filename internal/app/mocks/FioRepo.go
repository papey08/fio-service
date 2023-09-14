// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import mock "github.com/stretchr/testify/mock"
import model "fio-service/internal/model"

// FioRepo is an autogenerated mock type for the FioRepo type
type FioRepo struct {
	mock.Mock
}

// AddFio provides a mock function with given fields: ctx, f
func (_m *FioRepo) AddFio(ctx context.Context, f model.Fio) (model.Fio, error) {
	ret := _m.Called(ctx, f)

	var r0 model.Fio
	if rf, ok := ret.Get(0).(func(context.Context, model.Fio) model.Fio); ok {
		r0 = rf(ctx, f)
	} else {
		r0 = ret.Get(0).(model.Fio)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.Fio) error); ok {
		r1 = rf(ctx, f)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteFio provides a mock function with given fields: ctx, id
func (_m *FioRepo) DeleteFio(ctx context.Context, id int) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetFioByFilter provides a mock function with given fields: ctx, f
func (_m *FioRepo) GetFioByFilter(ctx context.Context, f model.Filter) ([]model.Fio, error) {
	ret := _m.Called(ctx, f)

	var r0 []model.Fio
	if rf, ok := ret.Get(0).(func(context.Context, model.Filter) []model.Fio); ok {
		r0 = rf(ctx, f)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Fio)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.Filter) error); ok {
		r1 = rf(ctx, f)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFioById provides a mock function with given fields: ctx, id
func (_m *FioRepo) GetFioById(ctx context.Context, id int) (model.Fio, error) {
	ret := _m.Called(ctx, id)

	var r0 model.Fio
	if rf, ok := ret.Get(0).(func(context.Context, int) model.Fio); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(model.Fio)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateFio provides a mock function with given fields: ctx, id, f
func (_m *FioRepo) UpdateFio(ctx context.Context, id int, f model.Fio) (model.Fio, error) {
	ret := _m.Called(ctx, id, f)

	var r0 model.Fio
	if rf, ok := ret.Get(0).(func(context.Context, int, model.Fio) model.Fio); ok {
		r0 = rf(ctx, id, f)
	} else {
		r0 = ret.Get(0).(model.Fio)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, model.Fio) error); ok {
		r1 = rf(ctx, id, f)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
