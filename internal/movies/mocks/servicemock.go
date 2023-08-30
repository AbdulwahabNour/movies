package mocks

import (
	"context"

	model "github.com/AbdulwahabNour/movies/internal/model/movie"
	"github.com/stretchr/testify/mock"
)

type MockService struct {
	mock.Mock
}

func (m *MockService) CreateMovie(ctx context.Context, movie *model.Movie) error {
	args := m.Called(ctx, movie)
	return args.Error(0)
}

func (m *MockService) GetMovie(ctx context.Context, id int64) (*model.Movie, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*model.Movie), args.Error(1)
}

func (m *MockService) ListMovies(ctx context.Context, query *model.MovieSearchQuery) ([]*model.Movie, error) {
	args := m.Called(ctx, query)
	return args.Get(0).([]*model.Movie), args.Error(1)
}

func (m *MockService) UpdateMovie(ctx context.Context, movie *model.Movie) error {
	args := m.Called(ctx, movie)
	return args.Error(0)
}

func (m *MockService) DeleteMovie(ctx context.Context, id int64) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
