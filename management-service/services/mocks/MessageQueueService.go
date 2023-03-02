// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	pubsub "github.com/caraml-dev/xp/common/pubsub"
	mock "github.com/stretchr/testify/mock"

	segmenters "github.com/caraml-dev/xp/common/segmenters"
)

// MessageQueueService is an autogenerated mock type for the MessageQueueService type
type MessageQueueService struct {
	mock.Mock
}

// PublishExperimentMessage provides a mock function with given fields: updateType, experiment
func (_m *MessageQueueService) PublishExperimentMessage(updateType string, experiment *pubsub.Experiment) error {
	ret := _m.Called(updateType, experiment)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *pubsub.Experiment) error); ok {
		r0 = rf(updateType, experiment)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PublishProjectSegmenterMessage provides a mock function with given fields: updateType, segmenter, projectId
func (_m *MessageQueueService) PublishProjectSegmenterMessage(updateType string, segmenter *segmenters.SegmenterConfiguration, projectId int64) error {
	ret := _m.Called(updateType, segmenter, projectId)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *segmenters.SegmenterConfiguration, int64) error); ok {
		r0 = rf(updateType, segmenter, projectId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PublishProjectSettingsMessage provides a mock function with given fields: updateType, settings
func (_m *MessageQueueService) PublishProjectSettingsMessage(updateType string, settings *pubsub.ProjectSettings) error {
	ret := _m.Called(updateType, settings)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *pubsub.ProjectSettings) error); ok {
		r0 = rf(updateType, settings)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewMessageQueueService interface {
	mock.TestingT
	Cleanup(func())
}

// NewMessageQueueService creates a new instance of MessageQueueService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMessageQueueService(t mockConstructorTestingTNewMessageQueueService) *MessageQueueService {
	mock := &MessageQueueService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
