package main

import (
	"dataServer/db"
	"dataServer/db/fetch"
	"dataServer/db/write"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// type updateBody struct {
// 	Name string
// 	Mode string
// 	Temperature float64
// 	Humidity float64
// 	WaterLevel float64
// }

func UpdateConfig(c *fiber.Ctx) error {
	name := c.Params("name")
	log.Info(name)
	if name == "" {
		return c.SendStatus(400)
	}
	// fa := fetch.Current()
	fc := fetch.Config()
	// ff := db.ConfigData{}
	body := db.ConfigDataNoID{}
	err := c.BodyParser(&body)
	if err != nil {
		log.Info(err)
		return c.SendStatus(400)
	}
	// for i, n := range fa {
	// 	if n.Name == name {
	// 		ff = fc[i]
	// 		break
	// 	}
	// }
	return c.JSON(fc)
	// return c.SendStatus(400)
}

func Update(c *fiber.Ctx) error {
	name := c.Params("name")
	// when := c.Params("*")
	if name == "" {
		return c.SendStatus(400)
	}
	fc := fetch.Current()
	ff := db.PlantData{}
	body := db.PlantDataNoID{}
	err := c.BodyParser(&body)
	if err != nil {
		log.Info(err)
		return c.SendStatus(400)
	}
	for _, n := range fc {
		if n.Name == name {
			ff = n
			break
		}
	}
	if body.Mode == "auto" || body.Mode == "manual" {
		ff.Mode = body.Mode
	}
	if body.Humidity != 0.0 {
		ff.Humidity = body.Humidity
	}
	if body.Temperature != 0.0 {
		ff.Temperature = body.Temperature
	}
	if body.WaterLevel != 0.0 {
		ff.WaterLevel = body.WaterLevel
	}
	ff.LastUpdate = time.Now().String()
	if write.Update("current", ff) {
		return c.JSON(ff)
	}
	return c.SendStatus(400)
}

func Create(c *fiber.Ctx) error {
	if c.Params("name") == "" {
		return c.SendStatus(400)
	}
	if write.Create[string, db.PlantDataNoID]("current", db.PlantDataNoID{
		LastUpdate:  time.Now().String(),
		Mode:        "auto",
		Name:        c.Params("name"),
		Temperature: 0.00,
		Humidity:    0.00,
		WaterLevel:  0.00,
		Birth:       time.Now().Unix(),
	}) == true {
		return c.SendStatus(201)
	}
	return c.SendStatus(400)
}

func Fetch(c *fiber.Ctx) error {
	where := c.Params("where")
	when := c.Params("*")
	names := []string{}
	fa := fetch.Current()
	for _, n := range fa {
		// log.Info(n.Name)
		names = append(names, n.Name)
	}
	// var fc
	if when == "config" {
		targetID := []primitive.ObjectID{}
		fc := fetch.Config()
		for _, n := range fc {
			// log.Info(n.TargetID)
			targetID = append(targetID, n.TargetID)
		}
		if where == "" || where == "current" {
			return c.JSON(fc)
		}
		if contain(names, where) {
			for i, n := range fa {
				if string(n.Name) == where {
					return c.JSON(fc[i])
				}
			}
		}
	}
	if when == "plant" {

		if where == "" || where == "current" {
			return c.JSON(fa)
		}
		if contain(names, where) {
			for i, n := range fa {
				if n.Name == where {
					return c.JSON(fa[i])
				}
			}
		}
	}

	return c.JSON("text no")
}

func contain[T comparable](src []T, v T) bool {
	for _, i := range src {
		if v == i {
			return true
		}
	}
	return false
}
