package redis

import (
	"fmt"
	"github.com/go-redis/redis/v8"
)

type Config struct {
	Host string
	Port string
}

type Server struct {
	conn *redis.Client
}

func New(config *Config) *Server {
	conn := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Host, config.Port),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return &Server{conn: conn}
}

func (server Server) Conn() *redis.Client {
	return server.conn
}
