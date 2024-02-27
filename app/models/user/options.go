package users

type UserOpts func(*User)

type UserAttr struct {
	u *User
}

func (a *UserAttr) ID(id int) UserOpts {
	return func(u *User) {
		u.ID = id
	}
}

func (a *UserAttr) MapField(u *User) []interface{} {
	return []interface{}{
		&u.ID,
		&u.Email,
	}
}

func (a *UserAttr) SelectMin() string {
	return "id, email"
}
