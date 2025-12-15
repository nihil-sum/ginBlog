package conf

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	App    Appconfig
	Logger Logger.Config
	ORM    orm.Config
	JWT    jwt.Config
}

type Appconfig struct {
	Name          string
	Version       string
	Mode          string
	Addr          string
	Host          string
	Resource      string
	Ffprobepath   string
	Env           string
	MpHost        string
	AdOctopusHost string
}

var (
	Conf = &Config{}
)

func Init(ConfigPath string) *Config {
	viper.SetConfigType("yml")
	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(&Conf); err != nil {
		panic(err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		if err := viper.Unmarshal(&Conf); err != nil {
			panic(err)
		}
	})
	return Conf
}
