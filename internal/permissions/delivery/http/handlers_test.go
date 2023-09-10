package http

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	model "github.com/AbdulwahabNour/movies/internal/model/permission"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddPermissionHandler(t *testing.T) {

	router, handlers, mockService := setupTest()
	router.POST("/permissions", handlers.AddPermissionHandler)

	testCases := []struct {
		description        string
		mocking            bool
		userPermissions    *model.Permission
		payload            []byte
		returnArguments    error
		expectedStatusCode int
	}{
		{
			description:        "Invalid payload",
			mocking:            false,
			userPermissions:    nil,
			payload:            nil,
			returnArguments:    nil,
			expectedStatusCode: http.StatusBadRequest,
		},

		{
			description:        "AddPermission  error",
			mocking:            true,
			userPermissions:    &model.Permission{Code: "test"},
			payload:            []byte(`{"code":"test"}`),
			returnArguments:    fmt.Errorf("something went wrong"),
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			description:        "success AddPermission",
			mocking:            true,
			userPermissions:    &model.Permission{Code: "add:movie"},
			payload:            []byte(`{"code":"add:movie"}`),
			returnArguments:    nil,
			expectedStatusCode: http.StatusCreated,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			if tc.mocking {
				mockService.On("AddPermission", mock.Anything, tc.userPermissions).
					Return(tc.returnArguments)
			}
			w := httptest.NewRecorder()

			req, _ := http.NewRequest("POST", "/permissions", bytes.NewBuffer(tc.payload))

			router.ServeHTTP(w, req)
			assert.Equal(t, tc.expectedStatusCode, w.Code)
		})
	}

	mockService.AssertExpectations(t)
}

func TestGetPermissioHandler(t *testing.T) {

	router, handlers, mockService := setupTest()
	router.GET("/:id", handlers.GetPermissioHandler)

	testCases := []struct {
		description        string
		mocking            bool
		query              string
		returnArguments    []interface{}
		expectedStatusCode int
	}{
		{
			description:        "Invalid id",
			mocking:            false,
			query:              "test",
			returnArguments:    nil,
			expectedStatusCode: http.StatusUnprocessableEntity,
		},

		{
			description: "GetPermission error",
			mocking:     true,
			query:       "10",
			returnArguments: []interface{}{
				&model.Permission{},
				fmt.Errorf("something went wrong"),
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			description: "success GetPermission",
			mocking:     true,
			query:       "11",
			returnArguments: []interface{}{
				&model.Permission{},
				nil,
			},
			expectedStatusCode: http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			if tc.mocking {
				pid, _ := strconv.ParseInt(tc.query, 10, 64)
				mockService.On("GetPermission", mock.Anything, pid).Return(tc.returnArguments...)
			}

			w := httptest.NewRecorder()

			req, _ := http.NewRequest("GET", fmt.Sprintf("/%s", tc.query), nil)

			router.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedStatusCode, w.Code)

		})
	}
	mockService.AssertExpectations(t)

}

func TestUpdatePermission(t *testing.T) {

	router, handlers, mockService := setupTest()
	router.PUT("/:id", handlers.UpdatePermissionHandler)

	testCases := []struct {
		description        string
		mocking            bool
		query              string
		payload            []byte
		permission         *model.Permission
		returnArgument     error
		expectedStatusCode int
	}{
		{
			description:        "Invalid id",
			mocking:            false,
			query:              "test",
			payload:            nil,
			permission:         nil,
			returnArgument:     nil,
			expectedStatusCode: http.StatusUnprocessableEntity,
		},
		{
			description:        "Invalid payload",
			mocking:            false,
			query:              "10",
			payload:            nil,
			permission:         nil,
			returnArgument:     nil,
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			description:        "UpdatePermission error",
			mocking:            true,
			query:              "10",
			payload:            []byte(`{"code":"addmovie"}`),
			permission:         &model.Permission{ID: 10, Code: "addmovie"},
			returnArgument:     fmt.Errorf("something went wrong"),
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			description:        "Success UpdatePermission ",
			mocking:            true,
			query:              "1",
			payload:            []byte(`{"code":"delete:movie"}`),
			permission:         &model.Permission{ID: 1, Code: "delete:movie"},
			returnArgument:     nil,
			expectedStatusCode: http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			if tc.mocking {
				mockService.On("UpdatePermission", mock.Anything, tc.permission).Return(tc.returnArgument)
			}

			w := httptest.NewRecorder()

			req, _ := http.NewRequest("PUT", fmt.Sprintf("/%s", tc.query), bytes.NewBuffer(tc.payload))

			router.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedStatusCode, w.Code)

		})
	}
	mockService.AssertExpectations(t)

}

func TestDeletePermissionHandler(t *testing.T) {
	router, handlers, mockService := setupTest()
	router.DELETE("/:id", handlers.DeletePermissionHandler)

	testCases := []struct {
		description        string
		mocking            bool
		query              string
		returnArgument     error
		expectedStatusCode int
	}{
		{
			description:        "Invalid id",
			mocking:            false,
			query:              "test",
			returnArgument:     nil,
			expectedStatusCode: http.StatusUnprocessableEntity,
		},

		{
			description:        "DeletePermission error",
			mocking:            true,
			query:              "10",
			returnArgument:     fmt.Errorf("something went wrong"),
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			description:        "success DeletePermission",
			mocking:            true,
			query:              "1",
			returnArgument:     nil,
			expectedStatusCode: http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			if tc.mocking {
				pid, _ := strconv.ParseInt(tc.query, 10, 64)
				mockService.On("DeletePermission", mock.Anything, pid).Return(tc.returnArgument)
			}

			w := httptest.NewRecorder()

			req, _ := http.NewRequest("DELETE", fmt.Sprintf("/%s", tc.query), nil)

			router.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedStatusCode, w.Code)

		})
	}
	mockService.AssertExpectations(t)
}

