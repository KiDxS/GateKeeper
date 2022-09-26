package web

// Field struct for handling change password requests
type ChangePasswordFields struct {
	CurrentPassword    string `json:"current_password" validate:"required"`
	NewPassword        string `json:"new_password" validate:"required,min=12"`
	ConfirmNewPassword string `json:"confirm_password" validate:"required,min=12,eqfield=NewPassword"`
}

// Field struct for handling the login requests
type LoginFields struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Field struct for handling SSH generation requests
type SSHGenerationFields struct {
	Label    string `json:"label" validate:"required"`
	Password string `json:"password"`
}
