package usecases

import "my-test-app/models"

type deleteUserer interface {
	DeleteUser(int) (*models.User, error)
}

type DeleteUserUsecase struct {
	repo deleteUserer
}

func NewDeleteUserUsecase(r deleteUserer) *DeleteUserUsecase {
	return &DeleteUserUsecase{r}
}

func (d *DeleteUserUsecase) DeleteUser(id int) (*models.User, error) {
	return d.repo.DeleteUser(id)
}
