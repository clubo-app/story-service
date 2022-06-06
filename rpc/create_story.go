package rpc

import (
	"context"
	"log"

	"github.com/clubo-app/packages/utils"
	sg "github.com/clubo-app/protobuf/story"
	"github.com/clubo-app/story-service/dto"
)

func (s storyServer) CreateStory(c context.Context, req *sg.CreateStoryRequest) (*sg.Story, error) {
	d := dto.Story{
		PartyId:       req.PartyId,
		UserId:        req.RequesterId,
		Url:           req.Url,
		TaggedFriends: req.TaggedFriends,
	}

	story, err := s.ss.Create(c, d)
	log.Println(story)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return story.ToGRPCStory(), err
}
