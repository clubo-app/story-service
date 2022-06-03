package main

import (
	"log"

	"github.com/clubo-app/story-service/config"
	"github.com/clubo-app/story-service/repository"
	rpc "github.com/clubo-app/story-service/rpc"
	"github.com/clubo-app/story-service/service"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	pool, err := repository.NewPGXPool(c.DB_USER, c.DB_PW, c.DB_NAME, c.DB_HOST, c.DB_PORT)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	q := repository.New(pool)

	s := service.NewStoryServie(q)
	us := service.NewUploadService(c.SPACES_KEY, c.SPACES_ENDPOINT, c.SPACES_KEY)

	st := rpc.NewStoryServer(s, us)
	rpc.Start(st, c.PORT)
}
