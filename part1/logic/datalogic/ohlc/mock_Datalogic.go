// Code generated by mockery v2.21.1. DO NOT EDIT.

package fs

import (
	context "context"

	model "github.com/ctopher7/gltc/v2/part1/model"
	mock "github.com/stretchr/testify/mock"
)

// MockDatalogic is an autogenerated mock type for the Datalogic type
type MockDatalogic struct {
	mock.Mock
}

// CalculateOhlc provides a mock function with given fields: dataset, stockName
func (_m *MockDatalogic) CalculateOhlc(dataset []model.Stock, stockName string) (model.Ohlc, []string, error) {
	ret := _m.Called(dataset, stockName)

	var r0 model.Ohlc
	var r1 []string
	var r2 error
	if rf, ok := ret.Get(0).(func([]model.Stock, string) (model.Ohlc, []string, error)); ok {
		return rf(dataset, stockName)
	}
	if rf, ok := ret.Get(0).(func([]model.Stock, string) model.Ohlc); ok {
		r0 = rf(dataset, stockName)
	} else {
		r0 = ret.Get(0).(model.Ohlc)
	}

	if rf, ok := ret.Get(1).(func([]model.Stock, string) []string); ok {
		r1 = rf(dataset, stockName)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]string)
		}
	}

	if rf, ok := ret.Get(2).(func([]model.Stock, string) error); ok {
		r2 = rf(dataset, stockName)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetOhlcFromCache provides a mock function with given fields: ctx, stockName
func (_m *MockDatalogic) GetOhlcFromCache(ctx context.Context, stockName string) (model.Ohlc, error) {
	ret := _m.Called(ctx, stockName)

	var r0 model.Ohlc
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (model.Ohlc, error)); ok {
		return rf(ctx, stockName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) model.Ohlc); ok {
		r0 = rf(ctx, stockName)
	} else {
		r0 = ret.Get(0).(model.Ohlc)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, stockName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PublishOhlc provides a mock function with given fields: ctx, data
func (_m *MockDatalogic) PublishOhlc(ctx context.Context, data []string) error {
	ret := _m.Called(ctx, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []string) error); ok {
		r0 = rf(ctx, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StoreOhlc provides a mock function with given fields: ctx, data
func (_m *MockDatalogic) StoreOhlc(ctx context.Context, data map[string]model.Ohlc) error {
	ret := _m.Called(ctx, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, map[string]model.Ohlc) error); ok {
		r0 = rf(ctx, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewMockDatalogic interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockDatalogic creates a new instance of MockDatalogic. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockDatalogic(t mockConstructorTestingTNewMockDatalogic) *MockDatalogic {
	mock := &MockDatalogic{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
