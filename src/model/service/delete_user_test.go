package service

// import (
// 	"testing"

// 	"github.com/LucasAntonioC-137/crud-go/src/configuration/rest_err"
// 	"github.com/LucasAntonioC-137/crud-go/src/tests/mocks"
// 	"github.com/stretchr/testify/assert"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.uber.org/mock/gomock"
// )

// func TestUserDomainService_DeleteUser(t *testing.T) {
// 	control := gomock.NewController(t)
// 	defer control.Finish()

// 	repo := mocks.NewMockUserRepository(control)
// 	service := NewUserDomainService(repo)

// 	t.Run("when_sending_a_valid_user_and_userId_returns_success", func(t *testing.T) {
// 		id := primitive.NewObjectID().Hex()

// 		repo.EXPECT().DeleteUser(id).Return(nil)

// 		err := service.DeleteUserService(id)

// 		assert.Nil(t, err)
// 	})

// 	t.Run("when_sending_a_invalid_user_and_userId_returns_error", func(t *testing.T) {
// 		id := primitive.NewObjectID().Hex()

// 		repo.EXPECT().DeleteUser(id).Return(rest_err.NewInternalServerError("error trying to delete user"))

// 		err := service.DeleteUserService(id)

// 		assert.NotNil(t, err)
// 		assert.EqualValues(t, err.Message, "error trying to delete user")
// 	})
// }