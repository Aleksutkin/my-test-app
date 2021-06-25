package usecases

import (
	"my-test-app/forms"
	"my-test-app/models"
)

type postUserer interface {
	PostUser(models.User) (models.User, error)
}

type PostUserUsecase struct {
	repo postUserer
}

func NewPostUserUsecase(r postUserer) *PostUserUsecase {
	return &PostUserUsecase{r}
}

func (p *PostUserUsecase) PostUser(gotUser forms.User) (models.User, error) {
	usr := &models.User{
		Name:  gotUser.Name,
		Email: gotUser.Email,
	}
	return p.repo.PostUser(*usr)
}
