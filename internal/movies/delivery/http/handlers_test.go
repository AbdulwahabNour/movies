package http

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"

	model "github.com/AbdulwahabNour/movies/internal/model/movie"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateMovieHandler(t *testing.T) {
	router, handlers, mockService := setupTest()
	router.POST("/movies", handlers.CreateMovieHandler)

	testCases := []struct {
		description        string
		mocking            bool
		payload            []byte
		movie              *model.Movie
		returnArguments    interface{}
		expectedStatusCode int
	}{
		{
			description:        "Invalid payload",
			mocking:            false,
			payload:            []byte("test"),
			movie:              nil,
			returnArguments:    nil,
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			description:        "Service error",
			mocking:            true,
			payload:            []byte(`{"title":"title", "year":2021}`),
			movie:              &model.Movie{Title: "title", Year: 2021},
			returnArguments:    fmt.Errorf("not found"),
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			description:        "Successful creation",
			mocking:            true,
			payload:            []byte(`{"title":"title", "year":2020}`),
			movie:              &model.Movie{Title: "title", Year: 2020},
			returnArguments:    nil,
			expectedStatusCode: http.StatusCreated,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {

			if tc.mocking {
				mockService.On("CreateMovie", mock.Anything, tc.movie).Return(tc.returnArguments)

			}

			w := httptest.NewRecorder()

			req, _ := http.NewRequest("POST", "/movies", bytes.NewBuffer(tc.payload))

			router.ServeHTTP(w, req)
			assert.Equal(t, tc.expectedStatusCode, w.Code)
		})

	}
	mockService.AssertExpectations(t)
}

func TestShowMovieHandler(t *testing.T) {
	router, handlers, mockService := setupTest()
	router.GET("/movies/:id", handlers.ShowMovieHandler)

	testCases := []struct {
		description        string
		mocking            bool
		query              string
		returnArguments    []interface{}
		expectedStatusCode int
	}{
		{
			description:        "Invalid ID format",
			mocking:            false,
			query:              "test",
			returnArguments:    nil,
			expectedStatusCode: http.StatusUnprocessableEntity,
		},

		{
			description:        "Service error",
			mocking:            true,
			query:              "11",
			returnArguments:    []interface{}{&model.Movie{}, fmt.Errorf("not found")},
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			description:        "Successful retrieval",
			mocking:            true,
			query:              "1",
			returnArguments:    []interface{}{&model.Movie{}, nil},
			expectedStatusCode: http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			if tc.mocking {
				id, _ := strconv.ParseInt(tc.query, 10, 64)
				mockService.On("GetMovie", mock.Anything, id).Return(tc.returnArguments...)
			}

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", fmt.Sprintf("/movies/%s", tc.query), nil)

			router.ServeHTTP(w, req)
			assert.Equal(t, tc.expectedStatusCode, w.Code)

		})
	}
	mockService.AssertExpectations(t)

}
func TestListMoviesHandler(t *testing.T) {
	router, handlers, mockService := setupTest()
	router.GET("/movies", handlers.ListMoviesHandler)

	testCases := []struct {
		description        string
		mocking            bool
		query              url.Values
		filter             *model.MovieSearchQuery
		returnArguments    []interface{}
		expectedStatusCode int
	}{
		{
			description:        "Invalid query parameters",
			mocking:            false,
			query:              url.Values{"invalid_param": []string{"test"}},
			filter:             nil,
			returnArguments:    nil,
			expectedStatusCode: http.StatusInternalServerError,
		},

		{
			description: "Service error",
			mocking:     true,
			query: url.Values{"title": []string{"test"},
				"page":      []string{"1"},
				"page_size": []string{"10"}},
			filter: &model.MovieSearchQuery{
				Title: "test",
				Filter: model.Filters{
					Page:     1,
					PageSize: 10,
				},
			},
			returnArguments:    []interface{}{[]*model.Movie{}, fmt.Errorf("not found")},
			expectedStatusCode: http.StatusInternalServerError,
		},

		{
			description: "Successful retrieval",
			mocking:     true,
			query:       url.Values{"page": []string{"2"}, "page_size": []string{"5"}},
			filter: &model.MovieSearchQuery{
				Filter: model.Filters{
					Page:     2,
					PageSize: 5,
				},
			},
			returnArguments:    []interface{}{[]*model.Movie{}, nil},
			expectedStatusCode: http.StatusOK,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			if tc.mocking {
				mockService.On("ListMovies", mock.Anything, tc.filter).Return(tc.returnArguments...)
			}

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/movies?"+tc.query.Encode(), nil)

			router.ServeHTTP(w, req)
			assert.Equal(t, tc.expectedStatusCode, w.Code)
		})
	}

	mockService.AssertExpectations(t)

}

