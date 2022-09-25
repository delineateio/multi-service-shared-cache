package main

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func loadConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../../")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}

func watchConfig() {
	viper.OnConfigChange(func(e fsnotify.Event) {
		load()
	})
	viper.WatchConfig()
}

func getString(key, def string) string {
	value := viper.GetString(key)
	if value == "" {
		return def
	} else {
		return value
	}
}
