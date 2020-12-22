package cvault

func CreateSession(url string, auth Auth) (Session, error) {
	result := &session{
		URL:      url,
		AuthData: auth,
	}

	return result, result.Auth()
}
