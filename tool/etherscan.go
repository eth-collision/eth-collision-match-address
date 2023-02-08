package tool

import (
	"encoding/json"
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
)

var ethScancfg EthScanCfg

func init() {
	ethScancfg.Init()
}

type EthScanCfg struct {
	Token string `yaml:"token"`
}

func (c *EthScanCfg) Init() {
	yamlFile, err := ioutil.ReadFile("ethscan-config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
		return
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Printf("Unmarshal: %v", err)
		return
	}
}

func GetBalanceFromEthScan(address string) (*big.Int, error) {
	if ethScancfg.Token == "" {
		// if config empty, pass check
		return nil, errors.New("ethscan config empty")
	}
	url := fmt.Sprintf("https://api.etherscan.io/api?module=account&action=balance&address=%s&tag=latest&apikey=%s",
		address, ethScancfg.Token)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	if result["status"] != "1" {
		return nil, err
	}
	balance := result["result"].(string)
	balanceInt, _ := new(big.Int).SetString(balance, 10)
	return balanceInt, nil
}
