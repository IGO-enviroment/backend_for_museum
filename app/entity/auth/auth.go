package entity

type SignUpEntity struct {
	Email           string `json:"email" validate:"required,email,min=3,max=255"`
	Password        string `json:"password" validate:"required,min=8,max=255"`
	PasswordConfirm string `json:"password_confir" validate:"required,eqfield=Password"`
}

func (e *SignUpEntity) Error() string {
	return "te"
}
