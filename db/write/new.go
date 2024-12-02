package write

import (
	"context"
	"dataServer/db"

	"github.com/gofiber/fiber/v2/log"
)

func Create[W comparable, G comparable](where string, what G) bool {
	client := db.New()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := client.Database("plantData").Collection(where)
	res, err := coll.InsertOne(context.TODO(), what)
	if err != nil {
		log.Infof(err.Error())
		return false
	}
	log.Info(res)
	return true
}
