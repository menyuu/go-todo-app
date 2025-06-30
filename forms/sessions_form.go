package forms

type SignUpForm struct {
	Name     string `form:"name" validate:"required,min=2,max=20"`
	Email    string `form:"email" validate:"required,email"`
	Password string `form:"password" validate:"required,min=6,max=20"`
}

type LoginForm struct {
	Email    string `form:"email" validate:"required,email"`
	Password string `form:"password" validate:"required"`
}
