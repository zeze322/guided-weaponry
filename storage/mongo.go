package storage

import (
	"context"
	"errors"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBStorage struct {
	coll *mongo.Collection
}

func NewMongoDBStorage(ctx context.Context, mongodbURL, mongodbDatabase, mongodbCollection string) (*MongoDBStorage, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongodbURL))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to mongo: %s", err)
	}

	if err = client.Ping(ctx, nil); err != nil {
		log.Println(err)
	}

	coll := client.Database(mongodbDatabase).Collection(mongodbCollection)

	return &MongoDBStorage{
		coll: coll,
	}, nil
}

func (s *MongoDBStorage) WeaponParams(ctx context.Context, category string) ([]*Params, error) {
	filter := bson.M{"category": category}
	cursor, err := s.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var result []*Params

	if err = cursor.All(ctx, &result); err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("category %s doesn't exist", category)
	}

	return result, nil
}

func (s *MongoDBStorage) InsertWeapon(ctx context.Context, params *Params) error {
	weapon := NewWeapon(params)
	filter := bson.M{"name": weapon.Name}

	if err := s.coll.FindOne(ctx, filter).Err(); errors.Is(err, mongo.ErrNoDocuments) {
		_, err = s.coll.InsertOne(ctx, weapon)
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("weapon %s already exists", weapon.Name)
	}

	return nil
}

func (s *MongoDBStorage) UpdateWeapon(ctx context.Context, params *Params) error {
	update := bson.M{"$set": UpdateWeaponParams(params)}
	filter := bson.M{"name": params.Name}

	if err := s.coll.FindOne(ctx, filter).Err(); !errors.Is(err, mongo.ErrNoDocuments) {
		_, err = s.coll.UpdateOne(ctx, filter, update)
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("weapon %s doesn't exist", params.Name)
	}

	return nil
}

func (s *MongoDBStorage) ListCategory(ctx context.Context) ([]Category, error) {

	return nil, nil
}
