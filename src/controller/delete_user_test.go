package controller

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	// "github.com/LucasAntonioC-137/crud-go/src/configuration/rest_err"
	"github.com/LucasAntonioC-137/crud-go/src/configuration/rest_err"
	"github.com/LucasAntonioC-137/crud-go/src/tests/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"

	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
)

func TestUserControllerInterface_DeleteUser(t *testing.T) {
	control := gomock.NewController(t)
	defer control.Finish()

	service := mocks.NewMockUserDomainService(control)

	controller := NewUserControllerInterface(service)

	t.Run("userId_is_invalid_returns_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)
		// ID := primitive.NewObjectID().Hex()

		param := []gin.Param{
			{
				Key:   "userId",
				Value: "test",
			},
		}

		// service.EXPECT().DeleteUserService("test").Return(nil, rest_err.NewBadRequestError("error test"))

		MakeRequest(context, param, url.Values{}, "DELETE", nil)
		controller.DeleteUser(context)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("id_is_valid_service_returns_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)
		ID := primitive.NewObjectID().Hex()

		param := []gin.Param{
			{
				Key:   "userId",
				Value: ID,
			},
		}

		service.EXPECT().DeleteUserService(ID).Return(rest_err.NewInternalServerError("error test"))

		MakeRequest(context, param, url.Values{}, "DELETE", nil)
		controller.DeleteUser(context)

		assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("id_is_valid_service_returns_success", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)
		ID := primitive.NewObjectID().Hex()

		param := []gin.Param{
			{
				Key:   "userId",
				Value: ID,
			},
		}

		service.EXPECT().DeleteUserService(ID).Return(nil)

		MakeRequest(context, param, url.Values{}, "DELETE", nil)
		controller.DeleteUser(context)

		assert.EqualValues(t, http.StatusOK, recorder.Code)
	})
}
