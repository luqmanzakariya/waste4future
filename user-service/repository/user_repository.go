package repository

import (
	"context"
	"user-service/model/domain"

	"gorm.io/gorm"
)

type IUserRepository interface {
	FindAll(ctx context.Context) ([]domain.User, error)
	FindById(ctx context.Context, id int) (domain.User, error)
	FindByEmail(ctx context.Context, email string) (domain.User, error)
	Create(ctx context.Context, user domain.User) (domain.User, error)
	Update(ctx context.Context, user domain.User) (domain.User, error)
	Delete(ctx context.Context, user domain.User) error
}

type userRepository struct {
	Db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{
		Db: db,
	}
}

func (u *userRepository) FindAll(ctx context.Context) ([]domain.User, error) {
	var users []domain.User

	// res := u.Db.Table("Users").Find(&users) // Specify the table name || If your table name is Users (uppercase):
	res := u.Db.Find(&users)
	if res.Error != nil {
		return []domain.User{}, res.Error
	}

	return users, res.Error
}

func (u *userRepository) FindById(ctx context.Context, id int) (domain.User, error) {
	var user domain.User

	res := u.Db.First(&user, id)

	return user, res.Error
}

func (u *userRepository) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	var user domain.User

	err := u.Db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *userRepository) Create(ctx context.Context, user domain.User) (domain.User, error) {
	res := u.Db.Create(&user)

	if res.Error != nil {
		return domain.User{}, res.Error
	}

	return user, nil
}

func (u *userRepository) Update(ctx context.Context, user domain.User) (domain.User, error) {
	res := u.Db.Model(&user).Where("user_id = ?", user.ID).Updates(map[string]interface{}{
		"name":       user.Name,
		"email":      user.Email,
		"password":   user.Password,
		"address_id": user.AddressID,
	})

	if res.Error != nil {
		return domain.User{}, res.Error
	}

	return user, nil
}

func (s *userRepository) Delete(ctx context.Context, user domain.User) error {
	res := s.Db.Delete(&domain.User{}, user)
	return res.Error
}
