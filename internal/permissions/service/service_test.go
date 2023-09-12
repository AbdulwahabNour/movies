package service

import (
	"context"
	"fmt"
	"testing"
	"time"

	model "github.com/AbdulwahabNour/movies/internal/model/permission"
	"github.com/AbdulwahabNour/movies/pkg/httpError"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddPermission(t *testing.T) {

	service, mockRepo := setupTest()
	validator := validator.New()

	testCases := []struct {
		description    string
		permission     *model.Permission
		mocking        bool
		returnArgument error
		expectedErr    error
	}{
		{
			description:    "Invalid permission",
			permission:     &model.Permission{},
			mocking:        false,
			returnArgument: nil,
			expectedErr:    httpError.ParseValidationErrors(validator.Struct(&model.Permission{})),
		},
		{
			description:    "AddPermission Repo error",
			permission:     &model.Permission{Code: "test"},
			mocking:        true,
			returnArgument: fmt.Errorf("something went wrong"),
			expectedErr:    httpError.ParseErrors(fmt.Errorf("something went wrong")),
		},
		{
			description:    "Success AddPermission Repo ",
			permission:     &model.Permission{Code: "add:movie"},
			mocking:        true,
			returnArgument: nil,
			expectedErr:    nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			ctx, cancle := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancle()

			if tc.mocking {
				mockRepo.On("AddPermission", ctx, tc.permission).Return(tc.returnArgument)
			}
			err := service.AddPermission(ctx, tc.permission)

			assert.Equal(t, tc.expectedErr, err)
		})
	}
	mockRepo.AssertExpectations(t)
}

func TestGetPermission(t *testing.T) {
	service, mockRepo := setupTest()
	testCases := []struct {
		description     string
		permissionId    int64
		mocking         bool
		returnArguments []interface{}
		expectedErr     error
	}{
		{
			description:  "GetPermission Repo error",
			permissionId: -1,
			mocking:      true,
			returnArguments: []interface{}{
				&model.Permission{},
				fmt.Errorf("something went wrong"),
			},
			expectedErr: httpError.ParseErrors(fmt.Errorf("something went wrong")),
		},
		{
			description:  "Success GetPermission Repo ",
			permissionId: 10,
			mocking:      true,
			returnArguments: []interface{}{
				&model.Permission{},
				nil,
			},
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			ctx, cancle := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancle()

			if tc.mocking {
				mockRepo.On("GetPermission", ctx, tc.permissionId).Return(tc.returnArguments...)
			}
			_, err := service.GetPermission(ctx, tc.permissionId)

			assert.Equal(t, tc.expectedErr, err)
		})
	}
	mockRepo.AssertExpectations(t)

}

func TestUpdatePermission(t *testing.T) {

	service, mockRepo := setupTest()
	validator := validator.New()
	testCases := []struct {
		description    string
		permission     *model.Permission
		mocking        bool
		returnArgument error
		expectedErr    error
	}{
		{
			description:    "Invalid Permission",
			permission:     &model.Permission{},
			mocking:        false,
			returnArgument: nil,
			expectedErr:    httpError.ParseValidationErrors(validator.Struct(&model.Permission{})),
		},
		{
			description:    "UpdatePermission Repo error",
			permission:     &model.Permission{ID: -10, Code: "delete:movie"},
			mocking:        true,
			returnArgument: fmt.Errorf("something went wrong"),
			expectedErr:    httpError.ParseErrors(fmt.Errorf("something went wrong")),
		},
		{
			description:    "Success UpdatePermission",
			permission:     &model.Permission{ID: 11, Code: "delete:movie"},
			mocking:        true,
			returnArgument: nil,
			expectedErr:    nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			ctx, cancle := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancle()

			if tc.mocking {
				mockRepo.On("UpdatePermission", ctx, tc.permission).Return(tc.returnArgument)
			}
			err := service.UpdatePermission(ctx, tc.permission)

			assert.Equal(t, tc.expectedErr, err)
		})
	}
	mockRepo.AssertExpectations(t)
}

func TestDeletePermission(t *testing.T) {
	service, mockRepo := setupTest()

	testCases := []struct {
		description    string
		permissionId   int64
		mocking        bool
		returnArgument error
		expectedErr    error
	}{
		{
			description:    "DeletePermission Repo error",
			permissionId:   -1,
			mocking:        true,
			returnArgument: fmt.Errorf("something went wrong"),
			expectedErr:    httpError.ParseErrors(fmt.Errorf("something went wrong")),
		},
		{
			description:    "Success DeletePermission",
			permissionId:   10,
			mocking:        true,
			returnArgument: nil,
			expectedErr:    nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			ctx, cancle := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancle()

			if tc.mocking {
				mockRepo.On("DeletePermission", ctx, tc.permissionId).Return(tc.returnArgument)
			}
			err := service.DeletePermission(ctx, tc.permissionId)

			assert.Equal(t, tc.expectedErr, err)
		})
	}
	mockRepo.AssertExpectations(t)

}

