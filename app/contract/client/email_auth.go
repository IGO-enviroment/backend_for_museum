package contract

type CreateEmail struct {
	Email string `json:"email" validate:"required,email,min=3,max=255"`
}

type VerifyEmail struct {
	Code string `json:"code" validate:"required,min=1,max=255"`
}
