package rpc

import (
	"context"

	sg "github.com/clubo-app/protobuf/story"
	"github.com/segmentio/ksuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s storyServer) GetStory(c context.Context, req *sg.GetStoryRequest) (*sg.Story, error) {
	_, err := ksuid.Parse(req.StoryId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid Story id")
	}

	story, err := s.sService.Get(c, req.StoryId)
	if err != nil {
		return nil, status.Error(codes.NotFound, "Story not found")
	}

	return story.ToGRPCStory(), nil
}
