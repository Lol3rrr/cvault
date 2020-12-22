package cvault

type ApproleAuth struct {
	RoleID     string
	RoleSecret string
}

func NewApprole(role_id, role_secret string) *ApproleAuth {
	return &ApproleAuth{
		RoleID:     role_id,
		RoleSecret: role_secret,
	}
}

func (a *ApproleAuth) GetLoginData() map[string]interface{} {
	data := map[string]interface{}{
		"role_id":   a.RoleID,
		"secret_id": a.RoleSecret,
	}

	return data
}

func (a *ApproleAuth) GetLoginEndpoint() string {
	return "auth/approle/login"
}
