package write

import (
	"context"
	"dataServer/db"

	"github.com/gofiber/fiber/v2/log"
	"go.mongodb.org/mongo-driver/bson"
)

func Update(what string, where any) bool {
	client := db.New()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := client.Database("plantData").Collection(what)
	if worker, ok := where.(db.Worker); ok {
		res, err := coll.UpdateOne(context.TODO(), bson.D{{Key: "_id", Value: worker.ID}},
			bson.D{{Key: "$set", Value: bson.D{
				{Key: "name", Value: worker.Name},
				{Key: "mode", Value: worker.Mode},
				{Key: "last_update", Value: worker.LastUpdate},
				{Key: "temperature", Value: worker.Temperature},
				{Key: "humidity", Value: worker.Humidity},
				{Key: "water_level", Value: worker.WaterLevel},
				{Key: "water_level_target", Value: worker.WaterLevelTarget}}}})
		if err != nil {
			return false
		}
		log.Info(res)
		return true
	}
	if shiranai, ok := where.(db.Shiranaihito); ok {
		_, err := coll.UpdateOne(context.TODO(), bson.D{{Key: "_id", Value: shiranai.ID}},
			bson.D{{Key: "$set", Value: bson.D{
				{Key: "name", Value: shiranai.Name},
				{Key: "ip", Value: shiranai.Ip}}}})
		if err != nil {
			return false
		}
		log.Info("Updated " + shiranai.ID.Hex())
		return true
	}
	return false
}
