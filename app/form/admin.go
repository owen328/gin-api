package form

type RegisterAdmin struct {
	Username string `json:"username" binding:"required,min=4,max=12"`
	Password string `json:"password" binding:"required,min=6,max=16"`
	ConfirmPassword string `json:"confirmPassword" binding:"required,min=6,max=16"`
	Mobile string `json:"mobile" binding:"required,min=11,max=12"`
}

type LoginAdmin struct {
	Username string `json:"username" binding:"required,min=4,max=12"`
	Password string `json:"password" binding:"required,min=6,max=16"`
}