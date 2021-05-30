package utils

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

func InitViper() {
	viper.AddConfigPath("configs")
	viper.SetConfigName("configs")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("cannot read viper config %s", err)
	}

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}

func ViperGetString(path string) string {
	return viper.GetString(path)
}
