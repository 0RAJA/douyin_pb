package v1

type group struct {
	User       user
	Comment    comment
	UserFollow userFollow
	Video      video
}

var Group = new(group)
