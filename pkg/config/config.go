package config

type ViperConfig struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	// Cache    CacheConfig    `mapstructure:"cache"`
}

const (
	SERVER_MODE_DEV  = "dev"
	SERVER_MODE_PROD = "prod"
)

type ServerConfig struct {
	Mode   string       `mapstructure:"mode"`
	Port   uint         `mapstructure:"port"`
	Log    string       `mapstructure:"log"`
	Jwt    string       `mapstructure:"jwt"`
	Wechat WechatConfig `mapstructure:"wechat"`
}

type WechatConfig struct {
	AppID     string `mapstructure:"app_id"`
	AppSecret string `mapstructure:"app_secret"`
}

const (
	DATABASE_DRIVER_SQLITE   = "sqlite"
	DATABASE_DRIVER_MYSQL    = "mysql"
	DATABASE_DRIVER_POSTGRES = "postgres"
)

type DatabaseConfig struct {
	Driver string `mapstructure:"driver"`
	DSN    string `mapstructure:"dsn"`
}

const (
	CACHE_METHOD_MEMORY = "memory"
	CACHE_METHOD_REDIS  = "redis"
)

type CacheConfig struct {
	Method string `mapstructure:"method"`
	CDSN   string `mapstructure:"cdsn"`
}
