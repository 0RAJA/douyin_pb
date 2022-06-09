package token

import (
	"testing"
	"time"

	"github.com/0RAJA/douyin/internal/pkg/app/errcode"
	"github.com/0RAJA/douyin/internal/pkg/utils"
	"github.com/stretchr/testify/require"
)

func TestPasetoMaker(t *testing.T) {
	maker, err := NewPasetoMaker([]byte(utils.RandomString(32)))
	require.NoError(t, err)
	require.NotEmpty(t, maker)

	username := utils.RandomString(10)
	duration := time.Minute
	userID := utils.RandomInt(1, 1000)

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)
	token, err := maker.CreateToken(userID, username, duration)

	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.Equal(t, payload.UserName, username)
	require.Equal(t, payload.UserID, userID)
	require.WithinDuration(t, payload.IssuedAt, issuedAt, time.Millisecond)

	require.WithinDuration(t, payload.ExpiredAt, expiredAt, time.Second)
}

func TestMaker(t *testing.T) {
	maker, err := NewPasetoMaker([]byte(utils.RandomString(32)))
	require.NoError(t, err)
	require.NotEmpty(t, maker)
	username := utils.RandomOwner()
	duration := time.Second
	userID := utils.RandomInt(1, 1000)
	token, err := maker.CreateToken(userID, username, duration)
	require.NoError(t, err)
	result, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, result)
	time.Sleep(duration * 2)
	result2, err := maker.VerifyToken(token)
	require.ErrorIs(t, err, errcode.ErrUnauthorizedTokenTimeout)
	require.Empty(t, result2)
}
