package config

import (
	"github.com/spf13/viper"

	"github.com/lexkong/log"
)

var cfgDatabase *viper.Viper
var cfgApplication *viper.Viper
var cfgJwt *viper.Viper

func init() {
	viper.SetConfigName("settings")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Debug(err)
	}

	cfgDatabase = viper.Sub("settings.database")
	if cfgDatabase == nil {
		log.Fatal("config not found settings.database")
	}
	DatabaseConfig = InitDatabase(cfgDatabase)

	cfgApplication = viper.Sub("settings.application")
	if cfgApplication == nil {
		log.Fatal("config not found settings.application")
	}
	ApplicationConfig = InitApplication(cfgApplication)

	cfgJwt = viper.Sub("settings.jwt")
	if cfgJwt == nil {
		log.Fatal("config not found settings.jwt")
	}
	JwtConfig = InitJwt(cfgJwt)
}

func SetApplicationIsInit() {
	SetConfig("./config", "settings.application.isInit", false)
}

func SetConfig(configPath string, key string, value interface{}) {
	viper.AddConfigPath(configPath)
	viper.Set(key, value)
	viper.WriteConfig()
}