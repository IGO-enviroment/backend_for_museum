/*
 * Модель пользователя
 */
package users

type (
	User struct {
		ID    int
		Email string

		CreatedAt string
		UpdatedAt string

		Attr  *UserAttr
		Query *UserQuery
	}
)

func NewUser(opts ...UserOpts) *User {
	user := &User{
		Attr: &UserAttr{}, Query: NewQuery(),
	}
	for _, opt := range opts {
		opt(user)
	}
	return user
}
