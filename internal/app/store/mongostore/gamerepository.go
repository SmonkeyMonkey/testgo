package mongostore

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"

	"log"

	"test/internal/app/models"
	"time"
)

type GameRepository struct {
	store *Store
}

func (r *GameRepository) Create(g *models.Game, userId string) error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	_, err := r.store.db.Database("test").Collection("user_games").InsertOne(ctx, g)
	if err != nil {
		return err
	}
	err = r.store.redis.ZIncrBy("count_games", 1, userId).Err()
	if err != nil {
		return err
	}
	return nil
}
func (r *GameRepository) GetAll(page int) []models.Game {
	var games []models.Game
	var itemsPerPage = 20
	start := (page - 1) * itemsPerPage
	stop := start + itemsPerPage
	collection := r.store.db.Database("test").Collection("user_games")
	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
	cursor, _ := collection.Find(ctx, bson.M{})

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var game models.Game
		err := cursor.Decode(&game)
		if err != nil {
			log.Fatal("error on decoding the document", err)
		}
		games = append(games, game)

	}
	return games[start:stop]
}
func (r *GameRepository) GetTopUsers(page int) []models.User {
	var users []models.User
	var itemsPerPage = 15
	start := (page - 1) * itemsPerPage
	stop := start + itemsPerPage
	count, _ := r.store.db.Database("test").Collection("users").CountDocuments(context.TODO(), bson.M{})
	u, _ := r.store.redis.ZRevRangeWithScores("count_games", 0, count).Result()

	for _, id := range u {
		id, _ := primitive.ObjectIDFromHex(fmt.Sprintf("%v", id.Member))
		cursor, _ := r.store.db.Database("test").Collection("users").Find(context.TODO(), bson.M{"_id": id})
		for cursor.Next(context.TODO()) {
			var user models.User
			err := cursor.Decode(&user)
			if err != nil {
				log.Fatal("error on decoding the document", err)
			}
			users = append(users, user)
		}
	}

	return users[start:stop]
}
