package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	Log struct {
		Level     string `json:"level"`
		Timestamp bool   `json:"timestamp"`
	} `json:"log"`
	DNS struct {
		Servers       []Server `json:"servers"`
		Rules         []Rule   `json:"rules"`
		DisableCache  bool     `json:"disable_cache"`
		DisableExpire bool     `json:"disable_expire"`
	} `json:"dns"`
	Inbounds  []Inbound  `json:"inbounds"`
	Outbounds []Outbound `json:"outbounds"`
	Route     struct {
		Rules []RouteRule `json:"rules"`
	} `json:"route"`
}

type Server struct {
	Tag     string `json:"tag"`
	Address string `json:"address"`
	Detour  string `json:"detour,omitempty"`
}

type Rule struct {
	Domain  string `json:"domain"`
	Geosite string `json:"geosite"`
	Server  string `json:"server"`
}

type Inbound struct {
	Type string `json:"type"`
	Tag  string `json:"tag"`
	//Listen            string   `json:"listen"`
	//ListenPort        int      `json:"listen_port"`
	//Sniff             bool     `json:"sniff"`
	//SetSystemProxy    bool     `json:"set_system_proxy"`
	InterfaceName     string   `json:"interface_name,omitempty"`
	Inet4Address      string   `json:"inet4_address,omitempty"`
	AutoRoute         bool     `json:"auto_route,omitempty"`
	Inet4RouteAddress []string `json:"inet4_route_address,omitempty"`
}

type Outbound struct {
	Type       string `json:"type"`
	Tag        string `json:"tag"`
	Server     string `json:"server,omitempty"`
	ServerPort int    `json:"server_port,omitempty"`
	Version    string `json:"version,omitempty"`
	//ConnectTimeout string `json:"connect_timeout,omitempty"`
	//TCPFastOpen    bool   `json:"tcp_fast_open,omitempty"`
	//UDPFragment    bool   `json:"udp_fragment,omitempty"`
	Network    string `json:"network,omitempty"`
	UdpOverTCP bool   `json:"udp_over_tcp,omitempty"`
}

type RouteRule struct {
	Geosite  string `json:"geosite"`
	GeoIP    string `json:"geoip,omitempty"`
	Outbound string `json:"outbound"`
}

// LoadConfig 从指定的文件中加载配置
func LoadConfig(filename string) (*Config, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("读取配置文件失败: %v", err)
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("解析配置文件失败: %v", err)
	}

	return &config, nil
}

// SaveConfig 将配置保存到指定的文件
func SaveConfig(config *Config, filename string) error {
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("转换配置文件失败: %v", err)
	}

	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("写入新配置文件失败: %v", err)
	}

	return nil
}
