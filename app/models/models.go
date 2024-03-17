package models

import "time"

const (
	tabel_user      = "users"
	tabel_user_role = "user_roles"
	table_role      = "roles"
)

// Общая модель пользователя.
type User struct {
	ID             int
	Email          string
	DigestPassword string
	IsAdmin        bool
	CreatedAt      *time.Time
}

// Таблица храненеия ролей пользователя.
type UserRole struct {
	ID        int
	UserID    string
	RoleID    string
	CreatedAt *time.Time
}

// Роли доступные пользователям.
type Role struct {
	ID        int
	Named     string
	Propery   string
	CreatedAt *time.Time
}

// Мероприятия.
type Event struct {
	ID      int
	Title   string
	Publish bool

	TicketCount int

	StartAt *time.Time

	Duration int

	AreaID int
	TypeID int

	CreatedAt *time.Time
}

// Теги мероприятия
type EventTag struct {
	ID        int
	EventID   int
	TagID     int
	CreatedAt *time.Time
}

// Теги/категории.
type Tag struct {
	ID        int
	Named     string
	Propery   string
	CreatedAt *time.Time
}

// Все типы событий в музее.
type TypeEvent struct {
	ID        int
	Named     string
	Propery   string
	Publish   bool
	CreatedAt *time.Time
}

// Площадки музея.
type Area struct {
	ID        int
	Named     string
	Propery   string
	Publish   bool
	CreatedAt *time.Time
}

// Контент внутри мероприятия, новости, инструкции и т.п.
type Content struct {
	ID         int
	TypeValue  string
	DataValue  string
	OrderValue int
	ModelID    int
	ModelType  string
	Options    string
	CreatedAt  *time.Time
}

// Сохраненные фильтрации для дальнейшего переиспользования
type Filter struct {
	ID         int
	Named      string
	Property   string
	Publish    bool
	OrderValue int
	UserID     int
	Options    string
	CreatedAt  *time.Time
}

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
