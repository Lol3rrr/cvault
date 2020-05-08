package cvault

import "github.com/hashicorp/vault/api"

func (s *session) Renew(id string, interval int) (*api.Secret, error) {
	return s.Client.Sys().Renew(id, interval)
}
