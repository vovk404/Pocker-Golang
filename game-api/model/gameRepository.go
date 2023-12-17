package model

import (
	"context"
	"log"
	"time"
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
const dbName = "pocker"
const collection = "games"

var client *mongo.Client

type Models struct {
	GameEntry GameEntry
}

type GameEntry struct {
	ID string `bson:"_id,omitempty" json:"id,omitempty"`
	UserId int `bson:"user_id,omitempty" json:"user_id,omitempty"`
	Data string `bson:"data" json:"data"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

func InitClient(mongo *mongo.Client) *mongo.Client {
	client = mongo
	return client
}

func Insert(game *Game, userId int) error {
	collection := client.Database(dbName).Collection(collection)
	gameJson, err := json.MarshalIndent(game, "", "\t")
	if err != nil {
		return err
	}

	_, error := collection.InsertOne(context.TODO(), GameEntry{
		UserId: userId,
		Data: string(gameJson),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	
	if error != nil {
		log.Println("Error inserting into games:", error)
		return error
	}

	return nil
}

func All() ([]*GameEntry, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := client.Database(dbName).Collection(collection)
	opts := options.Find()
	opts.SetSort(bson.D{{"created_at", -1}})

	cursor, err := collection.Find(context.TODO(), bson.D{}, opts)
	
	if err != nil {
		log.Println("Finding all docs error: ", err)
		return nil, err
	}

	defer cursor.Close(ctx)

	var games []*GameEntry

	for cursor.Next(ctx) {
		var item GameEntry
		err := cursor.Decode(&item)
		if err != nil {
			log.Println("Error decoding dames into slice:", err)
			return nil, err
		} else {
			games = append(games, &item)
		}
	}
	
	return games, nil
}

func GetOne(id string) (*GameEntry, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := client.Database(dbName).Collection(collection)
	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var entry GameEntry
	err = collection.FindOne(ctx, bson.M{"_id": docID}).Decode(&entry)
	if err != nil {
		return nil, err
	}
	return &entry, nil
}

func GetByUserId(id string) (*GameEntry, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := client.Database(dbName).Collection(collection)
	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var entry GameEntry
	err = collection.FindOne(ctx, bson.M{"user_id": docID}).Decode(&entry)
	if err != nil {
		return nil, err
	}
	return &entry, nil
}

func DropCollection() error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := client.Database(dbName).Collection(collection)

	if err := collection.Drop(ctx); err != nil {
		return err
	}

	return nil
}

func Update(entry GameEntry) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := client.Database(dbName).Collection(collection)
	docID, err := primitive.ObjectIDFromHex(entry.ID)
	if err != nil {
		return nil, err
	}
	result, err := collection.UpdateOne(
		ctx, 
		bson.M{"_id": docID}, 
		bson.D{
			{"$set", bson.D{
				{"data", entry.Data},
				{"updated_at", time.Now()},
			}},
		})
	if err != nil {
		return nil, err
	}
	return result, nil
}