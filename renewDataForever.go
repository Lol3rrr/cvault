package cvault

import (
	"time"

	"github.com/hashicorp/vault/api"
	"github.com/sirupsen/logrus"
)

func (s *session) RenewDataForever(secret *api.Secret, interval int, callback func()) {
	go func() {
		for {
			_, err := s.Client.Sys().Renew(secret.LeaseID, interval)
			if err != nil {
				logrus.Errorf("[Vault] Could not renew Secret: %s \n", err)

				callback()
				return
			}

			renewTime := interval / 2
			time.Sleep(time.Duration(renewTime) * time.Second)
		}
	}()
}
