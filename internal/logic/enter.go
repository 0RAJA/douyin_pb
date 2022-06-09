package logic

type group struct {
	User       user
	Comment    comment
	UserFollow userFollow
	Video      video
	UserVideo  uservideo
}

var Group = new(group)
