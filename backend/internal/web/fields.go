package web

type ChangePasswordFields struct {
	CurrentPassword    string `json:"current_password" validate:"required"`
	NewPassword        string `json:"new_password" validate:"required,min=12"`
	ConfirmNewPassword string `json:"confirm_password" validate:"required,min=12,eqfield=NewPassword"`
}

type LoginFields struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
