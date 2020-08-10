package main

import "github.com/spf13/viper"

// 加载配置
func loadConfig() (vi *viper.Viper, err error) {
	var v *viper.Viper
	v = viper.New()
	v.SetConfigName("config.yaml")
	v.AddConfigPath("./config/")
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}
	return v, nil
}
