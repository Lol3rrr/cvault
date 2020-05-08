package cvault

import "github.com/hashicorp/vault/api"

// Session is used to interact with Vault
type Session interface {
	Auth() error

	ReadData(name string) (*api.Secret, error)
	Renew(id string, interval int) (*api.Secret, error)
	RenewDataForever(secret *api.Secret, interval int, callback onFailedRenew)

	WriteMapData(name string, data map[string]interface{}) error
}

type session struct {
	URL        string
	RoleID     string
	RoleSecret string
	Client     *api.Client
}

type onFailedRenew func()
