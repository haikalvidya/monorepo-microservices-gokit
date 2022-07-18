package authservice

import (
	"sync"

	"github.com/go-kit/kit/metrics"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

type AuthService interface {
	// GenerateToken(ctx context.Context, id int64, username, email, role string) (token string, err error)
	// ParseToken(ctx context.Context, token string) (claims Claims, err error)
	// DeleteToken(ctx context.Context, token string) (valid bool, err error)
}

type authservice struct {
	mtx    sync.RWMutex
	db     *redis.Client
	logger *zap.Logger
}

func NewAuthService(db *redis.Client, logger *zap.Logger, ints metrics.Counter) {
	var svc AuthService
	{
		svc = newBasisService(db, logger)
		svc = LoggingMiddleware(logger)(svc)
		svc = InstrumentingMiddleware(ints)(svc)
	}
	return svc
}

func newBasisService(db *redis.Client, logger *zap.Logger) AuthService {
	return &authservice{
		db:     db,
		logger: logger,
	}
}

type Claims struct {
	Exp        int64  `json:"exp"`
	Created_at int64  `json:"created_at"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Role       string `json:"role"`
}

// func (as authservice) GenerateToken(ctx context.Context, id int64, username, email, role string) (token string, err error) {

// }
