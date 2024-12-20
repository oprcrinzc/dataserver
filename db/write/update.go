package write

import (
	"context"
	"dataServer/db"

	"github.com/gofiber/fiber/v2/log"
	"go.mongodb.org/mongo-driver/bson"
)

func Update(what string, where db.Worker) bool {
	client := db.New()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := client.Database("plantData").Collection(what)

	res, err := coll.UpdateOne(context.TODO(), bson.D{{Key: "_id", Value: where.ID}},
		bson.D{{Key: "$set", Value: bson.D{
			{Key: "name", Value: where.Name},
			{Key: "mode", Value: where.Mode},
			{Key: "last_update", Value: where.LastUpdate},
			{Key: "temperature", Value: where.Temperature},
			{Key: "humidity", Value: where.Humidity},
			{Key: "water_level", Value: where.WaterLevel},
			{Key: "water_level_target", Value: where.WaterLevelTarget}}}})
	if err != nil {
		return false
	}
	log.Info(res)
	return true
}
