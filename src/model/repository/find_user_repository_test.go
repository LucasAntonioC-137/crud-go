package repository

import (
	"fmt"
	"os"
	"testing"

	"github.com/LucasAntonioC-137/crud-go/src/model/repository/entity"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestUserRepository_FindUserByEmail(t *testing.T) {
	database_name := "user_database_test"
	collection_name := "user_collection_test"

	os.Setenv("MONGODB_USER_DB", collection_name)
	defer os.Clearenv()

	mtestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	// defer mtestDb.Close()

	mtestDb.Run("when_sending_a_valid_email_returns_success",
		func(mt *mtest.T) {
			userEntity := entity.UserEntity{
				ID:       primitive.NewObjectID(),
				Email:    "test@test.com",
				Password: "test",
				Name:     "test",
				Age:      50,
			}

			mt.AddMockResponses(mtest.CreateCursorResponse(
				1, fmt.Sprintf("%s.%s", database_name, collection_name),
				mtest.FirstBatch, convertEntityToBson(userEntity)))

			databaseMock := mt.Client.Database(database_name)

			repo := NewUserRepository(databaseMock)
			userDomain, err := repo.FindUserByEmail(userEntity.Email)
			assert.Nil(t, err)
			assert.EqualValues(t, userDomain.GetID(), userEntity.ID.Hex())
			assert.EqualValues(t, userDomain.GetEmail(), userDomain.GetEmail())
			assert.EqualValues(t, userDomain.GetName(), userDomain.GetName())
			assert.EqualValues(t, userDomain.GetAge(), userDomain.GetAge())
			assert.EqualValues(t, userDomain.GetPassword(), userDomain.GetPassword())
		},
	)

	mtestDb.Run("returns_error_when_mongodb_returns_error",
		func(mt *mtest.T) {

			mt.AddMockResponses(bson.D{
				{Key: "ok", Value: 0},
			})

			databaseMock := mt.Client.Database(database_name)

			repo := NewUserRepository(databaseMock)
			userDomain, err := repo.FindUserByEmail("test")

			assert.NotNil(t, err)
			assert.Nil(t, userDomain)
		},
	)

	mtestDb.Run("returns_no_document_found",
		func(mt *mtest.T) {

			mt.AddMockResponses(mtest.CreateCursorResponse(
				0, fmt.Sprintf("%s.%s", database_name, collection_name),
				mtest.FirstBatch))

			databaseMock := mt.Client.Database(database_name)

			repo := NewUserRepository(databaseMock)
			userDomain, err := repo.FindUserByEmail("test@gmail.com")

			assert.NotNil(t, err)
			assert.Nil(t, userDomain)
		},
	)
}

func TestUserRepository_FindUserById(t *testing.T) {
	database_name := "user_database_test"
	collection_name := "user_collection_test"

	os.Setenv("MONGODB_USER_DB", collection_name)
	defer os.Clearenv()

	mtestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	// defer mtestDb.Close()

	mtestDb.Run("when_sending_a_valid_id_returns_success",
		func(mt *mtest.T) {
			userEntity := entity.UserEntity{
				ID:       primitive.NewObjectID(),
				Email:    "test@test.com",
				Password: "test",
				Name:     "test",
				Age:      50,
			}

			mt.AddMockResponses(mtest.CreateCursorResponse(
				1, fmt.Sprintf("%s.%s", database_name, collection_name),
				mtest.FirstBatch, convertEntityToBson(userEntity)))

			databaseMock := mt.Client.Database(database_name)

			repo := NewUserRepository(databaseMock)
			userDomain, err := repo.FindUserByID(userEntity.ID.Hex())
			assert.Nil(t, err)
			assert.EqualValues(t, userDomain.GetID(), userEntity.ID.Hex())
			assert.EqualValues(t, userDomain.GetEmail(), userDomain.GetEmail())
			assert.EqualValues(t, userDomain.GetName(), userDomain.GetName())
			assert.EqualValues(t, userDomain.GetAge(), userDomain.GetAge())
			assert.EqualValues(t, userDomain.GetPassword(), userDomain.GetPassword())
		},
	)

	mtestDb.Run("returns_error_when_mongodb_returns_error",
		func(mt *mtest.T) {

			mt.AddMockResponses(bson.D{
				{Key: "ok", Value: 0},
			})

			databaseMock := mt.Client.Database(database_name)

			repo := NewUserRepository(databaseMock)
			userDomain, err := repo.FindUserByID("test")

			assert.NotNil(t, err)
			assert.Nil(t, userDomain)
		},
	)

	mtestDb.Run("returns_no_document_found",
		func(mt *mtest.T) {

			mt.AddMockResponses(mtest.CreateCursorResponse(
				0, fmt.Sprintf("%s.%s", database_name, collection_name),
				mtest.FirstBatch))

			databaseMock := mt.Client.Database(database_name)

			repo := NewUserRepository(databaseMock)
			userDomain, err := repo.FindUserByID("test@gmail.com")

			assert.NotNil(t, err)
			assert.Nil(t, userDomain)
		},
	)
}

func TestUserRepository_FindUserByEmailAndPassword(t *testing.T) {
	database_name := "user_database_test"
	collection_name := "user_collection_test"

	os.Setenv("MONGODB_USER_DB", collection_name)
	defer os.Clearenv()

	mtestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	// defer mtestDb.Close()

	mtestDb.Run("when_sending_a_valid_email_and_password_returns_success",
		func(mt *mtest.T) {
			userEntity := entity.UserEntity{
				ID:       primitive.NewObjectID(),
				Email:    "test@test.com",
				Password: "test",
				Name:     "test",
				Age:      50,
			}

			mt.AddMockResponses(mtest.CreateCursorResponse(
				1, fmt.Sprintf("%s.%s", database_name, collection_name),
				mtest.FirstBatch, convertEntityToBson(userEntity)))

			databaseMock := mt.Client.Database(database_name)

			repo := NewUserRepository(databaseMock)
			userDomain, err := repo.FindUserByEmailAndPassword(userEntity.Email, userEntity.Password)
			assert.Nil(t, err)
			assert.EqualValues(t, userDomain.GetID(), userEntity.ID.Hex())
			assert.EqualValues(t, userDomain.GetEmail(), userDomain.GetEmail())
			assert.EqualValues(t, userDomain.GetName(), userDomain.GetName())
			assert.EqualValues(t, userDomain.GetAge(), userDomain.GetAge())
			assert.EqualValues(t, userDomain.GetPassword(), userDomain.GetPassword())
		},
	)

	mtestDb.Run("returns_error_when_mongodb_returns_error",
		func(mt *mtest.T) {

			mt.AddMockResponses(bson.D{
				{Key: "ok", Value: 0},
			})

			databaseMock := mt.Client.Database(database_name)

			repo := NewUserRepository(databaseMock)
			userDomain, err := repo.FindUserByEmailAndPassword("test", "testpass")

			assert.NotNil(t, err)
			assert.Nil(t, userDomain)
		},
	)

	mtestDb.Run("returns_no_document_found",
		func(mt *mtest.T) {

			mt.AddMockResponses(mtest.CreateCursorResponse(
				0, fmt.Sprintf("%s.%s", database_name, collection_name),
				mtest.FirstBatch))

			databaseMock := mt.Client.Database(database_name)

			repo := NewUserRepository(databaseMock)
			userDomain, err := repo.FindUserByEmailAndPassword("test", "testpass")


			assert.NotNil(t, err)
			assert.Nil(t, userDomain)
		},
	)
}

func convertEntityToBson(userEntity entity.UserEntity) bson.D {
	return bson.D{
		{Key: "_id", Value: userEntity.ID},
		{Key: "email", Value: userEntity.Email},
		{Key: "password", Value: userEntity.Password},
		{Key: "name", Value: userEntity.Name},
		{Key: "age", Value: userEntity.Age},
	}
}
