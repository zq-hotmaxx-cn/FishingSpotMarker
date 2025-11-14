package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func InitViperConfig(on_config_change_callback func()) error {
	// 设置配置文件名（不含扩展名）
	viper.SetConfigName("fishing_spot_marker.conf")
	// 设置配置文件类型
	viper.SetConfigType("yaml")
	// 添加配置文件搜索路径
	viper.AddConfigPath(".")
	viper.AddConfigPath("/etc/FishingSpotMarker/")

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	// 监听配置文件变化
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		on_config_change_callback()
	})

	return nil
}

func LoadViperConfig() (*ViperConfig, error) {
	viper_config := &ViperConfig{}
	err := viper.Unmarshal(viper_config)
	return viper_config, err
}