func TestUpdateMovieHandler(t *testing.T) {
	router, handlers, mockService := setupTest()
	router.PUT("/movies/:id", handlers.UpdateMovieHandler)

	testCases := []struct {
		description        string
		mocking            bool
		query              string
		payload            []byte
		movie              *model.Movie
		returnArgument     error
		expectedStatusCode int
	}{
		{
			description:        "Invalid ID format",
			mocking:            false,
			query:              "test",
			payload:            nil,
			movie:              nil,
			returnArgument:     nil,
			expectedStatusCode: http.StatusUnprocessableEntity,
		},

		{
			description:        "Invalid payload",
			mocking:            false,
			query:              "1",
			payload:            nil,
			movie:              nil,
			returnArgument:     nil,
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			description:        "Service error",
			mocking:            true,
			query:              "2",
			payload:            []byte(`{"title":"error", "year":2020}`),
			movie:              &model.Movie{ID: 2, Title: "error", Year: 2020},
			returnArgument:     fmt.Errorf("not found"),
			expectedStatusCode: http.StatusInternalServerError,
		},

		{
			description:        "Successful update",
			mocking:            true,
			query:              "3",
			payload:            []byte(`{"title":"test", "year":2021}`),
			movie:              &model.Movie{ID: 3, Title: "test", Year: 2021},
			returnArgument:     nil,
			expectedStatusCode: http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			if tc.mocking {

				mockService.On("UpdateMovie", mock.Anything, tc.movie).Return(tc.returnArgument)
			}

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("PUT", fmt.Sprintf("/movies/%s", tc.query), bytes.NewBuffer(tc.payload))

			router.ServeHTTP(w, req)
			assert.Equal(t, tc.expectedStatusCode, w.Code)
		})
	}

	mockService.AssertExpectations(t)
}
func TestDeleteMovieHandler(t *testing.T) {
	router, handlers, mockService := setupTest()
	router.DELETE("/movies/:id", handlers.DeleteMovieHandler)

	testCases := []struct {
		description        string
		mocking            bool
		query              string
		returnArgument     error
		expectedStatusCode int
	}{
		{
			description:        "Invalid ID format",
			mocking:            false,
			query:              "test",
			returnArgument:     nil,
			expectedStatusCode: http.StatusUnprocessableEntity,
		},
		{
			description:        "Service error",
			mocking:            true,
			query:              "11",
			returnArgument:     fmt.Errorf("not found"),
			expectedStatusCode: http.StatusInternalServerError,
		},

		{
			description:        "Successful deletion",
			mocking:            true,
			query:              "12",
			returnArgument:     nil,
			expectedStatusCode: http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			if tc.mocking {
				id, _ := strconv.ParseInt(tc.query, 10, 64)
				mockService.On("DeleteMovie", mock.Anything, id).Return(tc.returnArgument)
			}

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("DELETE", fmt.Sprintf("/movies/%s", tc.query), nil)

			router.ServeHTTP(w, req)
			assert.Equal(t, tc.expectedStatusCode, w.Code)
		})
	}

	mockService.AssertExpectations(t)

}
