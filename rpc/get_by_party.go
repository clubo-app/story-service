package rpc

import (
	"context"

	"github.com/clubo-app/packages/utils"
	sg "github.com/clubo-app/protobuf/story"
	"github.com/segmentio/ksuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s storyServer) GetByParty(c context.Context, req *sg.GetByPartyRequest) (*sg.PagedStories, error) {
	_, err := ksuid.Parse(req.PartyId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid Party id")
	}

	stories, err := s.ss.GetByParty(c, req.PartyId, req.Offset, req.Limit)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	var result []*sg.Story
	for _, s := range stories {
		result = append(result, s.ToGRPCStory())
	}

	return &sg.PagedStories{Stories: result}, nil
}
