package cvault

import (
	"errors"
	"os"
)

func CreateSessionEnv() (Session, error) {
	vaultURL, found := os.LookupEnv("VAULT_URL")
	if !found {
		vaultURL = "http://localhost:8200"
	}
	appRoleID, found := os.LookupEnv("APPROLE_ID")
	if !found {
		return nil, errors.New("Env-Variable 'APPROLE_ID' could not be found")
	}
	appRoleSecret, found := os.LookupEnv("APPROLE_SECRET")
	if !found {
		return nil, errors.New("Env-Variable 'APPROLE_SECRET' could not be found")
	}

	return CreateSession(vaultURL, appRoleID, appRoleSecret)
}
