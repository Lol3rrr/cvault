package cvault

func CreateSession(url, roleID, roleSecret string) (Session, error) {
	result := &session{
		URL:        url,
		RoleID:     roleID,
		RoleSecret: roleSecret,
	}

	return result, result.auth()
}
