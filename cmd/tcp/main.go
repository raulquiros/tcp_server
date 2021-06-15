package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/raulquiros/tcp_server/internal/platform/bus/inmemory"
	"github.com/raulquiros/tcp_server/internal/platform/server"
	"github.com/raulquiros/tcp_server/internal/platform/storage/redis"
	"github.com/raulquiros/tcp_server/internal/sku"
	"github.com/raulquiros/tcp_server/kit/command"
	"github.com/raulquiros/tcp_server/kit/query"
	"os"
)

func main() {
	loadEnvFile()

	storageServer := redis.New(
		&redis.Config{Host: os.Getenv("REDIS_HOST"), Port: os.Getenv("REDIS_PORT")},
	)

	//initialize the command bus
	var commandBus = inmemory.NewCommandBus()
	var commands = make(map[command.Type]command.Handler)

	//registering commands to command bus
	commands[sku.SkuCreateType] = sku.NewCreateSkuCommandHandler(redis.NewRedisSkuRepository(storageServer.Conn()))
	commandBus.Register(commands)

	//initialize the command bus
	var queryBus = inmemory.NewQueryBus()
	var queries = make(map[query.Type]query.Handler)

	//registering queries to query bus
	queries[sku.SkuListType] = sku.NewListSkuQueryHandler(redis.NewRedisSkuRepository(storageServer.Conn()))
	queryBus.Register(queries)

	s := server.New(
		&server.Config{Host: os.Getenv("TCP_HOST"), Port: os.Getenv("TCP_PORT")},
		commandBus,
	)

	s.Run()

	resp, _ := queryBus.Dispatch(context.Background(), sku.NewListSkuQuery())
	fmt.Println(resp.Data)
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
