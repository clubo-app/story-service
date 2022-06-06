package rpc

import (
	"context"

	"github.com/clubo-app/packages/utils"
	sg "github.com/clubo-app/protobuf/story"
	"github.com/segmentio/ksuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s storyServer) GetByUser(c context.Context, req *sg.GetByUserRequest) (*sg.PagedStories, error) {
	_, err := ksuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid User id")
	}

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
