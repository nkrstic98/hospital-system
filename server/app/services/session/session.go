package session

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/exp/slog"
	"hospital-system/server/app/dto"
	"hospital-system/server/config"
	"time"
)

type ServiceImpl struct {
	config      config.Config
	redisClient *redis.Client
}

func NewService(redisClient *redis.Client, config config.Config) *ServiceImpl {
	return &ServiceImpl{
		redisClient: redisClient,
		config:      config,
	}
}

func (service *ServiceImpl) CreateSession(claims dto.TokenClaims) (string, error) {
	tokenExpiration := time.Second * time.Duration(service.config.AuthToken.ExpirationSeconds)

	claims.IssuedAt = &jwt.NumericDate{Time: time.Now()}

	// Create JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(service.config.AuthToken.Key))
	if err != nil {
		slog.Error("Error signing token: ", err)
		return "", err
	}

	// Save token to Redis
	marshalledClaims, err := json.Marshal(claims)
	if err != nil {
		slog.Error("Error marshalling claims: ", err)
		return "", err
	}

	if err = service.redisClient.Set(context.Background(), signedToken, marshalledClaims, tokenExpiration).Err(); err != nil {
		slog.Error("Error saving token to Redis: ", err)
		return "", err
	}

	return signedToken, err
}

func (service *ServiceImpl) GetSession(token string) (*dto.TokenClaims, error) {
	claims := dto.TokenClaims{}

	// Get claims from Redis
	claimsStr, err := service.redisClient.Get(context.Background(), token).Result()
	if err != nil {
		slog.Error("Error getting claims from Redis: ", err)
		return nil, err
	}

	if err = json.Unmarshal([]byte(claimsStr), &claims); err != nil {
		slog.Error("Error unmarshalling claims: ", err)
		return nil, err
	}

	return &claims, nil
}

func (service *ServiceImpl) RefreshSession(token string) error {
	expiration := time.Second * time.Duration(service.config.AuthToken.ExpirationSeconds)

	if _, err := service.redisClient.Expire(context.Background(), token, expiration).Result(); err != nil {
		slog.Error("Error refreshing session: ", err)
		return err
	}

	return nil
}

func (service *ServiceImpl) DeleteSession(token string) error {
	if _, err := service.redisClient.Del(context.Background(), token).Result(); err != nil {
		slog.Error("Error deleting session: ", err)
		return err
	}

	return nil
}
