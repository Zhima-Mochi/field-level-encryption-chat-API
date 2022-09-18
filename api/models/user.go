package models

type User struct {
	UserID   string `json:"userID" bson:"userID"`
	UserName string `json:"userName" bson:"userName"`
}
