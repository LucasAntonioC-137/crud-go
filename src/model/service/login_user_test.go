package service

import (
	"os"
	"testing"

	"github.com/LucasAntonioC-137/crud-go/src/configuration/rest_err"
	"github.com/LucasAntonioC-137/crud-go/src/model"
	"github.com/LucasAntonioC-137/crud-go/src/tests/mocks"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
)

func TestUserDomainService_LoginUserService(t *testing.T) {
	control := gomock.NewController(t)
	defer control.Finish()

	repository := mocks.NewMockUserRepository(control)
	service := &userDomainService{repository}

	t.Run("when_calling_repository_returns_error", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		userDomain := model.NewUserDomain("test@test.com", "test", "test", 50)
		userDomain.SetID(id)

		userDomainMock := model.NewUserDomain(userDomain.GetEmail(), userDomain.GetPassword(), userDomain.GetName(), userDomain.GetAge())
		userDomainMock.EncryptPassword()

		repository.EXPECT().FindUserByEmailAndPassword(userDomain.GetEmail(),
			userDomainMock.GetPassword()).Return(nil, rest_err.NewInternalServerError("error trying to find user by email and password"))

		user, token, err := service.LoginUserService(userDomain)
		assert.Nil(t, user)
		assert.Empty(t, token)
		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "error trying to find user by email and password")

	})

	t.Run("when_calling_create_token_returns_error", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		userDomain := model.NewUserDomain("test@test.com", "test", "test", 50)
		userDomain.SetID(id)

		userDomainMock := mocks.NewMockUserDomainInterface(control)

		userDomainMock.EXPECT().GetEmail().Return("test@test.com")
		userDomainMock.EXPECT().GetPassword().Return("test")
		userDomainMock.EXPECT().EncryptPassword()

		userDomainMock.EXPECT().GenerateToken().Return("", rest_err.NewInternalServerError("error trying to create token"))

		repository.EXPECT().FindUserByEmailAndPassword(
			"test@test.com", "test").Return(userDomainMock, nil)

		user, token, err := service.LoginUserService(userDomainMock)
		assert.Nil(t, user)
		assert.Empty(t, token)
		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "error trying to create token")
	})

	t.Run("when_user_and_password_is_valid_return_success", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		secret := "test"
		os.Setenv("JWT_SECRET_KEY", secret)
		defer os.Clearenv()

		userDomain := model.NewUserDomain("test@test.com", "test", "test", 50)
		userDomain.SetID(id)

		repository.EXPECT().FindUserByEmailAndPassword(
			userDomain.GetEmail(), gomock.Any()).Return(userDomain, nil)

		userDomainReturn, token, err := service.LoginUserService(userDomain)
		assert.Nil(t, err)
		assert.EqualValues(t, userDomainReturn.GetID(), id)
		assert.EqualValues(t, userDomainReturn.GetEmail(), userDomain.GetEmail())
		assert.EqualValues(t, userDomainReturn.GetPassword(), userDomain.GetPassword())
		assert.EqualValues(t, userDomainReturn.GetName(), userDomain.GetName())
		assert.EqualValues(t, userDomainReturn.GetAge(), userDomain.GetAge())

		tokenReturned, _ := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
				return []byte(secret), nil
			}

			return nil, rest_err.NewBadRequestError("invalid token")
		})

		_, ok := tokenReturned.Claims.(jwt.MapClaims)
		if !ok || !tokenReturned.Valid {
			t.FailNow()
			return
		}

	})
}
