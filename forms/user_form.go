package forms

type User struct {
	Name  string `json:"name" validate:"required,max=5"`
	Email string `json:"email" validate:"omitempty,email"`
}
