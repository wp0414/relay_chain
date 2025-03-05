package config

import (
	"github.com/hokaccha/go-prettyjson"
	"github.com/spf13/viper"
	"log"
)

type ListenConfig struct {
	RPCAdd string `mapstructure:"rpc"`
}

type ChainConfig struct {
	ChainType                       string `mapstructure:"type"`
	SdkConfPath                     string `mapstructure:"config"`
	ChainId                         string `mapstructure:"chainId"`
	CrossTransactionContractAddress string `mapstructure:"address"`
}

type DbConfig struct {
	DbFile string `mapstructure:"dbFile"`
}

type Config struct {
	ListenConfig `mapstructure:"listen"`
	ChainConfig  `mapstructure:"chain"`
	DbConfig     `mapstructure:"db"`
	// Mode mean this application work mode, one of consensus and listen
	Mode string `mapstructure:"mode"`
}

func InitConfig(configFile string) (*Config, error) {
	cfViper := viper.New()
	//cfViper.configFile = configFile
	cfViper.SetConfigFile(configFile)
	//从配置文件中读取相关参数
	if err := cfViper.ReadInConfig(); err != nil {
		log.Fatal("read config failed, ", err)
	}
	config := &Config{}
	//将相关参数反序列化到config结构体中
	if err := cfViper.Unmarshal(&config); err != nil {
		log.Fatal("Unmarshal config failed, ", err)
	}
	config.printLog()
	return config, nil
}

func (c *Config) printLog() {
	json, err := prettyjson.Marshal(c)
	if err != nil {
		log.Fatalf("marshal alarm config failed, %s", err.Error())
	}
	log.Println(string(json))
}
