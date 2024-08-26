package repo

import (
	"github.com/MXkodo/verif/models"
	"gorm.io/gorm"
)

type AuthRepo struct {
	db *gorm.DB
}

func New(dbRepo *DBRepo) *AuthRepo {
	return &AuthRepo{db: dbRepo.DB}
}

func (repo *AuthRepo) CreateToken(token *models.Token) error {
	return repo.db.Create(token).Error
}

func (repo *AuthRepo) FindTokenByRefreshToken(refreshToken string) (*models.Token, error) {
	var token models.Token
	if err := repo.db.Where("refresh_token = ?", refreshToken).First(&token).Error; err != nil {
		return nil, err
	}
	return &token, nil
}
