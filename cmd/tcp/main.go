package main

import (
	"github.com/joho/godotenv"
	"github.com/raulquiros/tcp_server/internal/platform/bus/inmemory"
	"github.com/raulquiros/tcp_server/internal/platform/server"
	"github.com/raulquiros/tcp_server/internal/platform/storage/redis"
	"github.com/raulquiros/tcp_server/internal/sku"
	"github.com/raulquiros/tcp_server/kit/command"
	"os"
)

func main() {
	loadEnvFile()

	//initialize the command bus
	var commandBus = inmemory.NewCommandBus()
	var commands = make(map[command.Type]command.Handler)

	//registering commands to command bus
	commands[sku.SkuCreateType] = sku.NewCreateSkuCommandHandler(redis.NewRedisSkuRepository())
	commandBus.Register(commands)

	s := server.New(
		&server.Config{Host: os.Getenv("TCP_HOST"), Port: os.Getenv("TCP_PORT")},
		commandBus,
	)

	s.Run()
}

func loadEnvFile() {
	_, err := os.Stat(".env")
	if os.IsNotExist(err) {
		return
	}

	err = godotenv.Load()

	if err != nil {
		panic("Error loading .env file")
	}
}
