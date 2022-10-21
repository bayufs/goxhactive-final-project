package resources

type InputUsers struct {
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=6"`
	Age      int    `json:"age" validate:"required,gte=10"`
}

type InputUserUpdate struct {
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required"`
}
