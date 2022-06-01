package datastruct

import (
	sg "github.com/clubo-app/protobuf/story"
	"github.com/gofrs/uuid"
	"github.com/mmcloughlin/geohash"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Story struct {
	Id            string   `json:"id"                       db:"id"             validate:"required"`
	PartyId       string   `json:"partyId"                  db:"party_id"       validate:"required"`
	UserId        string   `json:"userId"                   db:"user_id"        validate:"required"`
	GHash         string   `json:"geohash"                  db:"geohash"        validate:"required"`
	Url           string   `json:"url"                      db:"url"            validate:"required"`
	TaggedFriends []string `json:"tagged_friends,omitempty" db:"tagged_friends"`
}

type PagedStories struct {
	Stories  []sg.PublicStory `json:"stories,omitempty"`
	NextPage string           `json:"nextPage"`
}

func (s Story) ToPublicStory() *sg.PublicStory {
	lat, lon := geohash.DecodeCenter(s.GHash)

	uuidv1, err := uuid.FromString(s.Id)
	if err != nil {
		return &sg.PublicStory{}
	}
	timestamp, err := uuid.TimestampFromV1(uuidv1)
	if err != nil {
		return &sg.PublicStory{}
	}
	t, err := timestamp.Time()
	if err != nil {
		return &sg.PublicStory{}
	}

	return &sg.PublicStory{
		Id:            s.Id,
		PartyId:       s.PartyId,
		UserId:        s.UserId,
		Lat:           float32(lat),
		Long:          float32(lon),
		Url:           s.Url,
		TaggedFriends: s.TaggedFriends,
		CreatedAt:     timestamppb.New(t),
	}
}
