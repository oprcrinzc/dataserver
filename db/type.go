package db

import "go.mongodb.org/mongo-driver/bson/primitive"

type PlantData struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	LastUpdate  string             `bson:"last_update" json:"last_update"`
	Mode        string             `json:"mode"`
	Name        string             `json:"name"`
	Temperature float64            `json:temperature`
	Humidity    float64            `json:humidity`
	WaterLevel  float64            `json:water_level`
	Birth       int64              `json:birth`
}

type PlantDataNoID struct {
	LastUpdate  string  `bson:"last_update" json:"last_update"`
	Mode        string  `json:"mode"`
	Name        string  `json:"name"`
	Temperature float64 `json:temperature`
	Humidity    float64 `json:humidity`
	WaterLevel  float64 `json:water_level`
	// LightLevel  float64 `json:water_level`
	Birth int64 `json:birth`
}
