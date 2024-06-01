package user

import "time"

type User struct {
	ID             int
	Role           *string
	IsAdmin        bool
	Email          string
	DigestPassword string
	CreatedAt      *time.Time
}
