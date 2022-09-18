package models

import "time"

type Chat struct {
	ChatID   string    `json:"chatID" bson:"chatID"`
	UserID   string    `json:"userID" bson:"userID"`
	Text     string    `json:"text" bson:"text"`
	SendedAt time.Time `json:"sendedAt" bson:"sendedAt"`
}
