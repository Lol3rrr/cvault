package cvault

import (
	"errors"

	"github.com/hashicorp/vault/api"
)

func (s *session) ReadData(name string) (*api.Secret, error) {
	resp, err := s.Client.Logical().Read(name)
	if err != nil {
		return nil, err
	}

	if resp == nil {
		return nil, errors.New("No Response was returned")
	}

	if resp.Data == nil {
		return nil, errors.New("No Data was returned")
	}

	return resp, nil
}
