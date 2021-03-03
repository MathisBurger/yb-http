package Var

import (
	"github.com/MathisBurger/yb-http/models"
)

var configurations []*models.HttpConfig

// appends config to configurations
func AppendConfig(cfg *models.HttpConfig) {
	configurations = append(configurations, cfg)
}

// gets config by domainname
func GetConfig(domainname string) *models.HttpConfig {
	for _, cfg := range configurations {
		if cfg.Server.Name == domainname {
			return cfg
		}
	}
	return nil
}
