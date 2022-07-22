package datastruct

import "github.com/clubo-app/protobuf/story"

type Story struct {
	Id            string   `db:"story_id"       validate:"required"`
	PartyId       string   `db:"party_id"       validate:"required"`
	UserId        string   `db:"user_id"        validate:"required"`
	Url           string   `db:"url"            validate:"required"`
	TaggedFriends []string `db:"tagged_friends"`
}

func (s Story) ToGRPCStory() *story.Story {
	return &story.Story{
		Id:            s.Id,
		PartyId:       s.PartyId,
		UserId:        s.UserId,
		Url:           s.Url,
		TaggedFriends: s.TaggedFriends,
	}
}
