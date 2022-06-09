package db_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/0RAJA/douyin/internal/dao"
	db "github.com/0RAJA/douyin/internal/dao/mysql/sqlc"
	"github.com/0RAJA/douyin/internal/pkg/utils"
	"github.com/stretchr/testify/require"
)

func TestQueriesCreateUser(t *testing.T) {
	user := db.User{
		Name:     utils.RandomString(30),
		Password: utils.RandomString(30),
	}
	userID, err := dao.Group.Mysql.CreateUser(context.Background(), db.CreateUserParams{
		Name:     user.Name,
		Password: user.Password,
	})
	require.NoError(t, err)
	require.NotZero(t, userID)
	result, err := testQueriesGetUserByUserID(userID)
	require.NoError(t, err)
	require.Equal(t, userID, result.ID)
	require.Equal(t, user.Name, result.Name)
	require.Equal(t, user.Password, result.Password)
	result, err = testQueriesGetUserByUsername(user.Name)
	require.NoError(t, err)
	require.Equal(t, userID, result.ID)
	require.Equal(t, user.Name, result.Name)
	require.Equal(t, user.Password, result.Password)
	userID, err = dao.Group.Mysql.CreateUser(context.Background(), db.CreateUserParams{
		Name:     user.Name,
		Password: user.Password,
	})
	require.Error(t, err)
	require.Zero(t, userID)
}

func testCreateUser(t *testing.T) db.User {
	user := db.User{
		Name:     utils.RandomString(30),
		Password: utils.RandomString(30),
	}
	userID, err := dao.Group.Mysql.CreateUser(context.Background(), db.CreateUserParams{
		Name:     user.Name,
		Password: user.Password,
	})
	require.NoError(t, err)
	require.NotZero(t, userID)
	user.ID = userID
	return user
}

func testQueriesGetUserByUsername(userName string) (db.User, error) {
	return dao.Group.Mysql.GetUserByUsername(context.Background(), userName)
}

func TestQueriesGetUserByUsername(t *testing.T) {
	user := db.User{
		Name: utils.RandomString(30),
	}
	result, err := dao.Group.Mysql.GetUserByUsername(context.Background(), user.Name)
	require.ErrorIs(t, err, sql.ErrNoRows)
	require.Empty(t, result)
	user = testCreateUser(t)
	result, err = dao.Group.Mysql.GetUserByUsername(context.Background(), user.Name)
	require.NoError(t, err)
	require.Equal(t, result, user)
}

func testQueriesGetUserByUserID(userID int64) (db.User, error) {
	return dao.Group.Mysql.GetUserByUserID(context.Background(), userID)
}

func TestQueriesGetUserByUserID(t *testing.T) {
	user := testCreateUser(t)
	result, err := dao.Group.Mysql.GetUserByUserID(context.Background(), user.ID)
	require.NoError(t, err)
	require.Equal(t, result, user)
}

func TestQueries_UpdateUserFollowerCount(t *testing.T) {
	user := testCreateUser(t)
	sum := int64(0)
	type args struct {
		ctx context.Context
		arg db.UpdateUserFollowerCountParams
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "good",
			args: args{
				ctx: context.Background(),
				arg: db.UpdateUserFollowerCountParams{
					FollowerCount: utils.RandomInt(1, 100),
					ID:            user.ID,
				},
			},
		},
		{
			name: "good",
			args: args{
				ctx: context.Background(),
				arg: db.UpdateUserFollowerCountParams{
					FollowerCount: utils.RandomInt(1, 100),
					ID:            user.ID,
				},
			},
		},
		{
			name: "good negativeNumber",
			args: args{
				ctx: context.Background(),
				arg: db.UpdateUserFollowerCountParams{
					FollowerCount: -sum / 2,
					ID:            user.ID,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sum += tt.args.arg.FollowerCount
			err := dao.Group.Mysql.UpdateUserFollowerCount(tt.args.ctx, tt.args.arg)
			require.NoError(t, err)
			user, err := testQueriesGetUserByUserID(tt.args.arg.ID)
			require.NoError(t, err)
			require.Equal(t, user.FollowerCount, sum)
		})
	}
}

func TestQueries_UpdateUserFollowCount(t *testing.T) {
	user := testCreateUser(t)
	sum := int64(0)
	type args struct {
		ctx context.Context
		arg db.UpdateUserFollowCountParams
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "good",
			args: args{
				ctx: context.Background(),
				arg: db.UpdateUserFollowCountParams{
					FollowCount: utils.RandomInt(1, 100),
					ID:          user.ID,
				},
			},
		},
		{
			name: "good",
			args: args{
				ctx: context.Background(),
				arg: db.UpdateUserFollowCountParams{
					FollowCount: utils.RandomInt(1, 100),
					ID:          user.ID,
				},
			},
		},
		{
			name: "good negativeNumber",
			args: args{
				ctx: context.Background(),
				arg: db.UpdateUserFollowCountParams{
					FollowCount: -sum / 2,
					ID:          user.ID,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sum += tt.args.arg.FollowCount
			err := dao.Group.Mysql.UpdateUserFollowCount(tt.args.ctx, tt.args.arg)
			require.NoError(t, err)
			user, err := testQueriesGetUserByUserID(tt.args.arg.ID)
			require.NoError(t, err)
			require.Equal(t, user.FollowCount, sum)
		})
	}
}
