package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/zjyl1994/shortlinkd/infra/vars"
	"github.com/zjyl1994/shortlinkd/server"
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
	err = vars.ApplyConfig(cfg)
	if err != nil {
		return err
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		for sig := range sigs {
			logrus.Infoln("Singal", sig.String(), "received")
			switch sig {
			case syscall.SIGHUP:
				cfg, err := vars.LoadConfig(vars.CONFIG_FILE)
				if err != nil {
					logrus.Errorln(err.Error())
					return
				}
				err = vars.ApplyConfig(cfg)
				if err != nil {
					logrus.Errorln(err.Error())
					return
				}
			case syscall.SIGINT:
				err := server.App.ShutdownWithTimeout(5 * time.Second)
				if err != nil {
					logrus.Errorln(err.Error())
					return
				}
			case syscall.SIGTERM:
				err := server.App.Shutdown()
				if err != nil {
					logrus.Errorln(err.Error())
					return
				}
			}
		}
	}()

	return server.Run(vars.LISTEN)
}
