package vars

import (
	"os"

	"gopkg.in/yaml.v3"
)

var (
	LISTEN      string
	CONFIG_FILE string
	DEBUG_MODE  bool
)

func LoadConfig(configFile string) (*ConfigS, error) {
	b, err := os.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	var config ConfigS
	err = yaml.Unmarshal(b, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
