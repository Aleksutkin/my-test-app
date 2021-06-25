package usecases

import (
	"my-test-app/forms"
	"my-test-app/models"
)

type UpdateUserer interface {
	UpdateUser(models.User, int) (models.User, error)
}

type UpdateUserUsecase struct {
	repo UpdateUserer
}

func NewUpdateUserUsecase(r UpdateUserer) *UpdateUserUsecase {
	return &UpdateUserUsecase{r}
}

func (p *UpdateUserUsecase) UpdateUser(gotUser forms.User, id int) (models.User, error) {
	usr := &models.User{
		Id:    id,
		Name:  gotUser.Name,
		Email: gotUser.Email,
	}
	return p.repo.UpdateUser(*usr, id)
}
