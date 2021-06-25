package forms

type User struct {
	Name  string `json:"name" validate:"required,max=50"`
	Email string `json:"email" validate:"omitempty,email"`
}
