package repository

import (
	"context"

	domain "github.com/linn-phyo/go_gin_clean_architecture/src/domain"
	interfaces "github.com/linn-phyo/go_gin_clean_architecture/src/repository/interface"
	"gorm.io/gorm"
)

type userDatabase struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &userDatabase{DB}
}

func (c *userDatabase) FindAll(ctx context.Context) ([]domain.Users, error) {
	var users []domain.Users
	err := c.DB.Find(&users).Error

	//log.Fatal("USERS >> ", users)

	return users, err
}

func (c *userDatabase) FindByID(ctx context.Context, id string) (domain.Users, error) {
	var user domain.Users
	// err := c.DB.First(&user, id).Error
	err := c.DB.Where("id = ?", id).First(&user).Error
	return user, err
}

func (c *userDatabase) Save(ctx context.Context, user domain.Users) (domain.Users, error) {
	err := c.DB.Save(&user).Error

	return user, err
}

func (c *userDatabase) Delete(ctx context.Context, user domain.Users) error {
	err := c.DB.Delete(&user).Error

	return err
}
