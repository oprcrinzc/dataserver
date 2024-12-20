package main

import (
	"dataServer/db"
	"dataServer/db/fetch"
	"dataServer/db/write"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
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
	fc := fetch.Configs()
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
	fc := fetch.Workers()
	ff := db.Worker{}
	body := db.WorkerNoID{}
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
	if write.Create[string, db.WorkerNoID]("current", db.WorkerNoID{
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

// type struct

func Fetch(c *fiber.Ctx) error {
	what := c.Params("what")
	where := c.Params("where")
	names := []string{}
	workers := fetch.Workers()
	configs := fetch.Configs()
	for _, n := range workers {
		// log.Info(n.Name)
		names = append(names, n.Name)
	}
	if what == "" {
		workersNconfigs := struct {
			Workers []db.Worker
			Configs []db.ConfigData
		}{
			Workers: workers,
			Configs: configs}
		return c.JSON(workersNconfigs)
	}
	if what == "config" && where != "" {
		if contain(names, where) {
			for i, n := range workers {
				if string(n.Name) == where {
					return c.JSON(configs[i])
				}
			}
		} else {
			return c.JSON("not found")
		}
	}

	if what == "worker" && where != "" {
		if contain(names, where) {
			for i, n := range workers {
				if n.Name == where {
					return c.JSON(workers[i])
				}
			}
		} else {
			return c.JSON("not found")
		}
	}

	if what == "configs" {
		return c.JSON(configs)
	}

	if what == "workers" {
		return c.JSON(workers)
	}

	return c.JSON("text no")
}

func Gatekeeper(c *fiber.Ctx) error {
	return c.JSON("")
}

func contain[T comparable](src []T, v T) bool {
	for _, i := range src {
		if v == i {
			return true
		}
	}
	return false
}