func TestGetUserPermissionsHandler(t *testing.T) {
	router, handlers, mockService := setupTest()
	router.GET("/user/:id", handlers.GetUserPermissionsHandler)

	testCases := []struct {
		description        string
		mocking            bool
		query              string
		returnArguments    []interface{}
		expectedStatusCode int
	}{
		{
			description:        "Invalid id",
			mocking:            false,
			query:              "test",
			returnArguments:    nil,
			expectedStatusCode: http.StatusUnprocessableEntity,
		},

		{
			description: "GetUserPermissions error",
			mocking:     true,
			query:       "10",
			returnArguments: []interface{}{
				[]*model.Permission{},
				fmt.Errorf("something went wrong"),
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			description: "Success GetUserPermissions",
			mocking:     true,
			query:       "8",
			returnArguments: []interface{}{
				[]*model.Permission{},
				nil,
			},
			expectedStatusCode: http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			if tc.mocking {
				pid, _ := strconv.ParseInt(tc.query, 10, 64)
				mockService.On("GetUserPermissions", mock.Anything, pid).Return(tc.returnArguments...)
			}

			w := httptest.NewRecorder()

			req, _ := http.NewRequest("GET", fmt.Sprintf("/user/%s", tc.query), nil)

			router.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedStatusCode, w.Code)

		})
	}
	mockService.AssertExpectations(t)
}

func TestSetUserPermissionHandler(t *testing.T) {
	router, handlers, mockService := setupTest()
	router.POST("/user/:id", handlers.SetUserPermissionHandler)

	testCases := []struct {
		description        string
		mocking            bool
		query              string
		userPermissions    []string
		payload            []byte
		returnArgument     error
		expectedStatusCode int
	}{
		{
			description:        "Invalid id",
			mocking:            false,
			query:              "test",
			userPermissions:    nil,
			payload:            nil,
			returnArgument:     nil,
			expectedStatusCode: http.StatusUnprocessableEntity,
		},
		{
			description:        "Invalid payload",
			mocking:            false,
			query:              "11",
			userPermissions:    nil,
			payload:            nil,
			returnArgument:     nil,
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			description:        "SetUserPermissions  error",
			mocking:            true,
			query:              "11",
			userPermissions:    []string{"addmovie", "deletemovie"},
			payload:            []byte(`{"permissions":["addmovie","deletemovie" ]}`),
			returnArgument:     fmt.Errorf("something went wrong"),
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			description:        "success  SetUserPermissions",
			mocking:            true,
			query:              "12",
			userPermissions:    []string{"update:user", "add:user"},
			payload:            []byte(`{"permissions":["update:user","add:user" ]}`),
			returnArgument:     nil,
			expectedStatusCode: http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			if tc.mocking {
				userId, _ := strconv.ParseInt(tc.query, 10, 64)
				mockService.On("SetUserPermissions", mock.Anything, userId, tc.userPermissions).Return(tc.returnArgument)
			}

			w := httptest.NewRecorder()

			req, _ := http.NewRequest("POST", fmt.Sprintf("/user/%s", tc.query), bytes.NewBuffer(tc.payload))

			router.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedStatusCode, w.Code)

		})
	}
	mockService.AssertExpectations(t)
}

func TestDeleteUserPermissionHandler(t *testing.T) {
	router, handlers, mockService := setupTest()
	router.DELETE("/user/:id", handlers.DeleteUserPermissionHandler)
	testCases := []struct {
		description        string
		mocking            bool
		query              string
		userPermissions    []string
		payload            []byte
		returnArgument     error
		expectedStatusCode int
	}{
		{
			description:        "Invalid id",
			mocking:            false,
			query:              "test",
			userPermissions:    nil,
			payload:            nil,
			returnArgument:     nil,
			expectedStatusCode: http.StatusUnprocessableEntity,
		},
		{
			description:        "Invalid payload",
			mocking:            false,
			query:              "10",
			userPermissions:    nil,
			payload:            nil,
			returnArgument:     nil,
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			description:        "DeleteUserPermission  error",
			mocking:            true,
			query:              "11",
			userPermissions:    []string{"updatemovie", "deletemovie"},
			payload:            []byte(`{"permissions":["updatemovie","deletemovie" ]}`),
			returnArgument:     fmt.Errorf("something went wrong"),
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			description:        "success  DeleteUserPermission",
			mocking:            true,
			query:              "4",
			userPermissions:    []string{"update:user", "add:user"},
			payload:            []byte(`{"permissions":["update:user","add:user" ]}`),
			returnArgument:     nil,
			expectedStatusCode: http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			if tc.mocking {
				userId, _ := strconv.ParseInt(tc.query, 10, 64)
				mockService.On("DeleteUserPermission", mock.Anything, userId, tc.userPermissions).Return(tc.returnArgument)
			}

			w := httptest.NewRecorder()

			req, _ := http.NewRequest("DELETE", fmt.Sprintf("/user/%s", tc.query), bytes.NewBuffer(tc.payload))

			router.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedStatusCode, w.Code)

		})
	}
	mockService.AssertExpectations(t)

}
