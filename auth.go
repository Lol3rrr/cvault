package cvault

import (
	"errors"
	"time"

	"github.com/hashicorp/vault/api"
	"github.com/sirupsen/logrus"
)

func (s *session) auth() error {
	vaultConfig := &api.Config{
		Address: s.URL,
	}
	vaultClient, err := api.NewClient(vaultConfig)
	if err != nil {
		return err
	}

	vaultData := map[string]interface{}{
		"role_id":   s.RoleID,
		"secret_id": s.RoleSecret,
	}
	resp, err := vaultClient.Logical().Write("auth/approle/login", vaultData)
	if err != nil {
		return err
	}
	if resp.Auth == nil {
		return errors.New("No Auth info returned")
	}

	vaultClient.SetToken(resp.Auth.ClientToken)
	s.Client = vaultClient

	if resp.Auth.Renewable {

		go func() {
			for {
				s, err := vaultClient.Auth().Token().RenewSelf(resp.LeaseDuration)
				if err != nil {
					logrus.Errorf("Could not renew Lease: %s \n", err)
					return
				}

				nextRenew := s.Auth.LeaseDuration / 2
				time.Sleep(time.Duration(nextRenew) * time.Second)
			}
		}()
	}

	return nil
}
