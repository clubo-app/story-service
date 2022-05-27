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

	sess, err := repository.NewDB(c.CQL_KEYSPACE, c.CQL_HOSTS)
	if err != nil {
		log.Fatal(err)
	}
	defer sess.Close()

	dao := repository.NewDAO(&sess)

	s := service.NewStoryServie(dao.NewStoryRepository())
	us := service.NewUploadService(c.SPACES_KEY, c.SPACES_ENDPOINT, c.SPACES_KEY)

	st := rpc.NewStoryServer(s, us)
	rpc.Start(st, c.PORT)
}
