package write

import (
	"context"
	"dataServer/db"

	"github.com/gofiber/fiber/v2/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Del(where string, what primitive.ObjectID) bool {
	client := db.New()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := client.Database("plantData").Collection(where)
	res, err := coll.DeleteOne(context.TODO(), bson.D{{Key: "_id", Value: what}})
	if err != nil {
		log.Infof(err.Error())
		return false
	}
	log.Info(res)
	return true
}
