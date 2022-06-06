package rpc

import (
	"context"

	"github.com/clubo-app/packages/utils"
	sg "github.com/clubo-app/protobuf/story"
)

func (s storyServer) GetByUser(c context.Context, req *sg.GetByUserRequest) (*sg.PagedStories, error) {
	stories, err := s.ss.GetByUser(c, req.UserId, req.Offset, req.Limit)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	var result []*sg.Story
	for _, s := range stories {
		result = append(result, s.ToGRPCStory())
	}

	return &sg.PagedStories{Stories: result}, nil
}
