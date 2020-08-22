package cvault

import "github.com/hashicorp/vault/api"

// Session is used to interact with Vault
type Session interface {
	// Auth authenticates the current Session with vault
	// and automatically renews the Session afterwards forever
	Auth() error

	// ReadMap is used to read Data from vault and convert it
	// to a map which makes it easier to work with
	ReadMap(name string) (map[string]interface{}, error)
	// ReadData is used to read the Data from vault for the
	// given path
	ReadData(name string) (*api.Secret, error)
	// Renew simply renews the Token/Lease with the given id
	// for the given amount of time
	Renew(id string, interval int) (*api.Secret, error)
	// RenewDataForever is used to Renew the secret forever and
	// extends the lease by the given interval
	// calls the callback in case of an error
	RenewDataForever(secret *api.Secret, interval int, callback onFailedRenew)

	// WriteMapData writes the given data for the name to vault
	WriteMapData(name string, data map[string]interface{}) error

	// GetVaultClient returns the underlying client used by
	// the given cvault session
	GetVaultClient() *api.Client
}

type session struct {
	URL        string
	RoleID     string
	RoleSecret string
	Client     *api.Client
}

type onFailedRenew func()
