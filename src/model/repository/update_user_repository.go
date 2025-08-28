package repository

import (
	"context"
	"os"

	"github.com/LucasAntonioC-137/crud-go/src/configuration/logger"
	"github.com/LucasAntonioC-137/crud-go/src/configuration/rest_err"
	"github.com/LucasAntonioC-137/crud-go/src/model"
	"github.com/LucasAntonioC-137/crud-go/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (ur *userRepository) UpdateUser(userId string, userDomain model.UserDomainInterface) *rest_err.RestErr {
	
	logger.Info("Init updateUser repository",
			zap.String("journey", "UpdateUser"))

	collection_name := os.Getenv(MONGODB_USER_DB)

	collection := ur.databseConnection.Collection(collection_name)

	value := converter.ConvertDomainToEntity(userDomain)

	objectId, _ := primitive.ObjectIDFromHex(userId)

	filter := bson.D{{Key: "_id", Value: objectId}}
	update := bson.D{{Key: "$set", Value: value}}

	_, err := collection.UpdateOne(context.Background(), filter, update)
		if err != nil {
			logger.Error("Error trying to update user", err,
				zap.String("journey", "updateUser"))
			return rest_err.NewInternalServerError(err.Error())
		}

	logger.Info(
		"updateUser repository executed sucessfully",
		zap.String("userId", userId),
		zap.String("journey", "UpdateUser"))

	return nil
}