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

func TestQueries_CreateUserFollower(t *testing.T) {
	user1 := testCreateUser(t)
	user2 := testCreateUser(t)
	type args struct {
		ctx context.Context
		arg db.CreateUserFollowerParams
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		err     error
	}{
		{
			name: "good a to b",
			args: args{
				ctx: context.Background(),
				arg: db.CreateUserFollowerParams{
					FromID: user1.ID,
					ToID:   user2.ID,
				},
			},
		},
		{
			name: "bad a to a",
			args: args{
				ctx: context.Background(),
				arg: db.CreateUserFollowerParams{
					FromID: user1.ID,
					ToID:   user2.ID,
				},
			},
			wantErr: true,
		},
		{
			name: "good b to a",
			args: args{
				ctx: context.Background(),
				arg: db.CreateUserFollowerParams{
					FromID: user2.ID,
					ToID:   user2.ID,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := dao.Group.Mysql.CreateUserFollower(tt.args.ctx, tt.args.arg)
			if tt.wantErr {
				require.Error(t, err)
				if tt.err != nil {
					require.Equal(t, err, tt.err)
				}
			} else {
				require.NoError(t, err)
				_, err := dao.Group.Mysql.GetUserFollower(tt.args.ctx, db.GetUserFollowerParams{
					FromID: tt.args.arg.FromID,
					ToID:   tt.args.arg.ToID,
				})
				require.NoError(t, err)
			}
		})
	}
}

func testCreateUserFollower(t *testing.T) (fromID, toID int64) {
	user1 := testCreateUser(t)
	user2 := testCreateUser(t)
	err := dao.Group.Mysql.CreateUserFollower(context.Background(), db.CreateUserFollowerParams{
		FromID: user1.ID,
		ToID:   user2.ID,
	})
	require.NoError(t, err)
	return user1.ID, user2.ID
}

func TestQueries_DeleteUserFollower(t *testing.T) {
	fromID, toID := testCreateUserFollower(t)
	type args struct {
		ctx context.Context
		arg db.DeleteUserFollowerParams
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		f       func()
	}{
		{
			name: "good",
			args: args{
				ctx: context.Background(),
				arg: db.DeleteUserFollowerParams{
					FromID: fromID,
					ToID:   toID,
				},
			},
			f: func() {
				result, err := dao.Group.Mysql.GetUserFollower(context.Background(), db.GetUserFollowerParams{
					FromID: fromID,
					ToID:   toID,
				})
				require.Error(t, err)
				require.Empty(t, result)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := dao.Group.Mysql.DeleteUserFollower(tt.args.ctx, tt.args.arg)
			if tt.wantErr {
				require.Error(t, err)
			}
			tt.f()
		})
	}
}

func TestQueries_GetUserFollower(t *testing.T) {
	fromID, toID := testCreateUserFollower(t)
	type args struct {
		ctx context.Context
		arg db.GetUserFollowerParams
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		err     error
		f       func(result db.UserFollower)
	}{
		{
			name: "good",
			args: args{
				ctx: context.Background(),
				arg: db.GetUserFollowerParams{
					FromID: fromID,
					ToID:   toID,
				},
			},
			f: func(result db.UserFollower) {
				require.Equal(t, result.FromID, fromID)
				require.Equal(t, result.ToID, toID)
			},
		},
		{
			name: "bad not found",
			args: args{
				ctx: context.Background(),
				arg: db.GetUserFollowerParams{
					FromID: toID,
					ToID:   fromID,
				},
			},
			wantErr: true,
			err:     sql.ErrNoRows,
			f:       func(result db.UserFollower) {},
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := dao.Group.Mysql.GetUserFollower(tt.args.ctx, tt.args.arg)
			if tt.wantErr {
				require.ErrorIs(t, err, tt.err)
				require.Empty(t, got)
			} else {
				require.NoError(t, err)
				tt.f(got)
			}
		})
	}
}

func TestQueries_GetUserFollowersByFromUserID(t *testing.T) {
	user1 := testCreateUser(t)
	type args struct {
		ctx    context.Context
		fromID int64
	}
	type Test struct {
		name string
		args args
		f    func(Test)
	}
	tests := []struct {
		name string
		args args
		f    func(Test)
	}{
		{
			name: "good",
			args: args{
				ctx:    context.Background(),
				fromID: user1.ID,
			},
			f: func(tt Test) {
				nums := utils.RandomInt(1, 100)
				for i := 0; i < int(nums); i++ {
					user2 := testCreateUser(t)
					require.NoError(t, dao.Group.Mysql.CreateUserFollower(context.Background(), db.CreateUserFollowerParams{
						FromID: tt.args.fromID,
						ToID:   user2.ID,
					}))
				}
				results, err := dao.Group.Mysql.GetUserFollowersByFromUserID(tt.args.ctx, tt.args.fromID)
				require.NoError(t, err)
				require.Len(t, results, int(nums))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.f(tt)
		})
	}
}

func TestQueries_GetUserFollowersByToUserID(t *testing.T) {
	user1 := testCreateUser(t)
	type args struct {
		ctx  context.Context
		toID int64
	}
	type Test struct {
		name string
		args args
		f    func(Test)
	}
	tests := []struct {
		name string
		args args
		f    func(Test)
	}{
		{
			name: "good",
			args: args{
				ctx:  context.Background(),
				toID: user1.ID,
			},
			f: func(tt Test) {
				nums := utils.RandomInt(1, 100)
				for i := 0; i < int(nums); i++ {
					user2 := testCreateUser(t)
					require.NoError(t, dao.Group.Mysql.CreateUserFollower(context.Background(), db.CreateUserFollowerParams{
						FromID: user2.ID,
						ToID:   tt.args.toID,
					}))
				}
				results, err := dao.Group.Mysql.GetUserFollowersByToUserID(tt.args.ctx, tt.args.toID)
				require.NoError(t, err)
				require.Len(t, results, int(nums))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.f(tt)
		})
	}
}
