package core

import "github.com/Sirupsen/logrus"

type coreConfig struct {
	log *logrus.Logger
}

var cc coreConfig

func init() {
	cc = coreConfig{}
	cc.log = logrus.New()
	cc.log.Level = logrus.ErrorLevel
}

func InitLogger(l string) {
	var err error
	cc.log.Level, err = logrus.ParseLevel(l)
	if err != nil {
		cc.log.Level = logrus.ErrorLevel
	}
}
