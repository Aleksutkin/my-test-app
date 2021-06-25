package usecases

import (
	"my-test-app/forms"
	"my-test-app/models"
)

type putUserer interface {
	PutUser(models.User, int) (models.User, error)
}

type PutUserUsecase struct {
	repo putUserer
}

func NewPutUserUsecase(r putUserer) *PutUserUsecase {
	return &PutUserUsecase{r}
}

func (p *PutUserUsecase) PutUser(gotUser forms.User, id int) (models.User, error) {
	usr := &models.User{
		Id:    id,
		Name:  gotUser.Name,
		Email: gotUser.Email,
	}
	return p.repo.PutUser(*usr, id)
}
