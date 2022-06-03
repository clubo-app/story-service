package rpc

import (
	"context"

	"github.com/clubo-app/packages/utils"
	sg "github.com/clubo-app/protobuf/story"
)

func (s storyServer) GetByParty(c context.Context, req *sg.GetByPartyRequest) (*sg.PagedStories, error) {
	stories, err := s.sService.GetByParty(c, req.PartyId, req.Offset, req.Limit)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	var result []*sg.Story
	for _, s := range stories {
		result = append(result, s.ToGRPCStory())
	}

	return &sg.PagedStories{Stories: result}, nil
}
