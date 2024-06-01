package main

import (
	"flag"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/zjyl1994/shortlinkd/infra/vars"
	"github.com/zjyl1994/shortlinkd/server"
	"github.com/zjyl1994/shortlinkd/service/code"
)

func main() {
	if err := errMain(); err != nil {
		logrus.Fatalln(err.Error())
	}
}

func errMain() error {
	flag.StringVar(&vars.LISTEN, "listen", "localhost:9900", "listen address")
	flag.StringVar(&vars.CONFIG_FILE, "config", "config.yaml", "config file")
	flag.BoolVar(&vars.DEBUG_MODE, "debug", false, "enter debug mode")
	flag.Parse()
	logrus.Infoln("Shortlinkd start at", vars.LISTEN, "use config", vars.CONFIG_FILE)

	if vars.DEBUG_MODE {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.Debugln("Shortlinkd run in debug mode.")
	}

	cfg, err := vars.LoadConfig(vars.CONFIG_FILE)
	if err != nil {
		return err
	}
	initCodes(cfg.List)

	return server.Run(vars.LISTEN)
}

func initCodes(items map[string]vars.ListItemS) error {
	data := make([]code.CodeItem, 0, len(items))
	for shortCode, item := range items {
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
