package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/Asqar95/crud-app/internal/domain"
	audit "github.com/Asqar95/crud-audit-log/pkg/domain"
	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)

// PasswordHasher provides hashing logic to securely store passwords.
type PasswordHasher interface {
	Hash(password string) (string, error)
}

type UsersRepository interface {
	Create(ctx context.Context, user domain.User) error
	GetByCredentials(ctx context.Context, email, password string) (domain.User, error)
}

type SessionsRepository interface {
	Create(ctx context.Context, token domain.RefreshSession) error
	Get(ctx context.Context, token string) (domain.RefreshSession, error)
}

type AuditClient interface {
	SendLogRequest(ctx context.Context, req audit.LogItem) error
}

type Users struct {
	repo         UsersRepository
	hasher       PasswordHasher
	sessionsRepo SessionsRepository
	auditClient  AuditClient
	hmacSecret   []byte
	tokenTtl     time.Duration
}

func NewUsers(repo UsersRepository, sessionsRepo SessionsRepository, auditClient AuditClient, hasher PasswordHasher, secret []byte) *Users {
	return &Users{
		repo:         repo,
		sessionsRepo: sessionsRepo,
		hasher:       hasher,
		auditClient:  auditClient,
		hmacSecret:   secret,
	}
}

func (s *Users) SignUp(ctx context.Context, inp domain.SignUpInput) error {
	password, err := s.hasher.Hash(inp.Password)
	if err != nil {
		return err
	}

	user := domain.User{
		Name:         inp.Name,
		Email:        inp.Email,
		Password:     password,
		RegisteredAt: time.Now(),
	}

	if err := s.repo.Create(ctx, user); err != nil {
		return err
	}

	user, err = s.repo.GetByCredentials(ctx, inp.Email, password)
	if err != nil {
		return err
	}

	if err := s.auditClient.SendLogRequest(ctx, audit.LogItem{
		Action:    audit.ACTION_REGISTER,
		Entity:    audit.ENTITY_USER,
		EntityID:  user.ID,
		Timestamp: time.Now(),
	}); err != nil {
		logrus.WithFields(logrus.Fields{
			"method": "Users.SignUp",
		}).Error("failed to send log request:", err)
	}

	return nil
}

func (s *Users) SignIn(ctx context.Context, inp domain.SignInInput) (string, error) {
	password, err := s.hasher.Hash(inp.Password)
	if err != nil {
		return "", err
	}

	user, err := s.repo.GetByCredentials(ctx, inp.Email, password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", domain.ErrUserNotFound
		}

		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Subject:   strconv.Itoa(int(user.ID)),
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(s.tokenTtl).Unix(),
	})

	return token.SignedString(s.hmacSecret)
}

func (s *Users) ParseToken(ctx context.Context, token string) (int64, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return s.hmacSecret, nil
	})
	if err != nil {
		return 0, err
	}

	if !t.Valid {
		return 0, errors.New("invalid token")
	}

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid claims")
	}

	subject, ok := claims["sub"].(string)
	if !ok {
		return 0, errors.New("invalid subject")
	}

	id, err := strconv.Atoi(subject)
	if err != nil {
		return 0, errors.New("invalid subject")
	}

	return int64(id), nil
}
