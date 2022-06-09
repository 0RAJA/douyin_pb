package token

import (
	"time"

	"github.com/0RAJA/douyin/internal/pkg/app/errcode"
	"github.com/o1egl/paseto"
	"golang.org/x/crypto/chacha20poly1305"
)

type PasetoMaker struct {
	paseto *paseto.V2
	key    []byte
}

func NewPasetoMaker(key []byte) (Maker, error) {
	if len(key) != chacha20poly1305.KeySize {
		return nil, ErrSecretLen
	}
	return &PasetoMaker{
		paseto: paseto.NewV2(),
		key:    key,
	}, nil
}

func (p *PasetoMaker) CreateToken(userID int64, username string, expireDate time.Duration) (string, error) {
	paload, err := NewPayload(userID, username, expireDate)
	if err != nil {
		return "", nil
	}
	token, err := p.paseto.Encrypt(p.key, paload, nil)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (p *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	paload := &Payload{}
	err := p.paseto.Decrypt(token, p.key, paload, nil)
	if err != nil {
		return nil, err
	}
	if paload.ExpiredAt.Before(time.Now()) {
		return nil, errcode.ErrUnauthorizedTokenTimeout
	}
	return paload, nil
}
