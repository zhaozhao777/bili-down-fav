package conf

import (
	"gopkg.in/ini.v1"
)

const (
	IniPath  = "assets/config.ini"
	UserPath = "assets/user.ini"
	QrPath   = "assets/qr.png"
)

var file, _ = ini.Load(IniPath)

func List(section string) map[string]string {
	return file.Section(section).KeysHash()
}

func Get(section string, key string) string {
	return file.Section(section).Key(key).String()
}

func Save(section string, key string, value string) error {
	file.Section(section).Key(key).SetValue(value)
	return file.SaveTo(IniPath)
}
