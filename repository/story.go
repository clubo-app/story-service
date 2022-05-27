package repository

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/clubo-app/story-service/datastruct"
	"github.com/go-playground/validator/v10"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scylladb/gocqlx/v2/table"
)

const (
	TABLE_NAME     string = "stories"
	STORY_BY_PARTY string = "stories_by_party"
	STORY_BY_USER  string = "stories_by_user"
)

var storyMetadata = table.Metadata{
	Name:    TABLE_NAME,
	Columns: []string{"id", "party_id", "user_id", "geohash", "url", "tagged_friends"},
	PartKey: []string{"id", "party_id"},
}
var storyTable = table.New(storyMetadata)

type StoryRepository interface {
	Create(c context.Context, s datastruct.Story) (datastruct.Story, error)
	Delete(c context.Context, uId, sId string) error
	Get(c context.Context, sId string) (datastruct.Story, error)
	GetByUser(c context.Context, uId string, page []byte, limit uint32) ([]datastruct.Story, []byte, error)
	GetByParty(c context.Context, pId string, page []byte, limit uint32) ([]datastruct.Story, []byte, error)
}

type storyRepository struct {
	sess *gocqlx.Session
}

func (r *storyRepository) Create(c context.Context, s datastruct.Story) (datastruct.Story, error) {
	v := validator.New()
	err := v.Struct(s)
	if err != nil {
		return datastruct.Story{}, err
	}

	stmt, names := qb.
		Insert(TABLE_NAME).
		Columns(storyMetadata.Columns...).
		TTL(time.Hour * 24).
		ToCql()

	// TODO: insert story id with party id intt story_by_party table

	err = r.sess.
		Query(stmt, names).
		BindStruct(s).
		ExecRelease()
	if err != nil {
		return datastruct.Story{}, err
	}

	return s, err
}

func (r *storyRepository) Delete(c context.Context, uId, sId string) error {
	stmt, names := qb.
		Delete(TABLE_NAME).
		Where(qb.Eq("id")).
		If(qb.Eq("user_id")).
		ToCql()

	err := r.sess.
		Query(stmt, names).
		BindMap((qb.M{"id": sId, "user_id": uId})).
		ExecRelease()
	if err != nil {
		return err
	}
	return nil
}

func (r *storyRepository) Get(c context.Context, sId string) (datastruct.Story, error) {
	var result datastruct.Story
	err := r.sess.
		Query(storyTable.Get()).
		BindMap((qb.M{"id": sId})).
		GetRelease(&result)
	if err != nil {
		if err.Error() == "not found" {
			return datastruct.Story{}, errors.New("story not found")
		}
		return datastruct.Story{}, err
	}

	return result, nil
}

func (r *storyRepository) GetByUser(c context.Context, uId string, page []byte, limit uint32) (result []datastruct.Story, nextPage []byte, err error) {
	stmt, names := qb.
		Select(STORY_BY_USER).
		Where(qb.Eq("user_id")).
		ToCql()

	q := r.sess.
		Query(stmt, names).
		BindMap((qb.M{"user_id": uId}))
	defer q.Release()

	q.PageState(page)
	if limit == 0 {
		q.PageSize(10)
	} else {
		q.PageSize(int(limit))
	}

	iter := q.Iter()
	err = iter.Select(&result)
	if err != nil {
		log.Println(err)
		return []datastruct.Story{}, nil, errors.New("no stories found")
	}

	return result, iter.PageState(), nil
}

func (r *storyRepository) GetByParty(c context.Context, pId string, page []byte, limit uint32) (result []datastruct.Story, nextPage []byte, err error) {
	stmt, names := qb.
		Select(STORY_BY_PARTY).
		Where(qb.Eq("party_id")).
		ToCql()

	q := r.sess.
		Query(stmt, names).
		BindMap((qb.M{"party_id": pId}))
	defer q.Release()

	q.PageState(page)
	if limit == 0 {
		q.PageSize(10)
	} else {
		q.PageSize(int(limit))
	}

	iter := q.Iter()
	err = iter.Select(&result)
	if err != nil {
		log.Println(err)
		return []datastruct.Story{}, nil, errors.New("no stories found")
	}

	return result, iter.PageState(), err
}
