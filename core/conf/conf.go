package conf

import (
	"FishingSpotMarker/pkg/config"
	"sync"
)

var (
	// 钓点标记配置
	fishing_spot_marker_viper_config *config.ViperConfig

	// 钓点标记配置并发锁
	fishing_spot_marker_viper_config_lock sync.Mutex
)

// 初始化
func Init() error {
	err := config.InitViperConfig(func() {
		// 加载配置
		viper_config, err := config.LoadViperConfig()
		if err != nil {
			panic(err)
		}

		// 并发安全
		fishing_spot_marker_viper_config_lock.Lock()
		defer fishing_spot_marker_viper_config_lock.Unlock()

		fishing_spot_marker_viper_config = viper_config
	})

	// 加载配置
	viper_config, err := config.LoadViperConfig()
	if err != nil {
		return err
	}

	// 并发安全
	fishing_spot_marker_viper_config_lock.Lock()
	defer fishing_spot_marker_viper_config_lock.Unlock()

	fishing_spot_marker_viper_config = viper_config

	return err
}

// 获取配置
func GetConf() *config.ViperConfig {
	// 并发安全
	fishing_spot_marker_viper_config_lock.Lock()
	defer fishing_spot_marker_viper_config_lock.Unlock()

	return fishing_spot_marker_viper_config
}
