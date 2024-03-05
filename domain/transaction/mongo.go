package transaction

import (
	"GemiApp/domain"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository struct {
	col *mongo.Collection
}

func NewMongoRepo(db *domain.MongoDB) Repository {
	return &MongoRepository{
		col: db.GetCollection("transactions"),
	}
}
func (r *MongoRepository) GetAccountTransactions(user_id string) []Transaction {
	filters := bson.D{
		{Key: "$or", Value: []interface{}{
			bson.D{{Key: "to", Value: user_id}},
			bson.D{{Key: "from", Value: user_id}},
		}},
	}

	var arr = []Transaction{}
	cursor, err := r.col.Find(context.TODO(), filters)
	if err != nil {
    log.Println(err)
		return arr
	}

	if err = cursor.All(context.TODO(), &arr); err != nil {
    log.Println(err)
		return arr
	}
	return arr
}
