package service

import "github.com/casbin/casbin/config"

type systemService struct{}

var (
	SystemService = new(systemService)
)

func (s *systemService) GetSystemConfig() *config.Config {
	return &config.Config{}
}
