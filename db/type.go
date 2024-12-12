package db

import "go.mongodb.org/mongo-driver/bson/primitive"

type PlantData struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	LastUpdate  string             `bson:"last_update" json:"last_update"`
	Mode        string             `json:"mode"`
	Name        string             `json:"name"`
	Temperature float64            `json:temperature`
	Humidity    float64            `json:humidity`
	WaterLevel  float64            `bson:"water_level" json:"water_level"`
	LightLevel  float64            `bson:"light_level" json:"light_level"`
	Ph          float64            `bson:"ph" json:"ph"`
	Birth       int64              `json:"birth"`
}

type PlantDataNoID struct {
	LastUpdate  string  `bson:"last_update" json:"last_update"`
	Mode        string  `json:"mode"`
	Name        string  `json:"name"`
	Temperature float64 `json:temperature`
	Humidity    float64 `json:humidity`
	WaterLevel  float64 `bson:"water_level" json:"water_level"`
	LightLevel  float64 `bson:"light_level" json:"light_level"`
	Ph          float64 `bson:"ph" json:"ph"`
	Birth       int64   `json:"birth"`
}

type ConfigData struct {
	ID                primitive.ObjectID `bson:"_id" json:"id"`
	TargetID          primitive.ObjectID `bson:"target_id" json:"target_id"`
	TargetTemperature float64            `bson:"target_temperature" json:"target_temperature"`
	TargetHumidity    float64            `bson:"target_humidity" json:"target_humidity"`
	TargetLight       float64            `bson:"target_light" json:"target_light"`
}

type ConfigDataNoID struct {
	TargetID          primitive.ObjectID `bson:"target_id" json:"target_id"`
	TargetTemperature float64            `bson:"target_temperature" json:"target_temperature"`
	TargetHumidity    float64            `bson:"target_humidity" json:"target_humidity"`
	TargetLight       float64            `bson:"target_light" json:"target_light"`
}
