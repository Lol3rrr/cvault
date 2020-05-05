package cvault

import (
	"time"

	"github.com/hashicorp/vault/api"
	"github.com/sirupsen/logrus"
)

func (s *session) RenewDataForever(secret *api.Secret, interval int) {
	go func() {
		for {
			_, err := s.Client.Sys().Renew(secret.LeaseID, interval)
			if err != nil {
				logrus.Errorf("[Vault] Could not renew Secret \n")
			}

			renewTime := interval / 2
			time.Sleep(time.Duration(renewTime) * time.Second)
		}
	}()
}
