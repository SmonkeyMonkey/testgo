package mongostore

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"test/internal/app/models"
	"time"
)

type GameRepository struct {
	store *Store
}

func (r *GameRepository) Create(g *models.Game) error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	_, err := r.store.db.Database("test").Collection("user_games").InsertOne(ctx, g)
	if err != nil {
		return err
	}
	return nil
}
func (r *GameRepository) GetAll() []models.Game {
	var games []models.Game
	pagination := options.Find().SetLimit(25)
	collection := r.store.db.Database("test").Collection("user_games")
	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
	cursor, _ := collection.Find(ctx, bson.M{}, pagination)
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var user models.Game
		err := cursor.Decode(&user)
		if err != nil {
			log.Fatal("error on decoding the document", err)
		}
		games = append(games, user)
	}
	return games
}
