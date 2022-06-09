package routing

type group struct {
	User       user
	Comment    comment
	UserFollow userFollow
	Video      video
	UserVideo  userVideo
}

var Group = new(group)
