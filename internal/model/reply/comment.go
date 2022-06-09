package reply

import (
	"time"
)

type CommentAction struct {
	ID         int64     `json:"id"`
	Content    string    `json:"content"`
	CreateDate time.Time `json:"create_date"`
	User       User      `json:"user"`
}

type CommentList struct {
	CommentList []CommentAction `json:"comment_list,omitempty"`
}

func (p CommentList) Len() int {
	return len(p.CommentList)
}

func (p CommentList) Less(i, j int) bool {
	return p.CommentList[i].CreateDate.Before(p.CommentList[j].CreateDate)
}

func (p CommentList) Swap(i, j int) {
	p.CommentList[i], p.CommentList[j] = p.CommentList[j], p.CommentList[i]
}
