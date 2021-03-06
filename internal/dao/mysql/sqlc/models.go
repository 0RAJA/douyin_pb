// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package db

import (
	"time"
)

type Comment struct {
	ID         int64     `json:"id"`
	UserID     int64     `json:"user_id"`
	VideoID    int64     `json:"video_id"`
	Content    string    `json:"content"`
	CreateDate time.Time `json:"create_date"`
}

type User struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	Password      string `json:"password"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
}

type UserFollower struct {
	ID     int64 `json:"id"`
	FromID int64 `json:"from_id"`
	ToID   int64 `json:"to_id"`
}

type UserVideo struct {
	ID      int64 `json:"id"`
	UserID  int64 `json:"user_id"`
	VideoID int64 `json:"video_id"`
}

type Video struct {
	ID            int64     `json:"id"`
	UserID        int64     `json:"user_id"`
	Title         string    `json:"title"`
	PlayUrl       string    `json:"play_url"`
	CoverUrl      string    `json:"cover_url"`
	FavoriteCount int64     `json:"favorite_count"`
	CommentCount  int64     `json:"comment_count"`
	CreatedAt     time.Time `json:"created_at"`
}
