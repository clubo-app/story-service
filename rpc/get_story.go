package rpc

import (
	"context"

	sg "github.com/clubo-app/protobuf/story"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s storyServer) GetStory(c context.Context, req *sg.GetStoryRequest) (*sg.PublicStory, error) {
	if req.StoryId == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid Story id")
	}

	story, err := s.sService.Get(c, req.StoryId)
	if err != nil {
		return nil, status.Error(codes.NotFound, "Story not found")
	}

	return story.ToPublicStory(), nil
}
