package cvault

import (
	"errors"
	"io/ioutil"
	"os"
)

func CreateKubernetesSessionEnv() (Session, error) {
	vaultURL, found := os.LookupEnv("VAULT_URL")
	if !found {
		vaultURL = "http://localhost:8200"
	}
	kubernetesRole, found := os.LookupEnv("VAULT_ROLE")
	if !found {
		return nil, errors.New("Env-Variable 'VAULT_ROLE' could not be found")
	}
	kubernetesJWT, err := ioutil.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/token")
	if err != nil {
		return nil, err
	}

	kubernetesAuth := NewKubernetes(kubernetesRole, string(kubernetesJWT))

	return CreateSession(vaultURL, kubernetesAuth)
}
