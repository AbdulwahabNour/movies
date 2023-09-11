package mocks

import (
	"context"

	model "github.com/AbdulwahabNour/movies/internal/model/permission"
	"github.com/stretchr/testify/mock"
)

type MockService struct {
	mock.Mock
}

func (m *MockService) AddPermission(ctx context.Context, p *model.Permission) error {
	args := m.Called(ctx, p)
	return args.Error(0)
}

func (m *MockService) GetPermission(ctx context.Context, id int64) (*model.Permission, error) {
	args := m.Called(ctx, id)

	return args.Get(0).(*model.Permission), args.Error(1)
}

func (m *MockService) UpdatePermission(ctx context.Context, p *model.Permission) error {
	args := m.Called(ctx, p)
	return args.Error(0)
}

func (m *MockService) DeletePermission(ctx context.Context, id int64) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockService) GetUserPermissions(ctx context.Context, userId int64) ([]*model.Permission, error) {
	args := m.Called(ctx, userId)
	return args.Get(0).([]*model.Permission), args.Error(1)
}

func (m *MockService) SetUserPermissions(ctx context.Context, userId int64, permissions ...string) error {
	args := m.Called(ctx, userId, permissions)
	return args.Error(0)
}
func (m *MockService) DeleteUserPermission(ctx context.Context, userId int64, permissions ...string) error {
	args := m.Called(ctx, userId, permissions)
	return args.Error(0)
}
