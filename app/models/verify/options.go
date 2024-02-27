package verify

type VerifyOpts func(*Verify)

type VerifyAttr struct {
	u *Verify
}

func (a *VerifyAttr) ID(id int) VerifyOpts {
	return func(u *Verify) {
		u.ID = id
	}
}

func (a *VerifyAttr) MapField(u *Verify) []interface{} {
	return []interface{}{
		&u.ID,
		&u.ModelID,
	}
}

func (a *VerifyAttr) SelectMin() string {
	return "id, model_id, model_name"
}
