package config

type Config struct {
	Log       LogConfig   `json:"log"`
	DNS       DNSConfig   `json:"dns"`
	Inbounds  []Inbound   `json:"inbounds"`
	Outbounds []Outbound  `json:"outbounds"`
	Route     RouteConfig `json:"route"`
}

type LogConfig struct {
	Level     string `json:"level"`
	Timestamp bool   `json:"timestamp"`
}

type DNSConfig struct {
	Servers       []DNServer `json:"servers"`
	Rules         []DNSRule  `json:"rules"`
	DisableCache  bool       `json:"disable_cache"`
	DisableExpire bool       `json:"disable_expire"`
}

type DNServer struct {
	Tag     string `json:"tag"`
	Address string `json:"address"`
	Detour  string `json:"detour,omitempty"`
}

type DNSRule struct {
	Domain  string `json:"domain"`
	Geosite string `json:"geosite"`
	Server  string `json:"server"`
}

type Inbound struct {
	Type              string   `json:"type"`
	Tag               string   `json:"tag"`
	Listen            string   `json:"listen"`
	ListenPort        int      `json:"listen_port"`
	Sniff             bool     `json:"sniff"`
	SetSystemProxy    bool     `json:"set_system_proxy"`
	InterfaceName     string   `json:"interface_name,omitempty"`
	Inet4Address      string   `json:"inet4_address,omitempty"`
	AutoRoute         bool     `json:"auto_route,omitempty"`
	Inet4RouteAddress []string `json:"inet4_route_address,omitempty"`
}

type Outbound struct {
	Type           string `json:"type"`
	Tag            string `json:"tag"`
	Server         string `json:"server,omitempty"`
	ServerPort     int    `json:"server_port,omitempty"`
	ConnectTimeout string `json:"connect_timeout,omitempty"`
	TCPFastOpen    bool   `json:"tcp_fast_open,omitempty"`
	UDPFragment    bool   `json:"udp_fragment,omitempty"`
}

type RouteConfig struct {
	Rules []RouteRule `json:"rules"`
}

type RouteRule struct {
	Geosite  string `json:"geosite"`
	GeoIP    string `json:"geoip,omitempty"`
	Outbound string `json:"outbound"`
}
