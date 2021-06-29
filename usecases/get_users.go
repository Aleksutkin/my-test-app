package usecases

import "my-test-app/models"

type getAllUserer interface {
	GetUsers() (*models.Users, error)
}

type GetUsersUsecase struct {
	repo getAllUserer
}

func NewGetUsersUsecase(r getAllUserer) *GetUsersUsecase {
	return &GetUsersUsecase{r}
}

func (g *GetUsersUsecase) GetAllUsers() (*models.Users, error) {
	return g.repo.GetUsers()
}
