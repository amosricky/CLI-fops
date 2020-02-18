package setting

import (
	"github.com/go-ini/ini"
	"github.com/sirupsen/logrus"
)

type System struct {
	Version string
}

var SystemSetting = &System{}

var cfg *ini.File

// Setup initialize the configuration instance
func Setup() {
	var err error
	cfg, err = ini.Load("./conf/app.ini")
	if err != nil {
		logrus.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}

	MapTo("system", SystemSetting)
}

// Map to map section
func MapTo(section string, v interface{}) {
	var err error
	err = cfg.Section(section).MapTo(v)
	if err != nil {
		logrus.Fatalf("Cfg.MapTo %s err: %v", section, err)

	}
}