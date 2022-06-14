// Code generated by mockery v2.10.4. DO NOT EDIT.

package mocks

import (
	models "github.com/gojek/xp/management-service/models"
	pagination "github.com/gojek/xp/management-service/pagination"
	mock "github.com/stretchr/testify/mock"

	services "github.com/gojek/xp/management-service/services"
)

// ExperimentService is an autogenerated mock type for the ExperimentService type
type ExperimentService struct {
	mock.Mock
}

// CreateExperiment provides a mock function with given fields: settings, expData
func (_m *ExperimentService) CreateExperiment(settings models.Settings, expData services.CreateExperimentRequestBody) (*models.Experiment, error) {
	ret := _m.Called(settings, expData)

	var r0 *models.Experiment
	if rf, ok := ret.Get(0).(func(models.Settings, services.CreateExperimentRequestBody) *models.Experiment); ok {
		r0 = rf(settings, expData)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Experiment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.Settings, services.CreateExperimentRequestBody) error); ok {
		r1 = rf(settings, expData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DisableExperiment provides a mock function with given fields: projectId, experimentId
func (_m *ExperimentService) DisableExperiment(projectId int64, experimentId int64) error {
	ret := _m.Called(projectId, experimentId)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64, int64) error); ok {
		r0 = rf(projectId, experimentId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EnableExperiment provides a mock function with given fields: settings, experimentId
func (_m *ExperimentService) EnableExperiment(settings models.Settings, experimentId int64) error {
	ret := _m.Called(settings, experimentId)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Settings, int64) error); ok {
		r0 = rf(settings, experimentId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetDBRecord provides a mock function with given fields: projectId, experimentId
func (_m *ExperimentService) GetDBRecord(projectId models.ID, experimentId models.ID) (*models.Experiment, error) {
	ret := _m.Called(projectId, experimentId)

	var r0 *models.Experiment
	if rf, ok := ret.Get(0).(func(models.ID, models.ID) *models.Experiment); ok {
		r0 = rf(projectId, experimentId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Experiment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.ID, models.ID) error); ok {
		r1 = rf(projectId, experimentId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetExperiment provides a mock function with given fields: projectId, experimentId
func (_m *ExperimentService) GetExperiment(projectId int64, experimentId int64) (*models.Experiment, error) {
	ret := _m.Called(projectId, experimentId)

	var r0 *models.Experiment
	if rf, ok := ret.Get(0).(func(int64, int64) *models.Experiment); ok {
		r0 = rf(projectId, experimentId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Experiment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, int64) error); ok {
		r1 = rf(projectId, experimentId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListExperiments provides a mock function with given fields: projectId, params
func (_m *ExperimentService) ListExperiments(projectId int64, params services.ListExperimentsParams) ([]*models.Experiment, *pagination.Paging, error) {
	ret := _m.Called(projectId, params)

	var r0 []*models.Experiment
	if rf, ok := ret.Get(0).(func(int64, services.ListExperimentsParams) []*models.Experiment); ok {
		r0 = rf(projectId, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Experiment)
		}
	}

	var r1 *pagination.Paging
	if rf, ok := ret.Get(1).(func(int64, services.ListExperimentsParams) *pagination.Paging); ok {
		r1 = rf(projectId, params)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*pagination.Paging)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(int64, services.ListExperimentsParams) error); ok {
		r2 = rf(projectId, params)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// RunCustomValidation provides a mock function with given fields: experiment, settings, context, operationType
func (_m *ExperimentService) RunCustomValidation(experiment models.Experiment, settings models.Settings, context services.ValidationContext, operationType services.OperationType) error {
	ret := _m.Called(experiment, settings, context, operationType)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Experiment, models.Settings, services.ValidationContext, services.OperationType) error); ok {
		r0 = rf(experiment, settings, context, operationType)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateExperiment provides a mock function with given fields: settings, experimentId, expData
func (_m *ExperimentService) UpdateExperiment(settings models.Settings, experimentId int64, expData services.UpdateExperimentRequestBody) (*models.Experiment, error) {
	ret := _m.Called(settings, experimentId, expData)

	var r0 *models.Experiment
	if rf, ok := ret.Get(0).(func(models.Settings, int64, services.UpdateExperimentRequestBody) *models.Experiment); ok {
		r0 = rf(settings, experimentId, expData)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Experiment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.Settings, int64, services.UpdateExperimentRequestBody) error); ok {
		r1 = rf(settings, experimentId, expData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidatePairwiseExperimentOrthogonality provides a mock function with given fields: projectId, listExpParams, segmenters
func (_m *ExperimentService) ValidatePairwiseExperimentOrthogonality(projectId int64, listExpParams services.ListExperimentsParams, segmenters []string) error {
	ret := _m.Called(projectId, listExpParams, segmenters)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64, services.ListExperimentsParams, []string) error); ok {
		r0 = rf(projectId, listExpParams, segmenters)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
