package db

import "go.mongodb.org/mongo-driver/bson/primitive"

type Worker struct {
	ID               primitive.ObjectID `bson:"_id" json:"id"`
	LastUpdate       string             `bson:"last_update" json:"last_update"`
	Mode             string             `json:"mode"`
	Name             string             `json:"name"`
	Temperature      *float64           `json:"temperature"`
	Humidity         *float64           `json:"humidity"`
	WaterLevel       *float64           `bson:"water_level" json:"water_level" form:"water_level"`
	WaterLevelTarget *float64           `bson:"water_level_target" json:"water_level_target" form:"water_level_target"`
	Birth            int64              `json:"birth"`
}

type WorkerNoID struct {
	LastUpdate       string   `bson:"last_update" json:"last_update"`
	Mode             string   `json:"mode"`
	Name             string   `json:"name"`
	Temperature      *float64 `json:"temperature"`
	Humidity         *float64 `json:"humidity"`
	WaterLevel       *float64 `bson:"water_level" json:"water_level"`
	WaterLevelTarget *float64 `bson:"water_level_target" json:"water_level_target"`
	Birth            int64    `json:"birth"`
}

type Shiranaihito struct {
	ID   primitive.ObjectID `bson:"_id" json:"id"`
	Name string             `bson:"name" json:"name"`
	Ip   string             `bson:"ip" json:"ip"`
}
type ShiranaihitoNoID struct {
	Name string `bson:"name" json:"name"`
	Ip   string `bson:"ip" json:"ip"`
}

type Temperature struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Name        string             `bson:"name" json:"name"`
	Temperature *float64           `json:"temperature"`
	When        int64              `bson:"when" json:"when"`
}
type TemperatureNoID struct {
	Name        string   `bson:"name" json:"name"`
	Temperature *float64 `json:"temperature"`
	When        int64    `bson:"when" json:"when"`
}

type Humidity struct {
	ID       primitive.ObjectID `bson:"_id" json:"id"`
	Name     string             `bson:"name" json:"name"`
	Humidity *float64           `json:"humidity"`
	When     int64              `bson:"when" json:"when"`
}
type HumidityNoID struct {
	Name     string   `bson:"name" json:"name"`
	Humidity *float64 `json:"humidity"`
	When     int64    `bson:"when" json:"when"`
}

type HumidityTemperature struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Name        string             `bson:"name" json:"name"`
	Humidity    *float64           `json:"humidity"`
	Temperature *float64           `json:"temperature"`
	When        int64              `bson:"when" json:"when"`
}
type HumidityTemperatureNoID struct {
	Name        string   `bson:"name" json:"name"`
	Humidity    *float64 `json:"humidity"`
	Temperature *float64 `json:"temperature"`
	When        int64    `bson:"when" json:"when"`
}

type WaterLevel struct {
	ID         primitive.ObjectID `bson:"_id" json:"id"`
	Name       string             `bson:"name" json:"name"`
	WaterLevel *float64           `bson:"water_level" json:"water_level"`
	When       int64              `bson:"when" json:"when"`
}
type WaterLevelNoID struct {
	Name       string   `bson:"name" json:"name"`
	WaterLevel *float64 `bson:"water_level" json:"water_level"`
	When       int64    `bson:"when" json:"when"`
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
