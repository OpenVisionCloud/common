package config

import (
	"log"

	"github.com/spf13/viper"
)

// LoadConfig 加载配置文件
func LoadConfig(configPath string, cfg Config) Config {
	// 如果没有在命令行参数中指定配置文件的具体路径，则从默认路径查找
	if len(configPath) == 0 {
		// 支持从以下目录查找配置文件： ./config/config.yaml,  /etc/sso-proxy/config.yaml
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./config")
		viper.AddConfigPath("/etc/sso-proxy/")
	} else {
		// 从命令行指定的目录加载配置文件
		viper.SetConfigFile(configPath)
	}

	// 加载配置文件
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Failed to load config.yaml: %v", err)
		return err
	}

	// 序列化成struct
	if err := viper.Unmarshal(cfg); err != nil {
		log.Fatalf("Failed to unmarshal config file: %v", err)
		return err
	}

	// 监控配置文件变化
	viper.WatchConfig()

	//viper.OnConfigChange(func(in fsnotify.Event) {
	//	fmt.Println("Config file is changed, reload it now")
	//	var newConfig = &Config{}
	//
	//	// 更新现有的配置
	//	if err := viper.Unmarshal(newConfig); err != nil {
	//		panic(fmt.Errorf("The config cann't be updated:%s \n", err))
	//	}
	//	config = newConfig
	//
	//	// todo：日志组件更新
	//})
	return cfg
}
