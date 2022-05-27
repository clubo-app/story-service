package rpc

import (
	"context"

	"github.com/clubo-app/packages/utils"
	sg "github.com/clubo-app/protobuf/story"
	"github.com/clubo-app/story-service/dto"
)

func (s storyServer) CreateStory(c context.Context, req *sg.CreateStoryRequest) (*sg.PublicStory, error) {
	d := dto.Story{
		PartyId:       req.PartyId,
		UserId:        req.RequesterId,
		Lat:           float64(req.GetLat()),
		Long:          float64(req.GetLong()),
		Url:           req.Url,
		TaggedFriends: req.TaggedFriends,
	}

	story, err := s.sService.Create(c, d)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return story.ToPublicStory(), err
}
