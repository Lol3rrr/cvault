package cvault

import "errors"

func (s *session) ReadMap(name string) (map[string]interface{}, error) {
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

	data, worked := resp.Data["data"].(map[string]interface{})
	if !worked {
		return nil, errors.New("Could not convert Data to map")
	}

	return data, nil
}
