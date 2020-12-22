package cvault

type KubernetesAuth struct {
	Role string
	JWT  string
}

func NewKubernetes(role, jwt string) *KubernetesAuth {
	return &KubernetesAuth{
		Role: role,
		JWT:  jwt,
	}
}

func (k *KubernetesAuth) GetLoginData() map[string]interface{} {
	return map[string]interface{}{
		"role": k.Role,
		"jwt":  k.JWT,
	}
}

func (k *KubernetesAuth) GetLoginEndpoint() string {
	return "auth/kubernetes/login"
}
