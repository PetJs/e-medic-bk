// Package auth contains authentication use cases.
package auth

import (
	"context"
	"errors"
	"time"

	"emedic-bk/internal/application/dto"
	"emedic-bk/internal/application/port"
	"emedic-bk/internal/domain/entity"
	"emedic-bk/internal/domain/repository"
)

// ErrEmailAlreadyExists is returned when the email is already registered.
var ErrEmailAlreadyExists = errors.New("email already exists")

// RegisterUseCase handles user registration.
type RegisterUseCase struct {
	userRepo repository.UserRepository
	hasher   port.Hasher
	tokenGen port.TokenGenerator
	idGen    port.IDGenerator
}

// NewRegisterUseCase creates a new RegisterUseCase.
func NewRegisterUseCase(
	userRepo repository.UserRepository,
	hasher port.Hasher,
	tokenGen port.TokenGenerator,
	idGen port.IDGenerator,
) *RegisterUseCase {
	return &RegisterUseCase{
		userRepo: userRepo,
		hasher:   hasher,
		tokenGen: tokenGen,
		idGen:    idGen,
	}
}

// Execute registers a new user.
func (uc *RegisterUseCase) Execute(ctx context.Context, req *dto.RegisterRequest) (*dto.AuthResponse, error) {
	// Check if email already exists
	existingUser, err := uc.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, ErrEmailAlreadyExists
	}

	// Hash password
	passwordHash, err := uc.hasher.Hash(req.Password)
	if err != nil {
		return nil, err
	}

	// Create user entity
	now := time.Now()
	user := &entity.User{
		ID:           uc.idGen.Generate(),
		Email:        req.Email,
		PasswordHash: passwordHash,
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		Role:         "student",
		IsActive:     true,
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	// Save user
	if err := uc.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}

	// Generate tokens
	accessToken, expiresIn, err := uc.tokenGen.GenerateAccessToken(user.ID, user.Role)
	if err != nil {
		return nil, err
	}

	refreshToken, err := uc.tokenGen.GenerateRefreshToken(user.ID)
	if err != nil {
		return nil, err
	}

	return &dto.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    expiresIn,
		User: dto.UserResponse{
			ID:        user.ID,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Role:      user.Role,
			CreatedAt: user.CreatedAt,
		},
	}, nil
}
