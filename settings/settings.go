package settings

import (
	"fmt"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// 全局配置信息
var Conf = new(Config)

type Config struct {
	App       AppConfig       `mapstructure:"app"`
	Log       LogConfig       `mapstructure:"log"`
	MySQL     MySQLConfig     `mapstructure:"mysql"`
	Redis     RedisConfig     `mapstructure:"redis"`
	Email     EmailConfig     `mapstructure:"email"`
	Jwt       JwtConfig       `mapstructure:"jwt"`
	Snowflake SnowflakeConfig `mapstructrue:"snowflake"`
	Proxy     ProxyConfig     `mapstructrue:"proxy"`
	Cos       CosConfig       `mapstructrue:"cos"`
}

type AppConfig struct {
	Name string `mapstructure:"name"`
	Mode string `mapstructure:"mode"`
	Port int    `mapstructure:"port"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxBackups int    `mapstructure:"max_backups"`
	MaxAge     int    `mapstructure:"max_age"`
	Compress   bool   `mapstructure:"compress"`
}

type MySQLConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type EmailConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Expires  int    `mapstructure:"expires"`
}

type JwtConfig struct {
	Timeout int    `mapstructure:"timeout"`
	Issuer  string `mapstructure:"issuer"`
	Secret  string `mapstructure:"secret"`
}

type SnowflakeConfig struct {
	StartTime string `mapstructure:"start_time"`
	MachineID int64  `mapstructure:"machine_id"`
}

type ProxyConfig struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type CosConfig struct {
	TmpSecret TmpSecretConfig `mapstructure:"tmpsecret"`
	Resource  ResourceConfig  `mapstructure:"resource"`
	Action    []string        `mapstructure:"action"`
}

type TmpSecretConfig struct {
	Id  string `mapstructure:"id"`
	Key string `mapstructure:"key"`
}

type ResourceConfig struct {
	AllowPath string   `mapstructure:"allowpath"`
	Region    string   `mapstructure:"region"`
	Appid     string   `mapstructure:"appid"`
	Bucket    string   `mapstructure:"bucket"`
	AllowKey  []string `mapstructure:"allow_key"`
}

func Init() (err error) {
	//设置配置文件路径
	viper.SetConfigFile("./conf/config.yaml")

	//读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("viper load config failed,err:" + err.Error())
		return err
	}

	//将配置信息映射到结构体变量
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Println("viper umarshal Conf falied,err:" + err.Error())
		return err
	}

	fmt.Printf("配置信息为%v\n", Conf)

	viper.WatchConfig()

	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config has changed, time:" + time.Now().String())
		if err := viper.Unmarshal(Conf); err != nil {
			zap.L().Error("viper umarshal Conf falied,Error:", zap.Error(err))
			return
		}
		zap.L().Debug("config has changed", zap.Any("new config:", Conf))
	})

	return
}
