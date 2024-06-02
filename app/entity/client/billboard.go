package entity

type BillboardEvents struct {
	Events []BillboardEvent `json:"events"`
}

type BillboardEvent struct {
	ID          int
	Title       string
	PreviewURL  string
	TypeName    string
	DateStartAt string
	TimeStartAt string
	AgeTag      string
}
