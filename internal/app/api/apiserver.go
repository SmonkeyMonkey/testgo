package api

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"test/internal/app/store/mongostore"
	"time"
)

func Start() error {
	db, err := newDb()
	if err != nil {
		return err
	}
	store := mongostore.New(db)
	server := NewServer(store)
	return http.ListenAndServe(":8080", server)
}
func newDb() (*mongo.Client, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	db, err := mongo.Connect(ctx, options.Client())
	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}
