package service

import (
	"context"
	"fmt"
	"testing"
	"time"

	model "github.com/AbdulwahabNour/movies/internal/model/movie"
	"github.com/AbdulwahabNour/movies/pkg/httpError"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateMovie(t *testing.T) {
	movieServ, mockRepo := setup_test()

	testCases := []struct {
		description     string
		mocking         bool
		movie           *model.Movie
		returnArguments error
		expectedErr     error
	}{
		{
			description:     "Validate Movie ",
			mocking:         false,
			movie:           &model.Movie{},
			returnArguments: nil,
			expectedErr:     checkMovie(&model.Movie{}),
		},
		{
			description:     "Validate Movie id ",
			mocking:         false,
			movie:           &model.Movie{ID: -10},
			returnArguments: nil,
			expectedErr:     httpError.NewBadRequestError("movie id less than 0"),
		},
		{
			description: "Validate Movie fields",
			mocking:     false,
			movie: &model.Movie{
				Title:   "test",
				Year:    1000,
				Runtime: 50,
				Genres:  []string{"comedy"},
			},
			returnArguments: nil,
			expectedErr: checkMovie(&model.Movie{
				Title:   "test",
				Year:    1000,
				Runtime: 50,
				Genres:  []string{"comedy"},
			}),
		},
		{
			description: "Error in Movie Creation",
			mocking:     true,
			movie: &model.Movie{
				Title:   "test title",
				Year:    2020,
				Runtime: 13,
				Genres:  []string{"comedy"},
			},
			returnArguments: fmt.Errorf("error happen during create movie try again later"),
			expectedErr:     httpError.NewInternalServerError("error happen during create movie try again later"),
		},
		{
			description: "success Movie Creation",
			mocking:     true,
			movie: &model.Movie{
				Title:   "success test",
				Year:    2020,
				Runtime: 13,
				Genres:  []string{"comedy"},
			},
			returnArguments: nil,
			expectedErr:     nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {

			ctx, cancle := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancle()

			if tc.mocking {
				mockRepo.On("CreateMovie", ctx, tc.movie).Return(tc.returnArguments)
			}

			err := movieServ.CreateMovie(ctx, tc.movie)

			assert.Equal(t, tc.expectedErr, err)

		})
	}
	mockRepo.AssertExpectations(t)
}

func TestGetMovie(t *testing.T) {
	movieServ, mockRepo := setup_test()
	type returnVals struct {
		movie *model.Movie
		err   error
	}
	testCases := []struct {
		description     string
		mocking         bool
		id              int64
		returnArguments []interface{}
		expectedReturn  returnVals
	}{

		{
			description:     "Invalid Movie id",
			mocking:         false,
			id:              0,
			returnArguments: nil,
			expectedReturn: returnVals{
				movie: nil,
				err:   httpError.NewNotFoundError("movie not found"),
			},
		},
		{
			description: "Repo Error",
			mocking:     true,
			id:          100,
			returnArguments: []interface{}{
				&model.Movie{},
				fmt.Errorf("not found"),
			},
			expectedReturn: returnVals{
				movie: nil,
				err:   httpError.NewInternalServerError("not found"),
			},
		},
		{
			description: "success GetMovie",
			mocking:     true,
			id:          10,
			returnArguments: []interface{}{
				&model.Movie{},
				nil,
			},
			expectedReturn: returnVals{
				movie: &model.Movie{},
				err:   nil,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {

			ctx, cancle := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancle()

			if tc.mocking {

				mockRepo.On("GetMovie", ctx, tc.id).Return(tc.returnArguments...)
			}

			movie, err := movieServ.GetMovie(ctx, tc.id)

			assert.Equal(t, tc.expectedReturn.err, err)
			assert.Equal(t, tc.expectedReturn.movie, movie)

		})
	}
	mockRepo.AssertExpectations(t)

}

func TestListMovies(t *testing.T) {
	movieServ, mockRepo := setup_test()

	validate := validator.New()
	type returnVals struct {
		movie []*model.Movie
		err   error
	}
	testCases := []struct {
		description     string
		mocking         bool
		query           *model.MovieSearchQuery
		returnArguments []interface{}
		expectedReturn  returnVals
	}{

		{
			description:     "Invalid query",
			mocking:         false,
			query:           &model.MovieSearchQuery{},
			returnArguments: nil,
			expectedReturn: returnVals{
				movie: nil,
				err:   httpError.ParseValidationErrors(validate.Struct(&model.MovieSearchQuery{})),
			},
		},
		{
			description: "Repo Error",
			mocking:     true,
			query: &model.MovieSearchQuery{
				Filter: model.Filters{
					Page:     10,
					PageSize: 20,
					Sort:     "id",
				},
			},
			returnArguments: []interface{}{
				[]*model.Movie{},
				fmt.Errorf("something went wrong"),
			},
			expectedReturn: returnVals{
				movie: nil,
				err:   httpError.NewInternalServerError("something went wrong"),
			},
		},
		{
			description: "Success ListMovies",
			mocking:     true,
			query: &model.MovieSearchQuery{
				Filter: model.Filters{
					Page:     1,
					PageSize: 10,
					Sort:     "id",
				},
			},
			returnArguments: []interface{}{
				[]*model.Movie{},
				nil,
			},
			expectedReturn: returnVals{
				movie: []*model.Movie{},
				err:   nil,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {

			ctx, cancle := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancle()

			if tc.mocking {

				mockRepo.On("ListMovies", ctx, tc.query).Return(tc.returnArguments...)
			}

			movie, err := movieServ.ListMovies(ctx, tc.query)

			assert.Equal(t, tc.expectedReturn.err, err)
			assert.Equal(t, tc.expectedReturn.movie, movie)

		})
	}
	mockRepo.AssertExpectations(t)
}

func TestUpdateMovie(t *testing.T) {
	movieServ, mockRepo := setup_test()
	type mockingfunc struct {
		funcName        string
		funcpram        interface{}
		returnArguments []interface{}
	}
	testCases := []struct {
		description string
		mocking     bool
		mockingfunc []*mockingfunc
		movie       *model.Movie
		expectedErr error
	}{
		{
			description: "Movie Is Empty",
			mocking:     false,
			mockingfunc: nil,
			movie:       &model.Movie{},
			expectedErr: httpError.NewBadRequestError("The JSON payload is empty. Please provide valid data to update the movie."),
		},
		{
			description: "GetMovie",
			mocking:     true,
			mockingfunc: []*mockingfunc{
				{
					funcName: "GetMovie",
					funcpram: int64(12),
					returnArguments: []interface{}{
						&model.Movie{},
						fmt.Errorf("not found"),
					},
				},
			},
			movie: &model.Movie{ID: 12, Title: "test"},

			expectedErr: httpError.NewInternalServerError("not found"),
		},
		{
			description: "checkMovie",
			mocking:     true,
			mockingfunc: []*mockingfunc{
				{
					funcName: "GetMovie",
					funcpram: int64(10),
					returnArguments: []interface{}{
						&model.Movie{
							ID: -100,
						},
						nil,
					},
				},
			},
			movie: &model.Movie{ID: 10, Title: "test"},

			expectedErr: checkMovie(&model.Movie{ID: -100}),
		},
		{
			description: "repo UpdateMovie",
			mocking:     true,
			mockingfunc: []*mockingfunc{
				{
					funcName: "GetMovie",
					funcpram: int64(200),
					returnArguments: []interface{}{
						&model.Movie{
							ID:      200,
							Title:   "new title",
							Year:    2020,
							Runtime: 65,
							Genres:  []string{"comedy"},
						},
						nil,
					},
				},
				{
					funcName: "UpdateMovie",
					funcpram: &model.Movie{
						ID:      200,
						Title:   "new title",
						Year:    2020,
						Runtime: 65,
						Genres:  []string{"comedy"},
					},
					returnArguments: []interface{}{
						fmt.Errorf("something went wrong")},
				},
			},
			movie: &model.Movie{ID: 200, Title: "new title"},

			expectedErr: httpError.NewInternalServerError("something went wrong"),
		},
		{
			description: "Success UpdateMovie",
			mocking:     true,
			mockingfunc: []*mockingfunc{
				{
					funcName: "GetMovie",
					funcpram: int64(15),
					returnArguments: []interface{}{
						&model.Movie{
							ID:      15,
							Title:   "title",
							Year:    2013,
							Runtime: 45,
							Genres:  []string{"comedy"},
						},
						nil,
					},
				},
				{
					funcName: "UpdateMovie",
					funcpram: &model.Movie{
						ID:      15,
						Title:   "title",
						Year:    2013,
						Runtime: 45,
						Genres:  []string{"comedy"},
					},
					returnArguments: []interface{}{nil},
				},
			},
			movie: &model.Movie{ID: 15, Title: "title"},

			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {

			ctx, cancle := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancle()

			if tc.mocking {
				for _, v := range tc.mockingfunc {
					mockRepo.On(v.funcName, mock.Anything, v.funcpram).Return(v.returnArguments...)
				}
			}

			err := movieServ.UpdateMovie(ctx, tc.movie)

			assert.Equal(t, tc.expectedErr, err)

		})
	}
	mockRepo.AssertExpectations(t)
}
func TestDeleteMovie(t *testing.T) {
	movieServ, mockRepo := setup_test()

	testCases := []struct {
		description     string
		mocking         bool
		id              int64
		returnArguments error
		expectedReturn  error
	}{

		{
			description:     "Invalid Movie id",
			mocking:         false,
			id:              -100,
			returnArguments: nil,
			expectedReturn:  httpError.NewBadRequestError("movie id less than 1"),
		},
		{
			description:     "Repo DeleteMovie",
			mocking:         true,
			id:              10,
			returnArguments: fmt.Errorf("something went wrong"),
			expectedReturn:  httpError.NewInternalServerError("something went wrong"),
		},
		{
			description:     "Success DeleteMovie",
			mocking:         true,
			id:              2,
			returnArguments: nil,
			expectedReturn:  nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {

			ctx, cancle := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancle()

			if tc.mocking {

				mockRepo.On("DeleteMovie", ctx, tc.id).Return(tc.returnArguments)
			}

			err := movieServ.DeleteMovie(ctx, tc.id)

			assert.Equal(t, tc.expectedReturn, err)

		})
	}
	mockRepo.AssertExpectations(t)

}
