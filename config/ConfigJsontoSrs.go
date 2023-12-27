package config

type RouteRules struct {
	Domain       []string `json:"domain"`
	DomainSuffix []string `json:"domain_suffix"`
	IPCIDR       []string `json:"ip_cidr"`
}

type Info struct {
	Version int    `json:"version"`
	Rules   []Rule `json:"rules"`
}
