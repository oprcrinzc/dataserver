package fetch

import (
	"context"
	"dataServer/db"

	"go.mongodb.org/mongo-driver/bson"
)

func Current() []db.PlantData {
	client := db.New()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := client.Database("plantData").Collection("current")
	// {Key: "status", Value: "on"}
	data, err := coll.Find(context.TODO(), bson.D{})
	var res []db.PlantData
	if err != nil {
		panic(err)
	} else {
		if err = data.All(context.TODO(), &res); err != nil {
			panic(err)
		}
		// fmt.Println(res)
	}
	return res
}

// func Name(name string) (db.PlantData, bool) {
// 	if name == "" {
// 		return db.PlantData{}, false
// 	}
// 	client := db.New()
// 	defer func() {
// 		if err := client.Disconnect(context.TODO()); err != nil {
// 			panic(err)
// 		}
// 	}()
// 	coll := client.Database("plantData").Collection("current")
// 	data := coll.FindOne(context.TODO(), bson.D{{Key: "name", Value: name}})
// 	if data == nil {
// 		return db.PlantData{}, false
// 	}
// 	var res db.PlantData
// 	if data.Decode(&res) != nil {

// 		log.Info(res)
// 	}
// 	return res, true
// }
