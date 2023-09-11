package mocks

import (
	"context"

	model "github.com/AbdulwahabNour/movies/internal/model/permission"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) AddPermission(ctx context.Context, p *model.Permission) error {
	args := m.Called(ctx, p)
	return args.Error(0)
}

func (m *MockRepository) GetPermission(ctx context.Context, id int64) (*model.Permission, error) {
	args := m.Called(ctx, id)

	return args.Get(0).(*model.Permission), args.Error(1)
}
func (m *MockRepository) UpdatePermission(ctx context.Context, p *model.Permission) error {
	args := m.Called(ctx, p)
	return args.Error(0)
}

func (m *MockRepository) DeletePermission(ctx context.Context, id int64) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
func (m *MockRepository) UserPermissions(ctx context.Context, userId int64) ([]*model.Permission, error) {
	args := m.Called(ctx, userId)
	return args.Get(0).([]*model.Permission), args.Error(1)
}

func (m *MockRepository) AddUserPermissions(ctx context.Context, userId int64, codes ...string) error {
	args := m.Called(ctx, userId, codes)
	return args.Error(0)
}

func (m *MockRepository) DeleteUserPermission(ctx context.Context, userId int64, codes ...string) error {
	args := m.Called(ctx, userId, codes)
	return args.Error(0)
}
