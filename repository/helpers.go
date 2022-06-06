package repository

import (
	"fmt"
	"net/url"

	sg "github.com/clubo-app/protobuf/story"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/pgx"
	g "github.com/golang-migrate/migrate/v4/source/github"
	"github.com/segmentio/ksuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s Story) ToGRPCStory() *sg.Story {
	id, err := ksuid.Parse(s.ID)
	if err != nil {
		return nil
	}

	return &sg.Story{
		Id:            s.ID,
		PartyId:       s.PartyID,
		UserId:        s.UserID,
		Url:           s.Url,
		TaggedFriends: s.TaggedFriends,
		CreatedAt:     timestamppb.New(id.Time()),
	}
}

const version = 1

func validateSchema(url url.URL) error {
	url.Scheme = "pgx"
	url2 := fmt.Sprintf("%v%v", url.String(), "?sslmode=disable")
	g := g.Github{}
	d, err := g.Open("github://clubo-app/story-service/repository/migrations")
	if err != nil {
		return err
	}
	defer d.Close()

	m, err := migrate.NewWithSourceInstance("github", d, url2)

	if err != nil {
		return err
	}
	err = m.Migrate(version) // current version
	if err != nil && err != migrate.ErrNoChange {
		return err
	}
	defer m.Close()
	return nil
}
