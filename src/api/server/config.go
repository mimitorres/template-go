package server

import (
	"fmt"
	"os"
	"path"

	"github.com/olebedev/config"
	"github.com/sirupsen/logrus"
)

var conf *config.Config

func init() {
	confDir := os.Getenv("CONF_DIR")
	if confDir == "" {
		base := os.Getenv("GOPATH")
		confDir = path.Join(base, "src/github.com/xmtorres/template-go/conf")
		logrus.Info("CONF_DIR is not set")
	}

	logrus.Info(fmt.Sprintf("conf is being read from: %s", confDir))

	if _, err := os.Stat(confDir); os.IsNotExist(err) {
		panic(err.Error())
	}

	var err error
	conf, err = config.ParseYaml(path.Join(confDir, "default.yml"))

	if err != nil {
		panic(err.Error())
	}
}
