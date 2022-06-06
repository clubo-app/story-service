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

	r, err := repository.NewStoryRepository(c.DB_USER, c.DB_PW, c.DB_NAME, c.DB_HOST, c.DB_PORT)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	ss := service.NewStoryServie(r)
	us := service.NewUploadService(c.SPACES_KEY, c.SPACES_ENDPOINT, c.SPACES_KEY)

	st := rpc.NewStoryServer(ss, us)
	rpc.Start(st, c.PORT)
}
