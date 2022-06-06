package rpc

import (
	"context"

	"github.com/clubo-app/packages/utils"
	common "github.com/clubo-app/protobuf/common"
	sg "github.com/clubo-app/protobuf/story"
)

func (s storyServer) DeleteStory(c context.Context, req *sg.DeleteStoryRequest) (*common.MessageResponse, error) {
	err := s.ss.Delete(c, req.RequesterId, req.StoryId)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return &common.MessageResponse{Message: "Story removed"}, nil
}
