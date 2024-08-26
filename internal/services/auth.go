package services

import (
	"errors"

	"github.com/MXkodo/verif/internal/repo"
	"github.com/MXkodo/verif/models"
	"github.com/MXkodo/verif/pkg"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo *repo.AuthRepo
}

func NewAuthService(repo *repo.AuthRepo) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) GenerateTokens(userID string, ip string) (string, string, error) {
	accessToken, err := pkg.GenerateJWT(userID, ip)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := pkg.GenerateRandomString(32)
	if err != nil {
		return "", "", err
	}

	refreshTokenHash, err := bcrypt.GenerateFromPassword([]byte(refreshToken), bcrypt.DefaultCost)
	if err != nil {
		return "", "", err
	}

	token := &models.Token{
		UserID:       userID,
		RefreshToken: string(refreshTokenHash),
		IP:           ip,
	}

	if err := s.repo.CreateToken(token); err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *AuthService) RefreshTokens(oldRefreshToken string, ip string) (string, string, error) {
	token, err := s.repo.FindTokenByRefreshToken(oldRefreshToken)
	if err != nil {
		return "", "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(token.RefreshToken), []byte(oldRefreshToken)); err != nil {
		return "", "", errors.New("invalid refresh token")
	}

	if token.IP != ip {
		pkg.SendEmailWarning(token.UserID)
		return "", "", errors.New("IP address has changed")
	}

	return s.GenerateTokens(token.UserID, ip)
}
