package dto

type Story struct {
	Id            string
	PartyId       string
	UserId        string
	Url           string
	TaggedFriends []string
}
