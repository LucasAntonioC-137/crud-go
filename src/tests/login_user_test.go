package tests

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/LucasAntonioC-137/crud-go/src/controller/model/request"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestLoginUser(t *testing.T) {


		t.Run("user_and_password_is_not_valid", func(t *testing.T) {
		recorderCreateUser := httptest.NewRecorder()
		ctxCreateUser := GetTestGinContext(recorderCreateUser)
		recorderLoginUser := httptest.NewRecorder()
		ctxLoginUser := GetTestGinContext(recorderLoginUser)

		email := fmt.Sprintf("%d@test.com", rand.Int())
		password := fmt.Sprintf("%d#@@#", rand.Int())

		userRequest := request.UserRequest{
			Email:    email,
			Password: password,
			Name:     "Test User",
			Age:      10,
		}

		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		MakeRequest(ctxCreateUser, []gin.Param{}, url.Values{}, "POST", stringReader)
		UserController.CreateUser(ctxCreateUser)

		userLoginRequest := request.UserLogin{
			Email:    "test@notpassing.com",
			Password: "dkfjshfkdsfh##@@",
		}

		b, _ = json.Marshal(userLoginRequest)
		stringReader = io.NopCloser(strings.NewReader(string(b)))

		MakeRequest(ctxLoginUser, []gin.Param{}, url.Values{}, "POST", stringReader)
		UserController.LoginUser(ctxLoginUser)

		assert.EqualValues(t, http.StatusUnauthorized, recorderLoginUser.Result().StatusCode)
	})

	t.Run("user_and_password_is_valid", func(t *testing.T) {
		recorderCreateUser := httptest.NewRecorder()
		ctxCreateUser := GetTestGinContext(recorderCreateUser)
		recorderLoginUser := httptest.NewRecorder()
		ctxLoginUser := GetTestGinContext(recorderLoginUser)

		email := fmt.Sprintf("%d@test.com", rand.Int())
		password := fmt.Sprintf("%d#@@#", rand.Int())

		userRequest := request.UserRequest{
			Email:    email,
			Password: password,
			Name:     "Test User",
			Age:      10,
		}

		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		MakeRequest(ctxCreateUser, []gin.Param{}, url.Values{}, "POST", stringReader)
		UserController.CreateUser(ctxCreateUser)

		userLoginRequest := request.UserLogin{
			Email:    email,
			Password: password,
		}

		b, _ = json.Marshal(userLoginRequest)
		stringReader = io.NopCloser(strings.NewReader(string(b)))

		MakeRequest(ctxLoginUser, []gin.Param{}, url.Values{}, "POST", stringReader)
		UserController.LoginUser(ctxLoginUser)

		assert.EqualValues(t, http.StatusOK, recorderLoginUser.Result().StatusCode)
		assert.NotEmpty(t, recorderLoginUser.Result().Header.Get("Authorization"))
	})

}
