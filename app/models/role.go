package models

type Role struct {
	TabelName string

	name string
}

func NewRole() *Role {
	return &Role{TabelName: "roles"}
}
