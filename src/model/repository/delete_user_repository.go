package repository

import (
	"context"
	"os"

	"github.com/LucasAntonioC-137/crud-go/src/configuration/logger"
	"github.com/LucasAntonioC-137/crud-go/src/configuration/rest_err"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (ur *userRepository) DeleteUser(userId string) *rest_err.RestErr{
	logger.Info("Init deleteUser repository",
		zap.String("journey", "DeleteUser"))

	collection_name := os.Getenv(MONGODB_USER_DB)

	collection := ur.databseConnection.Collection(collection_name)

	objectId, _ := primitive.ObjectIDFromHex(userId)

	filter := bson.D{{Key: "_id", Value: objectId}}

	_, err := collection.DeleteOne(context.Background(), filter)
		if err != nil {
			logger.Error("Error trying to delete user", err,
				zap.String("journey", "deleteUser"))
			return rest_err.NewInternalServerError(err.Error())
		}

	logger.Info(
		"deleteUser repository executed sucessfully",
		zap.String("userId", userId),
		zap.String("journey", "DeleteUser"))

	return nil

}