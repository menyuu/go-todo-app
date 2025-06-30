package forms

type UserForm struct {
	Name     string `form:"name" validate:"required,min=2,max=20"`
	Email    string `form:"email" validate:"required,email"`
	Password string `form:"password" validate:"min=6,max=50"`
}
