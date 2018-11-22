package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"strings"
)

type Config struct {
	Name string
}

func Init(cfg string) error {
	// 创建配置对象
	c := Config{
		Name: cfg,
	}

	// 初始化配置文件
	if err := c.initConfig(); err != nil {
		return err
	}

	// 监控配置文件变化并热加载程序
	c.watchConfig()

	return nil
}

// 初始化配置
func (c *Config) initConfig() error {
	if c.Name != "" {
		// ./groupon -c config.yaml
		// 如果指定了配置文件名称，则解析指定的配置文件
		viper.SetConfigFile(c.Name)
	} else {
		// 如果没有指定则解析默认配置文件
		viper.AddConfigPath("conf")
		viper.SetConfigName("config")
	}
	viper.SetConfigType("yaml")   // 设置配置文件格式为YAML
	viper.AutomaticEnv()          // 读取匹配的环境变量
	viper.SetEnvPrefix("GROUPON") // 读取环境变量前缀为GROUPON的内容
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)            // 读取环境变量参数时的符号转换
	if err := viper.ReadInConfig(); err != nil { // viper解析配置文件
		return err
	}
	return nil
}

// 监控配置文件变化并热加载程序
func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("config file change: %s", e.Name)
	})
}
