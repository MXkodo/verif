package repo

import (
	"github.com/MXkodo/verif/config"
	"github.com/MXkodo/verif/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBRepo struct {
	DB *gorm.DB
}

func NewDB() (*DBRepo, error) {
	db, err := gorm.Open(postgres.Open(config.DBDSN), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&models.User{}, &models.Token{})
	if err != nil {
		return nil, err
	}

	return &DBRepo{DB: db}, nil
}
