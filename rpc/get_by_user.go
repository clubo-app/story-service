package rpc

import (
	"context"
	"encoding/base64"

	"github.com/clubo-app/packages/utils"
	sg "github.com/clubo-app/protobuf/story"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s storyServer) GetByUser(c context.Context, req *sg.GetByUserRequest) (*sg.PagedStories, error) {
	p, err := base64.URLEncoding.DecodeString(req.NextPage)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid Next Page Param")
	}
	stories, p, err := s.sService.GetByUser(c, req.UserId, p, req.Limit)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	nextPage := base64.URLEncoding.EncodeToString(p)

	var result []*sg.PublicStory
	for _, s := range stories {
		result = append(result, s.ToPublicStory())
	}

	return &sg.PagedStories{Stories: result, NextPage: nextPage}, nil
}
