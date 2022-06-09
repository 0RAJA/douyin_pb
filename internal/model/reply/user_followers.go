package reply

import db "github.com/0RAJA/douyin/internal/dao/mysql/sqlc"

type RelationAction struct {
}

type RelationFollowList struct {
	UserList []User `json:"user_list"`
}

type RelationFollowerList struct {
	UserList []db.User `json:"user_list"`
}
