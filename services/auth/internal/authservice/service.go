package authservice

import (
	"context"
	"sync"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

type AuthService interface {
	GenerateToken(ctx context.Context, id int64, username, email, role string) (token string, err error)
	ParseToken(ctx context.Context, token string) (claims Claims, err error)
	DeleteToken(ctx context.Context, token string) (valid bool, err error)
}

type authservice struct {
	mtx    sync.RWMutex
	db     *redis.Client
	logger *zap.Logger
}

func NewAuthService(db *redis.Client) AuthService {
	var svc AuthService
	{
		svc = newBasisService(db)
	}
	return svc
}

func newBasisService(db *redis.Client) AuthService {
	return &authservice{
		db: db,
	}
}

type Claims struct {
	Exp        int64  `json:"exp"`
	Created_at int64  `json:"created_at"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Role       string `json:"role"`
}

func (as authservice) GenerateToken(ctx context.Context, id int64, username, email, role string) (token string, err error) {
	as.mtx.RLock()
	defer as.mtx.RUnlock()
}

func (as authservice) ParseToken(ctx context.Context, token string) (claims Claims, err error) {
	as.mtx.RLock()
	defer as.mtx.RUnlock()
}

func (as authservice) DeleteToken(ctx context.Context, token string) (valid bool, err error) {
	as.mtx.RLock()
	defer as.mtx.RUnlock()
}
