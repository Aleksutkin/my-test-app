package usecases

import "my-test-app/models"

type getOneUserer interface {
	GetUser(int) (models.User, error)
}

type GetUserUsecase struct {
	repo getOneUserer
}

func NewGetUserUsecase(r getOneUserer) *GetUserUsecase {
	return &GetUserUsecase{r}
}

func (g *GetUserUsecase) GetOneUser(id int) (models.User, error) {
	return g.repo.GetUser(id)
}
