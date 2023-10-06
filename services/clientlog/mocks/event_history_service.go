// Code generated by mockery v2.22.1. DO NOT EDIT.

package mocks

import (
	context "context"

	client "go-micro.dev/v4/client"

	mock "github.com/stretchr/testify/mock"

	v0 "github.com/owncloud/ocis/v2/protogen/gen/ocis/services/eventhistory/v0"
)

// EventHistoryService is an autogenerated mock type for the EventHistoryService type
type EventHistoryService struct {
	mock.Mock
}

// GetEvents provides a mock function with given fields: ctx, in, opts
func (_m *EventHistoryService) GetEvents(ctx context.Context, in *v0.GetEventsRequest, opts ...client.CallOption) (*v0.GetEventsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *v0.GetEventsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v0.GetEventsRequest, ...client.CallOption) (*v0.GetEventsResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v0.GetEventsRequest, ...client.CallOption) *v0.GetEventsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v0.GetEventsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v0.GetEventsRequest, ...client.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEventsForUser provides a mock function with given fields: ctx, in, opts
func (_m *EventHistoryService) GetEventsForUser(ctx context.Context, in *v0.GetEventsForUserRequest, opts ...client.CallOption) (*v0.GetEventsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *v0.GetEventsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v0.GetEventsForUserRequest, ...client.CallOption) (*v0.GetEventsResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v0.GetEventsForUserRequest, ...client.CallOption) *v0.GetEventsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v0.GetEventsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v0.GetEventsForUserRequest, ...client.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewEventHistoryService interface {
	mock.TestingT
	Cleanup(func())
}

// NewEventHistoryService creates a new instance of EventHistoryService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewEventHistoryService(t mockConstructorTestingTNewEventHistoryService) *EventHistoryService {
	mock := &EventHistoryService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
