package service

import (
	"context"

	"github.com/clubo-app/story-service/dto"
	"github.com/clubo-app/story-service/repository"
	"github.com/segmentio/ksuid"
)

const GEOHASH_PRECISION uint = 9

type StoryService interface {
	Create(c context.Context, s dto.Story) (repository.Story, error)
	Delete(c context.Context, uId, sId string) error
	Get(c context.Context, sId string) (repository.Story, error)
	GetByUser(ctx context.Context, uId string, offset int32, limit int32) ([]repository.Story, error)
	GetByParty(ctx context.Context, pId string, offset int32, limit int32) ([]repository.Story, error)
}

type storyService struct {
	q *repository.Queries
}

func NewStoryServie(q *repository.Queries) StoryService {
	return &storyService{q: q}
}

func (s *storyService) Create(ctx context.Context, ds dto.Story) (repository.Story, error) {
	res, err := s.q.CreateStory(ctx, repository.CreateStoryParams{
		ID:            ksuid.New().String(),
		UserID:        ds.UserId,
		PartyID:       ds.PartyId,
		Url:           ds.Url,
		TaggedFriends: ds.TaggedFriends,
	})
	if err != nil {
		return repository.Story{}, err
	}

	return res, nil
}

func (s *storyService) Get(ctx context.Context, id string) (repository.Story, error) {
	res, err := s.q.GetStory(ctx, id)
	if err != nil {
		return repository.Story{}, err
	}

	return res, nil
}

func (s *storyService) GetByUser(ctx context.Context, uId string, offset int32, limit int32) ([]repository.Story, error) {
	res, err := s.q.GetStoryByUser(ctx, repository.GetStoryByUserParams{
		UserID: uId,
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return []repository.Story{}, err
	}

	return res, nil
}

func (s *storyService) GetByParty(ctx context.Context, pId string, offset int32, limit int32) ([]repository.Story, error) {
	res, err := s.q.GetStoryByParty(ctx, repository.GetStoryByPartyParams{
		PartyID: pId,
		Limit:   limit,
		Offset:  offset,
	})
	if err != nil {
		return []repository.Story{}, err
	}

	return res, nil
}

func (s *storyService) Delete(ctx context.Context, uId, sId string) error {
	return s.q.DeleteStory(ctx, repository.DeleteStoryParams{
		ID:     sId,
		UserID: uId,
	})
}
