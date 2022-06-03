package repository

import (
	sg "github.com/clubo-app/protobuf/story"
	"github.com/segmentio/ksuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s Story) ToGRPCStory() *sg.Story {
	id, err := ksuid.Parse(s.ID)
	if err != nil {
		return nil
	}

	return &sg.Story{
		Id:            s.ID,
		PartyId:       s.PartyID,
		UserId:        s.UserID,
		Url:           s.Url,
		TaggedFriends: s.TaggedFriends,
		CreatedAt:     timestamppb.New(id.Time()),
	}
}
