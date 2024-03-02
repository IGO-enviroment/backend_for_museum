package contract

type VerifyQrCode struct {
	Code string `params:"code" validate:"required,min=8,max=255"`
}
