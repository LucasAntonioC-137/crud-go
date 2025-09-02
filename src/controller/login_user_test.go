package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/LucasAntonioC-137/crud-go/src/configuration/logger"
	"github.com/LucasAntonioC-137/crud-go/src/configuration/rest_err"
	"github.com/LucasAntonioC-137/crud-go/src/controller/model/request"
	"github.com/LucasAntonioC-137/crud-go/src/model"
	"github.com/LucasAntonioC-137/crud-go/src/tests/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
)

func TestUserControllerInterface_LoginUser(t *testing.T) {

	control := gomock.NewController(t)
	defer control.Finish()

	service := mocks.NewMockUserDomainService(control)

	controller := NewUserControllerInterface(service)

	t.Run("validation_got_error",
		func(t *testing.T) {
			recorder := httptest.NewRecorder()
			context := GetTestGinContext(recorder)

			userRequest := request.UserLogin{
				Email:    "ERROR@_email",
				Password: "teste",
			}

			b, err := json.Marshal(userRequest)
			if err != nil {
				logger.Error("Error trying to convert object to body", err,
					zap.String("journey", "testeCreateUser"))
				return
			}
			stringReader := io.NopCloser(strings.NewReader(string(b)))

			MakeRequest(context, []gin.Param{}, url.Values{}, "POST", stringReader)
			controller.LoginUser(context)

			assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
		},
	)

	t.Run("object_is_valid_but_service_returns_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		userRequest := request.UserLogin{
			Email:    "test@test.com",
			Password: "teste@123",
		}

		domain := model.NewUserLoginDomain(
			userRequest.Email,
			userRequest.Password,
		)

		b, err := json.Marshal(userRequest)
		if err != nil {
			logger.Error("Error trying to convert object to body", err,
				zap.String("journey", "testeLoginUser"))
			return
		}
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		service.EXPECT().LoginUserService(domain).Return(nil, "", rest_err.NewInternalServerError("error test"))

		MakeRequest(context, []gin.Param{}, url.Values{}, "POST", stringReader)
		controller.LoginUser(context)

		assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("object_is_valid_and_service_returns_success", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)
		id := primitive.NewObjectID().Hex()

		userRequest := request.UserLogin{
			Email:    "test@test.com",
			Password: "teste@123",
		}

		domain := model.NewUserLoginDomain(
			userRequest.Email,
			userRequest.Password,
		)

		b, err := json.Marshal(userRequest)
		if err != nil {
			logger.Error("Error trying to convert object to body", err,
				zap.String("journey", "testeLoginUser"))
			return
		}
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		service.EXPECT().LoginUserService(domain).Return(domain, id, nil)

		MakeRequest(context, []gin.Param{}, url.Values{}, "POST", stringReader)
		controller.LoginUser(context) 

		assert.EqualValues(t, http.StatusOK, recorder.Code)
		assert.EqualValues(t, recorder.Header().Values("Authorization")[0], id)
	})
}
