package mongostore

import (
	"context"
	"github.com/go-redis/redis/v7"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
	"time"
)

func TestDB(t *testing.T) *mongo.Client {
	t.Helper()
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	db, err := mongo.Connect(ctx, options.Client())
	if err != nil {
		t.Fatal(err)
	}
	if err := db.Ping(ctx, nil); err != nil {
		t.Fatal(err)
	}
	return db
}
func TestRedis(t *testing.T) *redis.Client {
	t.Helper()
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	err := client.Ping().Err()
	if err != nil {
		t.Fatal(err)
	}
	return client
}
