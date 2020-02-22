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
	pathList := []string{"./conf/app.ini", "../conf/app.ini"}

	for _, path := range pathList{
		cfg, err = ini.Load(path)
		if err == nil {
			MapTo("system", SystemSetting)
			break
		}
	}
	if err != nil{
		logrus.Fatalf("setting.Setup, fail to parse 'app.ini': %v", err)
	}
}

// Map to map section
func MapTo(section string, v interface{}) {
	var err error
	err = cfg.Section(section).MapTo(v)
	if err != nil {
		logrus.Fatalf("Cfg.MapTo %s err: %v", section, err)

	}
}