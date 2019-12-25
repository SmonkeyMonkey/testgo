package mongostore

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"test/internal/app/models"
	"time"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *models.User) error {
	if err := u.Validation(); err != nil {
		return err
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	_, err := r.store.db.Database("test").Collection("users").InsertOne(ctx, u)
	if err != nil {
		return err
	}
	return nil
}
func (r *UserRepository) GetAll(page int) []models.User {
	var users []models.User
	itemsPerPage := 30
	skip := (page - 1) * itemsPerPage

	collection := r.store.db.Database("test").Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)

	projection := options.Find().SetSkip(int64(skip)).SetLimit(int64(itemsPerPage))
	cursor, _ := collection.Find(ctx, bson.M{}, projection)
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var user models.User
		err := cursor.Decode(&user)
		if err != nil {
			log.Fatal("error on decoding the document", err)
		}
		users = append(users, user)
	}
	return users
}
