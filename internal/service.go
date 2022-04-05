package internal

import (
	"encoding/json"
	"errors"
	"net/http"
)

type Service struct {
	host string
}

func NewService() *Service {
	srvHost := GetEnvDefault("SERVER_HOST", "http://localhost:9000")
	return &Service{srvHost}
}

func (s *Service) Get() (string, error) {
	if res, err := http.Get(s.host); err == nil {
		var m map[string]interface{}
		json.NewDecoder(res.Body).Decode(&m)
		return m["message"].(string), nil
	}
	return "", errors.New("internal server error")
}
