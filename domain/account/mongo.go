package account

import (
	"GemiApp/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository struct {
	col *mongo.Collection
}

func NewMongoRepo(db *domain.MongoDB) Repository {
	return &MongoRepository{
		col: db.GetCollection("accounts"),
	}
}

func (r *MongoRepository) GetAccountByUsername(username string) (*Account, error) {
	filters := bson.D{{Key: "username", Value: username}}
	var acc Account
	if err := r.col.FindOne(context.TODO(), filters).Decode(&acc); err != nil {
		return nil, err
	}
	return &acc, nil
}

func (r *MongoRepository) GetAccountByID(user_id string) (*Account, error) {
	filters := bson.D{{Key: "_id", Value: user_id}}
	var acc Account
	if err := r.col.FindOne(context.TODO(), filters).Decode(&acc); err != nil {
		return nil, err
	}
	return &acc, nil
}
