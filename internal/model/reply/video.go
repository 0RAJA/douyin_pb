package reply

type Video struct {
	ID            int64  `json:"id"`
	Author        User   `json:"author"`
	Title         string `json:"title"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	FavoriteCount int64  `json:"favorite_count"`
	CommentCount  int64  `json:"comment_count"`
	IsFavorite    bool   `json:"is_favorite"`
}

type Feed struct {
	NextTime  int64   `json:"next_time"`
	VideoList []Video `json:"video_list"`
}

type PublishAction struct {
}

type PublishList struct {
	VideoList []Video `json:"video_list"`
}
