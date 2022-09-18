package handlers

import (
	"context"

	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
)

type ChatsHandler struct {
	collection  *mongo.Collection
	ctx         context.Context
	redisClient *redis.Client
}
