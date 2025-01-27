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

var zero float64 = 0

func Update(c *fiber.Ctx) error {
	what := c.Params("what")
	where := c.Params("where")
	// log.Info("" + what + " " + where)

	if what == "workers" && where != "" {
		names, workers := getWorkersName()
		// log.Info(workers)
		worker := db.Worker{}
		data := db.Worker{}
		err := c.BodyParser(&data)
		if err != nil {
			log.Info(err.Error())
			return c.SendStatus(400)
		}
		if r, _ := contain(names, where); !r {
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
		if data.WaterLevelToFill != nil {
			worker.WaterLevelToFill = data.WaterLevelToFill
		}
		if data.Mode == "manual" || data.Mode == "auto" {
			worker.Mode = data.Mode
		}
		if data.Name != "" {
			worker.Name = data.Name
		}
		worker.LastUpdate = time.Now().String()
		write.Update("workers", worker)
		write.New("logs", db.LogsNoID{
			Name:             worker.Name,
			Humidity:         worker.Humidity,
			Temperature:      worker.Temperature,
			WaterLevel:       worker.WaterLevel,
			WaterLevelTarget: worker.WaterLevelTarget,
			WaterLevelToFill: worker.WaterLevelToFill,
			When:             time.Now().Unix(),
		})
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
		if r, _ := contain(names, where); r {
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
		if r, _ := contain(names, where); r {
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
	names, workers := getWorkersName()
	ip := c.IP()
	if r, at := contain(names, who); r {
		log.Info(workers[*at].ID.Hex())
		return c.JSON(workers[*at].ID.Hex())
	}
	shiran := fetch.Shiranaihito()
	for _, n := range shiran {
		if n.Name == who {
			write.Update("shiranaihito", db.Shiranaihito{
				ID:   n.ID,
				Name: who,
				Ip:   ip,
			})
			return c.SendString("go to register")
		}
	}
	write.New("shiranaihito", db.ShiranaihitoNoID{
		Name: who,
		Ip:   ip,
	})
	return c.SendString("go to register")
}

func Register(c *fiber.Ctx) error {
	what := c.Params("what")
	where := c.Params("where")
	if what == "workers" && where != "" {
		shiranai := fetch.Shiranaihito()
		for _, n := range shiranai {
			if n.Name == where {
				if n.Ip != c.IP() {
					return c.SendStatus(400)
				}
				data := db.WorkerNoID{}
				err := c.BodyParser(&data)
				if err != nil {
					log.Info(err.Error())
					if err.Error() != "Unprocessable Entity" {
						return c.SendStatus(400)
					}
				}
				if !write.Del("shiranaihito", n.ID) {
					return c.SendString("please go to gatekeeper")
				}
				if write.New("workers", db.WorkerNoID{
					Name:             where,
					Mode:             ternary(data.Mode == "", "Auto", data.Mode),
					LastUpdate:       time.Now().String(),
					Temperature:      ternary(data.Temperature == nil, &zero, data.Temperature),
					Humidity:         ternary(data.Humidity == nil, &zero, data.Humidity),
					WaterLevel:       ternary(data.WaterLevel == nil, &zero, data.WaterLevel),
					WaterLevelTarget: ternary(data.WaterLevelTarget == nil, &zero, data.WaterLevelTarget),
					Birth:            time.Now().Unix(),
				}) {
					return c.SendStatus(201)
				}
			}
		}
		return c.SendString("shiran!")
	}
	return c.SendStatus(400)
}

func contain[T comparable](src []T, v T) (bool, *int) {
	for h, i := range src {
		if v == i {
			return true, &h
		}
	}
	return false, nil
}

func getWorkersName() ([]string, []db.Worker) {
	workers := fetch.Workers()
	names := []string{}
	for _, n := range workers {
		names = append(names, n.Name)
	}
	return names, workers
}

func ternary[T comparable](condition bool, ifTrue T, ifFalse T) T {
	if condition {
		return ifTrue
	}
	return ifFalse
}
