package database

import (
	"context"
	"time"

	"github.com/khihadysucahyo/orderaja/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	client *mongo.Client
}

func Connect() *DB {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	return &DB{client: client}
}

func (db *DB) Fetch() ([]*model.Item, error) {
	collection := db.client.Database("orderaja").Collection("items")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	var items []*model.Item
	for cursor.Next(ctx) {
		var item model.Item
		err := cursor.Decode(&item)
		if err != nil {
			return []*model.Item{}, err
		}

		items = append(items, &item)
	}

	return items, nil
}

func (db *DB) GetByID(id string) (*model.Item, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	collection := db.client.Database("orderaja").Collection("items")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var item model.Item
	err = collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&item)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (db *DB) Store(input *model.NewItem) (*model.Item, error) {
	collection := db.client.Database("orderaja").Collection("items")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := collection.InsertOne(ctx, input)
	if err != nil {
		return nil, err
	}

	return &model.Item{
		ID:       res.InsertedID.(primitive.ObjectID).Hex(),
		Name:     input.Name,
		Quantity: input.Quantity,
		Price:    input.Price,
	}, nil
}
