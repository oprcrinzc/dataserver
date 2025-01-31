package main

import (
	"os"

	"github.com/BurntSushi/toml"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Config struct {
	Ip   string `toml:"ip"`
	Port string `toml:"port"`
}

func main() {
	app := fiber.New()
	app.Use(cors.New())

	app.Get("/fetch/:what?/:where?", Fetch)
	app.Get("/gatekeeper/:who", Gatekeeper)
	app.Put("/update/:what/:where", Update)
	app.Post("/register/:what/:where", Register)

	// app.Listen("192.168.1.43:8888")
	var config Config
	if _, err := toml.DecodeFile("./config.toml", &config); err != nil {
		log.Info("config file error: " + err.Error())
	}
	if config.Ip == "" {
		log.Info("wlan ip")
		// os.Exit(1)
	}
	if config.Port == "" {
		log.Info("port error")
		os.Exit(2)
	}
	log.Info(config.Ip + ":" + config.Port)
	app.Listen(config.Ip + ":" + config.Port)
	// app.Listen(":" + config.Port)
	// app.Listen(":8888")
}
