package http

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	model "github.com/AbdulwahabNour/movies/internal/model/permission"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddPermissionHandler(t *testing.T) {

	router, handlers, mockService := setupTest()
	router.POST("/permissions", handlers.AddPermissionHandler)

	mockService.On("AddPermission", mock.Anything, &model.Permission{
		Code: "test",
	}).
		Return(nil)

	w := httptest.NewRecorder()
	payload := []byte(`{"code":"test"}`)
	req, _ := http.NewRequest("POST", "/permissions", bytes.NewBuffer(payload))

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	mockService.AssertExpectations(t)
}

func TestGetPermissioHandler(t *testing.T) {

	router, handlers, mockService := setupTest()
	router.GET("/:id", handlers.GetPermissioHandler)

	t.Run("Successful Get Permission", func(t *testing.T) {

		mockService.On("GetPermission", mock.Anything, int64(11)).
			Return(&model.Permission{}, nil)

		w := httptest.NewRecorder()

		req, _ := http.NewRequest("GET", "/11", nil)

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		mockService.AssertExpectations(t)

	})

	t.Run("Not fount Get prtmission", func(t *testing.T) {
		mockService.On("GetPermission", mock.Anything, int64(12)).
			Return(nil, fmt.Errorf("id not found"))
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/12", nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		mockService.AssertExpectations(t)
	})

}

func TestUpdatePermission(t *testing.T) {

	router, handlers, mockService := setupTest()
	router.PUT("/", handlers.UpdatePermissionHandler)

	t.Run("Successful Update Permission", func(t *testing.T) {

		mockService.On("UpdatePermission", mock.Anything, &model.Permission{ID: 12, Code: "test"}).
			Return(nil)

		w := httptest.NewRecorder()
		payload := []byte(`{
			"id":12,
			"code":"test"
		}`)
		req, _ := http.NewRequest("PUT", "/", bytes.NewBuffer(payload))

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		mockService.AssertExpectations(t)

	})

	t.Run("Not fount Update prtmission", func(t *testing.T) {
		mockService.On("UpdatePermission", mock.Anything, &model.Permission{ID: 100, Code: "test"}).
			Return(fmt.Errorf("not found"))
		w := httptest.NewRecorder()
		payload := []byte(`{
			"id":100,
			"code":"test"
		}`)
		req, _ := http.NewRequest("PUT", "/", bytes.NewBuffer(payload))
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		mockService.AssertExpectations(t)
	})

}

func TestDeletePermissionHandler(t *testing.T) {
	router, handlers, mockService := setupTest()
	router.DELETE("/:id", handlers.DeletePermissionHandler)

	t.Run("Successful Delete Permission", func(t *testing.T) {

		mockService.On("DeletePermission", mock.Anything, int64(11)).
			Return(nil)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("DELETE", "/11", nil)

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		mockService.AssertExpectations(t)
	})

	t.Run("Not found Delete permission", func(t *testing.T) {
		mockService.On("DeletePermission", mock.Anything, int64(12)).
			Return(fmt.Errorf("not found"))

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("DELETE", "/12", nil)

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)

		mockService.AssertExpectations(t)
	})

}

func TestGetUserPermissionsHandler(t *testing.T) {
	router, handlers, mockService := setupTest()
	router.GET("/user/:id", handlers.GetUserPermissionsHandler)

	t.Run("Successful Get User Permissions", func(t *testing.T) {

		mockService.On("GetUserPermissions", mock.Anything, int64(10)).
			Return([]*model.Permission{}, nil)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/user/10", nil)

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		mockService.AssertExpectations(t)
	})

	t.Run("User Permissions not found", func(t *testing.T) {
		mockService.On("GetUserPermissions", mock.Anything, int64(11)).
			Return(nil, fmt.Errorf("user permissions not found"))

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/user/11", nil)

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)

		mockService.AssertExpectations(t)
	})

}

func TestSetUserPermissionHandler(t *testing.T) {
	router, handlers, mockService := setupTest()
	router.POST("/user/", handlers.SetUserPermissionHandler)

	t.Run("Successful Set User Permission", func(t *testing.T) {
		mockService.On("SetUserPermission", mock.Anything, &model.UserPermission{UserId: 10, PermissionId: 1}).
			Return(nil)
		w := httptest.NewRecorder()

		payload := []byte(`{"user_id":10,"permission_id":1}`)

		req, _ := http.NewRequest("POST", "/user/", bytes.NewBuffer(payload))

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

	})

}

func TestDeleteUserPermissionHandler(t *testing.T) {
	router, handlers, mockService := setupTest()
	router.DELETE("/user/", handlers.DeleteUserPermissionHandler)

	t.Run("Successful Delete User Permission", func(t *testing.T) {

		mockService.On("DeleteUserPermission", mock.Anything, &model.UserPermission{UserId: 10, PermissionId: 5}).
			Return(nil)
		payload := []byte(`{"user_id":10,"permission_id":5}`)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("DELETE", "/user/", bytes.NewBuffer(payload))

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		mockService.AssertExpectations(t)
	})

}
