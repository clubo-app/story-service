package service

import (
	"context"
	"errors"

	"github.com/clubo-app/story-service/datastruct"
	"github.com/clubo-app/story-service/dto"
	"github.com/clubo-app/story-service/repository"
	"github.com/gofrs/uuid"
	"github.com/mmcloughlin/geohash"
)

const GEOHASH_PRECISION uint = 9

type StoryService interface {
	Create(c context.Context, s dto.Story) (datastruct.Story, error)
	Delete(c context.Context, uId, sId string) error
	Get(c context.Context, sId string) (datastruct.Story, error)
	GetByUser(c context.Context, uId string, page []byte, limit uint32) ([]datastruct.Story, []byte, error)
	GetByParty(c context.Context, pId string, page []byte, limit uint32) ([]datastruct.Story, []byte, error)
}

type storyService struct {
	repo repository.StoryRepository
}

func NewStoryServie(repo repository.StoryRepository) StoryService {
	return &storyService{repo: repo}
}

func (sService storyService) Create(c context.Context, s dto.Story) (datastruct.Story, error) {
	uuid, err := uuid.NewV1()
	if err != nil {
		return datastruct.Story{}, errors.New("failed to gen Story id")
	}

	gHash := geohash.EncodeWithPrecision(s.Lat, s.Long, GEOHASH_PRECISION)

	ds := datastruct.Story{
		Id:            uuid.String(),
		PartyId:       s.PartyId,
		UserId:        s.UserId,
		GHash:         gHash,
		Url:           s.Url,
		TaggedFriends: s.TaggedFriends,
	}
	return sService.repo.Create(c, ds)
}

func (sService storyService) Get(c context.Context, sId string) (datastruct.Story, error) {
	return sService.repo.Get(c, sId)
}

func (sService storyService) GetByUser(c context.Context, uId string, page []byte, limit uint32) ([]datastruct.Story, []byte, error) {
	return sService.repo.GetByUser(c, uId, page, limit)
}

func (sService storyService) GetByParty(c context.Context, pId string, page []byte, limit uint32) ([]datastruct.Story, []byte, error) {
	return sService.repo.GetByParty(c, pId, page, limit)
}

func (sService storyService) Delete(c context.Context, uId, sId string) error {
	return sService.repo.Delete(c, uId, sId)
}
