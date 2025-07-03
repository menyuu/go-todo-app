package forms

type TodoForm struct {
	Title string `form:"title" validate:"required,max=50"`
	Done  bool
}
