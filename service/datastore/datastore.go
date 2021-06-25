package datastore

import (
	"my-test-app/models"

	"github.com/go-pg/pg/v10"
)

type UserRepository struct {
	db  *pg.DB
	rep models.Users
}

func NewUserRepository(db *pg.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) GetUsers() (models.Users, error) {
	users := make(models.Users, 0)
	err := u.db.Model(&users).
		Select()
	return users, err
}

func (u *UserRepository) GetUser(id int) (models.User, error) {
	var user models.User
	err := u.db.Model(&user).
		Where("id = ?", id).
		Select()
	return user, err
}

func (u *UserRepository) CreateUser(gotUser models.User) (models.User, error) {
	_, err := u.db.Model(&gotUser).
		Insert()
	return gotUser, err
}

func (u *UserRepository) UpdateUser(gotUser models.User, id int) (models.User, error) {
	res, err := u.db.Model(&gotUser).
		Where("id = ?", id).
		Update()
	if res.RowsAffected() == 0 {
		return models.User{}, pg.ErrNoRows
	}

	return gotUser, err
}

func (u *UserRepository) DeleteUser(id int) (models.User, error) {
	var user models.User
	res, err := u.db.Model(&user).
		Where("id = ?", id).
		Delete()
	if res.RowsAffected() == 0 {
		return models.User{}, pg.ErrNoRows
	}

	return user, err
}