func TestGetUserPermissions(t *testing.T) {
	service, mockRepo := setupTest()

	testCases := []struct {
		description     string
		userId          int64
		mocking         bool
		returnArguments []interface{}
		expectedErr     error
	}{
		{
			description: "UserPermissions Repo error",
			userId:      -1,
			mocking:     true,
			returnArguments: []interface{}{
				[]*model.Permission{},
				fmt.Errorf("something went wrong"),
			},
			expectedErr: httpError.ParseErrors(fmt.Errorf("something went wrong")),
		},
		{
			description: "Success GetUserPermissions",
			userId:      10,
			mocking:     true,
			returnArguments: []interface{}{
				[]*model.Permission{},
				nil,
			},
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			ctx, cancle := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancle()

			if tc.mocking {
				mockRepo.On("UserPermissions", ctx, tc.userId).Return(tc.returnArguments...)
			}
			_, err := service.GetUserPermissions(ctx, tc.userId)

			assert.Equal(t, tc.expectedErr, err)
		})
	}
	mockRepo.AssertExpectations(t)

}

func TestSetUserPermissions(t *testing.T) {
	service, mockRepo := setupTest()

	testCases := []struct {
		description    string
		userId         int64
		mocking        bool
		codes          []string
		returnArgument error
		expectedErr    error
	}{
		{
			description:    "Invalid permission",
			userId:         3,
			mocking:        false,
			codes:          []string{"movie121:121"},
			returnArgument: nil,
			expectedErr:    httpError.NewBadRequestError("invalid permission format"),
		},
		{
			description:    "Invalid userId",
			userId:         -100,
			mocking:        false,
			codes:          []string{"add:movie"},
			returnArgument: nil,
			expectedErr:    httpError.NewBadRequestError("user id less than 1"),
		},
		{
			description:    "AddUserPermissions Repo error",
			userId:         10,
			mocking:        true,
			codes:          []string{"add:movie", "delete:movie"},
			returnArgument: fmt.Errorf("something went wrong"),
			expectedErr:    httpError.ParseErrors(fmt.Errorf("something went wrong")),
		},
		{
			userId:         12,
			mocking:        true,
			codes:          []string{"add:movie", "delete:movie"},
			returnArgument: nil,
			expectedErr:    nil,
		},
	}
	for _, tc := range testCases {

		t.Run(tc.description, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()

			if tc.mocking {
				mockRepo.On("AddUserPermissions", mock.Anything, tc.userId, tc.codes).Return(tc.returnArgument)
			}

			err := service.SetUserPermissions(ctx, tc.userId, tc.codes...)

			assert.Equal(t, tc.expectedErr, err)
		})

	}
	mockRepo.AssertExpectations(t)
}

func TestDeleteUserPermission(t *testing.T) {
	service, mockRepo := setupTest()

	testCases := []struct {
		description    string
		userId         int64
		mocking        bool
		codes          []string
		returnArgument error
		expectedErr    error
	}{
		{
			description:    "Invalid permissions",
			userId:         1,
			mocking:        false,
			codes:          []string{"add20:111"},
			returnArgument: nil,
			expectedErr:    httpError.NewBadRequestError("invalid permission format"),
		},
		{
			description:    "Invalid userId",
			userId:         -11,
			mocking:        false,
			codes:          []string{"add:movie"},
			returnArgument: nil,
			expectedErr:    httpError.NewBadRequestError("user id less than 1"),
		},
		{
			description:    "DeleteUserPermission Repo error",
			userId:         1,
			mocking:        true,
			codes:          []string{"add:movie", "delete:movie"},
			returnArgument: fmt.Errorf("something went wrong"),
			expectedErr:    httpError.ParseErrors(fmt.Errorf("something went wrong")),
		},
		{
			description:    "Success DeleteUserPermission Repo ",
			userId:         2,
			mocking:        true,
			codes:          []string{"delete:movie", "update:movie"},
			returnArgument: nil,
			expectedErr:    nil,
		},
	}
	for _, tc := range testCases {

		t.Run(tc.description, func(t *testing.T) {

			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()

			if tc.mocking {
				mockRepo.On("DeleteUserPermission", mock.Anything, tc.userId, tc.codes).Return(tc.returnArgument)
			}

			err := service.DeleteUserPermission(ctx, tc.userId, tc.codes...)

			assert.Equal(t, tc.expectedErr, err)
		})

	}
	mockRepo.AssertExpectations(t)
}
