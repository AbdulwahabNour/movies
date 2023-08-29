package service

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	model "github.com/AbdulwahabNour/movies/internal/model/permission"
	"github.com/AbdulwahabNour/movies/pkg/httpError"
	"github.com/stretchr/testify/assert"
)

func TestAddPermission(t *testing.T) {
	service, mockRepo := setupTest()
	ctx, cancle := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancle()
	err := service.AddPermission(ctx, &model.Permission{})
	errOut := httpError.NewHttpError(http.StatusBadRequest, httpError.ErrInvalidSyntax.Error(), map[string]string{"Code": "The field Code Required"})
	assert.EqualError(t, err, errOut.Error(), "Expected error does not match actual error")
	mockRepo.On("AddPermission", ctx, &model.Permission{ID: 10, Code: "test"}).
		Return(nil)
	err = service.AddPermission(ctx, &model.Permission{ID: 10, Code: "test"})
	assert.NoError(t, err)

	mockRepo.On("AddPermission", ctx, &model.Permission{ID: 11, Code: "test"}).
		Return(fmt.Errorf("not found"))
	err = service.AddPermission(ctx, &model.Permission{ID: 11, Code: "test"})
	assert.Error(t, err)

}

func TestGetPermission(t *testing.T) {
	service, mockRepo := setupTest()
	ctx, cancle := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancle()
	perm := model.Permission{ID: 11, Code: "test"}

	mockRepo.On("GetPermission", ctx, int64(11)).
		Return(&perm, nil)
	permOut, err := service.GetPermission(ctx, 11)
	assert.NoError(t, err)
	assert.Equal(t, &perm, permOut)

	mockRepo.On("GetPermission", ctx, int64(1)).
		Return(&model.Permission{}, fmt.Errorf("not found"))

	permOutNil, err := service.GetPermission(ctx, 1)
	assert.Error(t, err)
	assert.Nil(t, permOutNil)

}

func TestUpdatePermission(t *testing.T) {

	service, mockRepo := setupTest()

	ctx, cancle := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancle()

	err := service.UpdatePermission(ctx, &model.Permission{})
	errOut := httpError.NewHttpError(http.StatusBadRequest, httpError.ErrInvalidSyntax.Error(), map[string]string{"Code": "The field Code Required"})
	assert.EqualError(t, err, errOut.Error(), "Expected error does not match actual error")
	mockRepo.On("UpdatePermission", ctx, &model.Permission{ID: 1, Code: "test"}).
		Return(nil)

	err = service.UpdatePermission(ctx, &model.Permission{ID: 1, Code: "test"})
	assert.NoError(t, err)

	mockRepo.On("UpdatePermission", ctx, &model.Permission{ID: 2, Code: "test"}).
		Return(fmt.Errorf("not found"))

	err = service.UpdatePermission(ctx, &model.Permission{ID: 2, Code: "test"})
	assert.Error(t, err)

}

func TestDeletePermission(t *testing.T) {
	service, mockRepo := setupTest()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	mockRepo.On("DeletePermission", ctx, int64(1)).
		Return(nil)

	err := service.DeletePermission(ctx, 1)
	assert.NoError(t, err)

	mockRepo.On("DeletePermission", ctx, int64(2)).
		Return(fmt.Errorf("not found"))

	err = service.DeletePermission(ctx, 2)
	assert.Error(t, err)
}

func TestGetUserPermissions(t *testing.T) {
	service, mockRepo := setupTest()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	userPermissions := []*model.Permission{
		{ID: 1, Code: "permission1"},
		{ID: 2, Code: "permission2"},
	}

	mockRepo.On("UserPermissions", ctx, int64(11)).
		Return(userPermissions, nil)

	perms, err := service.GetUserPermissions(ctx, 11)
	assert.NoError(t, err)
	assert.Equal(t, userPermissions, perms)
	mockRepo.On("UserPermissions", ctx, int64(2)).
		Return([]*model.Permission{}, fmt.Errorf("not found"))

	permsNil, err := service.GetUserPermissions(ctx, 2)
	assert.Error(t, err)
	assert.Nil(t, permsNil)

}

func TestSetUserPermission(t *testing.T) {
	service, mockRepo := setupTest()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	userPermissions := []model.UserPermission{
		{UserId: 0, PermissionId: 0},
		{UserId: 3, PermissionId: 3},
		{UserId: 5, PermissionId: 5},
	}
	errOut := httpError.NewHttpError(http.StatusBadRequest,
		httpError.ErrInvalidSyntax.Error(),
		map[string]string{"PermissionId": "The field PermissionId Required",
			"UserId": "The field UserId Required"})

	err := service.SetUserPermission(ctx, &userPermissions[0])
	assert.EqualError(t, errOut, err.Error())

	mockRepo.On("AddUserPermission", ctx, &userPermissions[1]).
		Return(nil)

	err = service.SetUserPermission(ctx, &userPermissions[1])
	assert.NoError(t, err)

	mockRepo.On("AddUserPermission", ctx, &userPermissions[2]).
		Return(fmt.Errorf("error adding user permission"))

	err = service.SetUserPermission(ctx, &userPermissions[2])
	assert.Error(t, err)
}

func TestDeleteUserPermission(t *testing.T) {
	service, mockRepo := setupTest()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	userPermissions := []model.UserPermission{
		{UserId: 0, PermissionId: 0},
		{UserId: 3, PermissionId: 3},
		{UserId: 5, PermissionId: 5},
	}
	errOut := httpError.NewHttpError(http.StatusBadRequest,
		httpError.ErrInvalidSyntax.Error(),
		map[string]string{"PermissionId": "The field PermissionId Required",
			"UserId": "The field UserId Required"})

	err := service.DeleteUserPermission(ctx, &userPermissions[0])
	assert.EqualError(t, errOut, err.Error())

	mockRepo.On("DeleteUserPermission", ctx, &userPermissions[1]).
		Return(nil)

	err = service.DeleteUserPermission(ctx, &userPermissions[1])
	assert.NoError(t, err)

	mockRepo.On("DeleteUserPermission", ctx, &userPermissions[2]).
		Return(fmt.Errorf("error deleting user permission"))

	err = service.DeleteUserPermission(ctx, &userPermissions[2])
	assert.Error(t, err)
}
