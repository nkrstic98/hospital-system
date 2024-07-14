package services

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt/v5"
	"hospital-system/server/app/dto"
	"hospital-system/server/config"
	"time"
)

type SessionService struct {
	config      config.AuthTokenConfig
	redisClient *redis.Client
}

func NewSessionService(redisClient *redis.Client, config config.AuthTokenConfig) *SessionService {
	return &SessionService{
		redisClient: redisClient,
		config:      config,
	}
}

func (s *SessionService) CreateSession(ctx context.Context, claims dto.TokenClaims) (string, error) {
	tokenExpiration := time.Second * time.Duration(s.config.ExpirationSeconds)

	claims.IssuedAt = &jwt.NumericDate{Time: time.Now()}

	// Create JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(s.config.Key))
	if err != nil {
		return "", fmt.Errorf("failed to sign session token: %w", err)
	}

	// Save token to Redis
	marshalledClaims, err := json.Marshal(claims)
	if err != nil {
		return "", fmt.Errorf("failed to marshal session claims: %w", err)
	}

	if err = s.redisClient.Set(ctx, signedToken, marshalledClaims, tokenExpiration).Err(); err != nil {
		return "", fmt.Errorf("failed to save session token to Redis: %w", err)
	}

	return signedToken, nil
}

func (s *SessionService) GetSession(ctx context.Context, token string) (*dto.TokenClaims, error) {
	claims := dto.TokenClaims{}

	// Get claims from Redis
	claimsStr, err := s.redisClient.Get(ctx, token).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get session token from Redis: %w", err)
	}

	if err = json.Unmarshal([]byte(claimsStr), &claims); err != nil {
		return nil, fmt.Errorf("failed to unmarshal session claims: %w", err)
	}

	return &claims, nil
}

func (s *SessionService) RefreshSession(ctx context.Context, token string) error {
	expiration := time.Second * time.Duration(s.config.ExpirationSeconds)

	if _, err := s.redisClient.Expire(ctx, token, expiration).Result(); err != nil {
		return fmt.Errorf("failed to refresh session token: %w", err)
	}

	return nil
}

func (s *SessionService) DeleteSession(ctx context.Context, token string) error {
	if _, err := s.redisClient.Del(ctx, token).Result(); err != nil {
		return fmt.Errorf("failed to delete session token: %w", err)
	}

	return nil
}
