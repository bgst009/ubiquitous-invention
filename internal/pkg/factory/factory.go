package factory

import (
	"io/ioutil"

	"github.com/bgst009/ubiquitous-invention/internal/pkg/config"
	"gopkg.in/yaml.v3"
)

var MonitorCfg *config.Config

func InitConfigFactory(path string) error {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	MonitorCfg = config.NewConfig()
	err2 := yaml.Unmarshal([]byte(b), MonitorCfg)
	if err2 != nil {
		return err2
	}

	return nil
}
