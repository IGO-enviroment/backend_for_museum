package roles

type Role struct {
	name string
}

func NewRole() *Role {
	return &Role{}
}
