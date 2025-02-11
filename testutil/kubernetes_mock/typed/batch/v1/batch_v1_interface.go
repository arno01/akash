// Code generated by mockery v2.5.1. DO NOT EDIT.

package kubernetes_mocks

import (
	mock "github.com/stretchr/testify/mock"
	rest "k8s.io/client-go/rest"

	v1 "k8s.io/client-go/kubernetes/typed/batch/v1"
)

// BatchV1Interface is an autogenerated mock type for the BatchV1Interface type
type BatchV1Interface struct {
	mock.Mock
}

// CronJobs provides a mock function with given fields: namespace
func (_m *BatchV1Interface) CronJobs(namespace string) v1.CronJobInterface {
	ret := _m.Called(namespace)

	var r0 v1.CronJobInterface
	if rf, ok := ret.Get(0).(func(string) v1.CronJobInterface); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1.CronJobInterface)
		}
	}

	return r0
}

// Jobs provides a mock function with given fields: namespace
func (_m *BatchV1Interface) Jobs(namespace string) v1.JobInterface {
	ret := _m.Called(namespace)

	var r0 v1.JobInterface
	if rf, ok := ret.Get(0).(func(string) v1.JobInterface); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1.JobInterface)
		}
	}

	return r0
}

// RESTClient provides a mock function with given fields:
func (_m *BatchV1Interface) RESTClient() rest.Interface {
	ret := _m.Called()

	var r0 rest.Interface
	if rf, ok := ret.Get(0).(func() rest.Interface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(rest.Interface)
		}
	}

	return r0
}
