package main

import (
	"os"

	"github.com/mexisme/go-subprocess-mon/subprocess"

	"github.com/mexisme/go-config"
	//"github.com/hashicorp/go-multierror"
	log "github.com/sirupsen/logrus"
)

func main() {
	config.Init(config.Config{EnvPrefix: "subprocess", FromConfig: true})

	if len(os.Args) < 2 {
		log.Panic("No command provided to execute")
	}
	if err := subprocess.New().SetEnviron(os.Environ()).SetCommand(os.Args[1:]).Run(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
