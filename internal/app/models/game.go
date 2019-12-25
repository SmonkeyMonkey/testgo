package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Game struct {
	Id           primitive.ObjectID `bson:"_id,omitempty"`
	PointsGained string             `json:"points_gained" bson:"points_gained"`
	WinStatus    string             `json:"win_status" bson:"win_status"`
	GameType     string             `json:"game_type" bson:"game_type"`
	Created      string             `json:"created" bson:"created"`
}
