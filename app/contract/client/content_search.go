package contract

// Глобальный поиск по контенту.
type ContentSearch struct {
	TypeSearch string `json:"typeSearch" validate:"max=255"`
	Target     string `json:"target" validate:"required,min=1,max=255"`
}
