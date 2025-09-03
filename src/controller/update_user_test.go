package controller

// import (
// 	"encoding/json"
// 	"io"
// 	"net/http"
// 	"net/http/httptest"
// 	"net/url"
// 	"strings"
// 	"testing"

// 	"github.com/LucasAntonioC-137/crud-go/src/configuration/logger"
// 	"github.com/LucasAntonioC-137/crud-go/src/configuration/rest_err"
// 	"github.com/LucasAntonioC-137/crud-go/src/controller/model/request"
// 	"github.com/LucasAntonioC-137/crud-go/src/model"
// 	"github.com/LucasAntonioC-137/crud-go/src/tests/mocks"
// 	"github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/assert"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.uber.org/mock/gomock"
// 	"go.uber.org/zap"
// )

// func TestUserControllerInterface_UpdateUser(t *testing.T) {

// 	control := gomock.NewController(t)
// 	defer control.Finish()

// 	service := mocks.NewMockUserDomainService(control)

// 	controller := NewUserControllerInterface(service)

// 	t.Run("validation_got_error",
// 		func(t *testing.T) {
// 			recorder := httptest.NewRecorder()
// 			context := GetTestGinContext(recorder)

// 			userRequest := request.UserUpdateRequest{
// 				Name: "tes",
// 				Age:  0,
// 			}

// 			b, err := json.Marshal(userRequest)
// 			if err != nil {
// 				logger.Error("Error trying to convert object to body", err,
// 					zap.String("journey", "testeUpdateUser"))
// 				return
// 			}
// 			stringReader := io.NopCloser(strings.NewReader(string(b)))

// 			MakeRequest(context, []gin.Param{}, url.Values{}, "PUT", stringReader)
// 			controller.UpdateUser(context)

// 			assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
// 		},
// 	)

// 	t.Run("ID_is_invalid_returns_error", func(t *testing.T) {
// 		recorder := httptest.NewRecorder()
// 		context := GetTestGinContext(recorder)

// 		userRequest := request.UserUpdateRequest{
// 			Name: "Test User",
// 			Age:  10,
// 		}

// 		b, err := json.Marshal(userRequest)
// 		if err != nil {
// 			logger.Error("Error trying to convert object to body", err,
// 				zap.String("journey", "testeCreateUser"))
// 			return
// 		}
// 		stringReader := io.NopCloser(strings.NewReader(string(b)))

// 		param := []gin.Param{
// 			{
// 				Key:   "userId",
// 				Value: "test",
// 			},
// 		}

// 		MakeRequest(context, param, url.Values{}, "PUT", stringReader)
// 		controller.UpdateUser(context)

// 		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
// 	})

// 	t.Run("object_is_valid_but_service_returns_error", func(t *testing.T) {
// 		recorder := httptest.NewRecorder()
// 		context := GetTestGinContext(recorder)
// 		ID := primitive.NewObjectID().Hex()

// 		userRequest := request.UserUpdateRequest{
// 			Name: "Test User",
// 			Age:  10,
// 		}

// 		domain := model.NewUserUpdateDomain(
// 			userRequest.Name,
// 			userRequest.Age,
// 		)

// 		b, err := json.Marshal(userRequest)
// 		if err != nil {
// 			logger.Error("Error trying to convert object to body", err,
// 				zap.String("journey", "testeCreateUser"))
// 			return
// 		}
// 		stringReader := io.NopCloser(strings.NewReader(string(b)))

// 		param := []gin.Param{
// 			{
// 				Key:   "userId",
// 				Value: ID,
// 			},
// 		}

// 		service.EXPECT().UpdateUserService(ID, domain).Return(rest_err.NewInternalServerError("error test"))

// 		MakeRequest(context, param, url.Values{}, "PUT", stringReader)
// 		controller.UpdateUser(context)

// 		assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
// 	})

// 	t.Run("object_is_valid_and_returns_success", func(t *testing.T) {
// 		recorder := httptest.NewRecorder()
// 		context := GetTestGinContext(recorder)
// 		ID := primitive.NewObjectID().Hex()

// 		userRequest := request.UserUpdateRequest{
// 			Name: "Test User",
// 			Age:  10,
// 		}

// 		domain := model.NewUserUpdateDomain(
// 			userRequest.Name,
// 			userRequest.Age,
// 		)

// 		b, err := json.Marshal(userRequest)
// 		if err != nil {
// 			logger.Error("Error trying to convert object to body", err,
// 				zap.String("journey", "testeCreateUser"))
// 			return
// 		}
// 		stringReader := io.NopCloser(strings.NewReader(string(b)))

// 		param := []gin.Param{
// 			{
// 				Key:   "userId",
// 				Value: ID,
// 			},
// 		}

// 		service.EXPECT().UpdateUserService(ID, domain).Return(nil)

// 		MakeRequest(context, param, url.Values{}, "PUT", stringReader)
// 		controller.UpdateUser(context)

// 		assert.EqualValues(t, http.StatusOK, recorder.Code)
// 	})
// }
