package write

import (
	"context"
	"dataServer/db"

	"github.com/gofiber/fiber/v2/log"
	"go.mongodb.org/mongo-driver/bson"
)

func Update(where string, what db.Worker) bool {
	client := db.New()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := client.Database("plantData").Collection(where)

	res, err := coll.UpdateOne(context.TODO(), bson.D{{Key: "_id", Value: what.ID}},
		bson.D{{Key: "$set", Value: bson.D{
			{Key: "mode", Value: what.Mode},
			{Key: "last_update", Value: what.LastUpdate},
			{Key: "temperature", Value: what.Temperature},
			{Key: "humidity", Value: what.Humidity},
			{Key: "waterlevel", Value: what.WaterLevel}}}})
	if err != nil {
		return false
	}
	log.Info(res)
	return true
}
