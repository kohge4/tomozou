package datastore

import (
	"tomozou/domain"

	"github.com/jinzhu/gorm"
)

type UserDBRepository struct {
	DB *gorm.DB
}

func NewUserDBRepository(db *gorm.DB) domain.UserRepository {
	return &UserDBRepository{
		DB: db,
	}
}

func NewDevUserRepo(db *gorm.DB) *UserDBRepository {
	return &UserDBRepository{
		DB: db,
	}
}

func (repo *UserDBRepository) Save(user domain.User) (int, error) {
	repo.DB.Create(&user)
	return user.ID, nil
}

func (repo *UserDBRepository) Update(id int) (domain.User, error) {
	user := domain.User{}
	repo.DB.Where("ID = ?", id).Find(&user)
	return user, nil
}

func (repo *UserDBRepository) ReadAll() []domain.User {
	users := []domain.User{}
	repo.DB.Find(&users)
	return users
}

func (repo *UserDBRepository) ReadBySocialID(socialID string) ([]domain.User, error) {
	users := []domain.User{}
	repo.DB.Where("social_id = ?", socialID).Find(&users)
	return users, nil
}

func (repo *UserDBRepository) ReadByID(id int) (domain.User, error) {
	user := domain.User{}
	repo.DB.Where("ID = ?", id).Find(&user)
	return user, nil
}

// 以下開発時の デバッグ用

func (repo *UserDBRepository) CheckUser() (interface{}, error) {
	user := []domain.User{}
	repo.DB.Find(&user)
	return user, nil
}
