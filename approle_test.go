package cvault

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewApprole(t *testing.T) {
	inputID := "testID"
	inputSecret := "testSecret"

	outApprole := NewApprole(inputID, inputSecret)

	assert.Equal(t, inputID, outApprole.RoleID)
	assert.Equal(t, inputSecret, outApprole.RoleSecret)
}

func TestGetLoginEndpoint(t *testing.T) {
	inputAuth := &ApproleAuth{
		RoleID:     "testID",
		RoleSecret: "testSecret",
	}

	assert.Equal(t, "auth/approle/login", inputAuth.GetLoginEndpoint())
}

func TestGetLoginData(t *testing.T) {
	inputAuth := &ApproleAuth{
		RoleID:     "testID",
		RoleSecret: "testSecret",
	}

	expectedData := map[string]interface{}{
		"role_id":   "testID",
		"secret_id": "testSecret",
	}

	assert.Equal(t, expectedData, inputAuth.GetLoginData())
}
