package models

import "time"

// Коды подтверждения, аунтификация, подтверждения оплаты и т.п.
type Verification struct {
	ID        int
	Code      string
	UntilAt   *time.Time
	Sended    bool
	Expired   bool
	Options   string
	ModelID   int
	ModelType string
	CreatedAt *time.Time
}
