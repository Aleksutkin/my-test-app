package usecases

import (
	"my-test-app/forms"
	"my-test-app/models"
)

type createUserer interface {
	CreateUser(models.User) (*models.User, error)
}

type CreateUserUsecase struct {
	repo createUserer
}

func NewCreateUserUsecase(r createUserer) *CreateUserUsecase {
	return &CreateUserUsecase{r}
}

func (p *CreateUserUsecase) CreateUser(gotUser forms.User) (*models.User, error) {
	usr := &models.User{
		Name:  gotUser.Name,
		Email: gotUser.Email,
	}
	return p.repo.CreateUser(*usr)
}
