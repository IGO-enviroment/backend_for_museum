package entity

type CreateEmail struct {
	Email string `json:"email"`
}

type VerifyEmail struct {
	Code string `json:"code"`
}
