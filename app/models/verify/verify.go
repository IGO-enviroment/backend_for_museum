/*
 * Модель пользователя
 */
package verify

type (
	Verify struct {
		ID        int
		ModelID   int
		ModelName string

		UniqCode string
		QRCode   string

		CreatedAt string
		UpdatedAt string

		Attr  *VerifyAttr
		Query *VerifyQuery
	}
)

func NewVerify(opts ...VerifyOpts) *Verify {
	verify := &Verify{
		Attr:  &VerifyAttr{},
		Query: NewQuery(),
	}
	for _, opt := range opts {
		opt(verify)
	}
	return verify
}
