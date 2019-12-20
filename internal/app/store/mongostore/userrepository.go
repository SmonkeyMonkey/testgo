package mongostore

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
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
	var itemsPerPage = 20
	start := (page - 1) * itemsPerPage
	stop := start + itemsPerPage

	collection := r.store.db.Database("test").Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
	cursor, _ := collection.Find(ctx, bson.M{})
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var user models.User
		err := cursor.Decode(&user)
		if err != nil {
			log.Fatal("error on decoding the document", err)
		}
		users = append(users, user)
	}
	return users[start:stop]
}
