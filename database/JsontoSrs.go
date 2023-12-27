package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"net"
	"os"
)

type Rule struct {
	Domain       []string `json:"domain"`
	DomainSuffix []string `json:"domain_suffix"`
	IPCIDR       []string `json:"ip_cidr"`
}

type Config struct {
	Version int    `json:"version"`
	Rules   []Rule `json:"rules"`
}

func JsonToSrs() {
	db, err := sql.Open("sqlite3", "path_to_your_sqlite_database.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT rule_name, domain, domain_suffix, ip_address FROM rules")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	rules := make(map[string]Rule)

	for rows.Next() {
		var ruleName string
		var domain string
		var domainSuffix string
		var ipAddress string
		err = rows.Scan(&ruleName, &domain, &domainSuffix, &ipAddress)
		if err != nil {
			fmt.Println(err)
			return
		}

		_, ipNet, err := net.ParseCIDR(ipAddress)
		if err != nil {
			fmt.Println(err)
			return
		}

		rule, ok := rules[ruleName]
		if !ok {
			rule = Rule{}
		}

		rule.Domain = append(rule.Domain, domain)
		rule.DomainSuffix = append(rule.DomainSuffix, domainSuffix)
		rule.IPCIDR = append(rule.IPCIDR, ipNet.String())

		rules[ruleName] = rule
	}

	config := Config{
		Version: 1,
		Rules:   make([]Rule, 0, len(rules)),
	}

	for _, v := range rules {
		config.Rules = append(config.Rules, v)
	}

	configJSON, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}

	configFile, err := os.Create("config.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer configFile.Close()

	_, err = configFile.Write(configJSON)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Created config.json")
}
