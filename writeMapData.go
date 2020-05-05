package cvault

func (s *session) WriteMapData(name string, data map[string]interface{}) error {
	resData := map[string]interface{}{
		"data": data,
	}
	_, err := s.Client.Logical().Write(name, resData)

	return err
}
