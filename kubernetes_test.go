package cvault

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewKubernetes(t *testing.T) {
	inputRole := "testRole"
	inputJWT := "testJWT"

	outAuth := NewKubernetes(inputRole, inputJWT)

	assert.Equal(t, inputRole, outAuth.Role)
	assert.Equal(t, inputJWT, outAuth.JWT)
}

func TestGetLoginData_Kubernetes(t *testing.T) {
	outAuth := NewKubernetes("testRole", "testJWT")

	expectedData := map[string]interface{}{
		"role": "testRole",
		"jwt":  "testJWT",
	}

	assert.Equal(t, expectedData, outAuth.GetLoginData())
}

func TestGetLoginEndpoint_Kubernetes(t *testing.T) {
	outAuth := NewKubernetes("testRole", "testJWT")

	assert.Equal(t, "auth/kubernetes/login", outAuth.GetLoginEndpoint())
}
