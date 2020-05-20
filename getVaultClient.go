package cvault

import "github.com/hashicorp/vault/api"

func (s *session) GetVaultClient() *api.Client {
	return s.Client
}
