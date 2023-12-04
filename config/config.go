package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

// Config 结构定义了配置文件的结构
type Config struct {
	// 配置项
	Device   string        `yaml:"device"`
	Proxy    string        `yaml:"proxy"`
	IP       string        `yaml:"ip"`
	Mask     string        `yaml:"mask"`
	Gateway  string        `yaml:"gateway"`
	ServerIP string        `yaml:"server_ip"`
	Networks []NetworkInfo `yaml:"networks"`
}
type NetworkInfo struct {
	Name    string `yaml:"name"`
	Address string `yaml:"address"`
}

// LoadConfig 函数从指定的 YAML 文件中加载配置
func LoadConfig(filename string) (*Config, error) {
	// 读取文件内容
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// 解析 YAML 数据
	cfg := &Config{}
	err = yaml.Unmarshal(data, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
