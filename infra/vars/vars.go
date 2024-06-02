package vars

import (
	"os"
	"time"

	"github.com/zjyl1994/shortlinkd/service/code"
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

func ApplyConfig(cfg *ConfigS) error {
	data := make([]code.CodeItem, 0, len(cfg.List))
	for shortCode, item := range cfg.List {
		newItem := code.CodeItem{
			Code:    shortCode,
			URL:     item.URL,
			Enabled: true,
		}

		if item.Disabled {
			continue
		}

		if item.Expired > "" {
			t, err := time.Parse(time.DateTime, item.Expired)
			if err != nil {
				return err
			}
			newItem.Expired = &t
		}
		data = append(data, newItem)
	}
	code.InitCode(data)
	return nil
}
