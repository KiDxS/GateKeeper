package ssh

// SSHGenerationFields is a struct that is used to store the values coming from the user's request to generate a new SSH key pair.
type SSHGenerationFields struct {
	Label    string `json:"label" validate:"required"`
	Password string `json:"password"`
}
