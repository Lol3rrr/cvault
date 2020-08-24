package mocks

import (
	"github.com/hashicorp/vault/api"
	"github.com/stretchr/testify/mock"
)

// MockSession is a simple mock struct that can be used for testing
type MockSession struct {
	mock.Mock
}

// Auth is needed to satisfy the interface
func (m *MockSession) Auth() error {
	args := m.Mock.Called()
	return args.Error(0)
}

// ReadMap is needed to satisfy the interface
func (m *MockSession) ReadMap(name string) (map[string]interface{}, error) {
	args := m.Mock.Called(name)
	return args.Get(0).(map[string]interface{}), args.Error(1)
}

// ReadData is needed to satisfy the interface
func (m *MockSession) ReadData(name string) (*api.Secret, error) {
	args := m.Mock.Called(name)
	return args.Get(0).(*api.Secret), args.Error(1)
}

// Renew is needed to satisfy the interface
func (m *MockSession) Renew(id string, interval int) (*api.Secret, error) {
	args := m.Mock.Called(id, interval)
	return args.Get(0).(*api.Secret), args.Error(1)
}

// RenewDataForever is needed to satisfy the interface
func (m *MockSession) RenewDataForever(secret *api.Secret, interval int, callback func()) {
	m.Mock.Called(secret, interval, callback)
	return
}

// WriteMapData is needed to satisfy the interface
func (m *MockSession) WriteMapData(name string, data map[string]interface{}) error {
	args := m.Mock.Called(name, data)
	return args.Error(0)
}

// GetVaultClient is needed to satisfy the interface
func (m *MockSession) GetVaultClient() *api.Client {
	args := m.Mock.Called()
	return args.Get(0).(*api.Client)
}
