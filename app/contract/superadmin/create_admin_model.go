package superadmin

type CreateAdminModel struct {
	Email        string `json:"email"`
	IsSuperAdmin bool   `json:"isSuperAdmin"`
}
