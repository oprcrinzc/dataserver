package main

import (
	"dataServer/db"
	"dataServer/db/fetch"
	"dataServer/db/write"
	"fmt"
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

func Update_(c *fiber.Ctx) error {
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
	if body.Humidity != nil {
		ff.Humidity = body.Humidity
	}
	if body.Temperature != nil {
		ff.Temperature = body.Temperature
	}
	if body.WaterLevel != nil {
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
	if zero := 0.00; write.New[string, db.WorkerNoID]("current", db.WorkerNoID{
		LastUpdate:  time.Now().String(),
		Mode:        "auto",
		Name:        c.Params("name"),
		Temperature: &zero,
		Humidity:    &zero,
		WaterLevel:  &zero,
		Birth:       time.Now().Unix(),
	}) == true {
		return c.SendStatus(201)
	}
	return c.SendStatus(400)
}

func Update(c *fiber.Ctx) error {
	what := c.Params("what")
	where := c.Params("where")

	if what == "workers" && where != "" {
		names, workers := getWorkersName()
		log.Info(workers)
		worker := db.Worker{}
		data := db.Worker{}
		err := c.BodyParser(&data)
		if err != nil {
			log.Info(err.Error())
			return c.SendStatus(400)
		}
		if !contain(names, where) {
			return c.SendString(fmt.Sprintf("the \"%s\" does not exits", where))
		}
		for i, n := range workers {
			if n.Name == where {
				worker = workers[i]
			}
		}
		if !(worker.ID == data.ID) {
			return c.SendString(fmt.Sprintf("the ID \"%v\" does not exits", data.ID))
		}
		if data.Humidity != nil {
			worker.Humidity = data.Humidity
		}
		if data.Temperature != nil {
			worker.Temperature = data.Temperature
		}
		if data.WaterLevel != nil {
			worker.WaterLevel = data.WaterLevel
		}
		if data.WaterLevelTarget != nil {
			worker.WaterLevelTarget = data.WaterLevelTarget
		}
		if data.Mode == "manual" || data.Mode == "auto" {
			worker.Mode = data.Mode
		}
		if data.Name != "" {
			worker.Name = data.Name
		}
		worker.LastUpdate = time.Now().String()
		write.Update("workers", worker)
		return c.JSON(worker)
	}

	return c.SendStatus(400)
}

func Fetch(c *fiber.Ctx) error {
	what := c.Params("what")
	where := c.Params("where")
	names, _ := getWorkersName()
	workers := fetch.Workers()
	configs := fetch.Configs()
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
	who := c.Params("who")
	// workers := fetch.Workers()
	names, _ := getWorkersName()
	if contain(names, who) {
		return c.SendString("ok")
	}
	write.New[string, db.Shiranaihito]("shiranaihito", db.Shiranaihito{
		Name: who,
		Ip:   c.IP(),
	})
	return c.JSON(db.Shiranaihito{
		Name: who,
		Ip:   c.IP(),
	})
	return c.SendString("who are u?, huh")
}

func Register(c *fiber.Ctx) error {
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

func getWorkersName() ([]string, []db.Worker) {
	workers := fetch.Workers()
	names := []string{}
	for _, n := range workers {
		names = append(names, n.Name)
	}
	return names, workers
}
