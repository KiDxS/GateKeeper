package ssh

// Field struct for handling SSH generation requests
type SSHGenerationFields struct {
	Label    string `json:"label" validate:"required"`
	Password string `json:"password"`
}
