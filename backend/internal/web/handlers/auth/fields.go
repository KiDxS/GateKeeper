package auth

// ChangePasswordFields is a struct that is used to store the user's request to change their password.
type ChangePasswordFields struct {
	CurrentPassword    string `json:"current_password" validate:"required"`
	NewPassword        string `json:"new_password" validate:"required,min=12"`
	ConfirmNewPassword string `json:"confirm_password" validate:"required,min=12,eqfield=NewPassword"`
}

// LoginFields struct is a struct that is used to store the user's request to login to their account.
type LoginFields struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
