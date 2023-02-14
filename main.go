package main

import (
	"crypto/ecdsa"
	"encoding/hex"
	"eth-collision/tool"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"
	"strings"
	"time"
)

var totalFile = "total.txt"
var matchFile = "match.txt"
var findFile = "find.txt"
var speedFile = "speed.txt"

// second
const rollupTime time.Duration = 1 * 60 * 60
const submitTime time.Duration = 1 * 60
const goroutineNum = 1024

func main() {
	LoadFromModelFile()
	msg := make(chan *big.Int)
	for i := 0; i < goroutineNum; i++ {
		go generateAccountJob(msg)
	}
	total, lastTotal := getInitTotal()
	ticker, callback := tool.NewProxyTicker(rollupTime * time.Second)
	go callback()
	for {
		select {
		case <-ticker:
			speed := new(big.Int).Sub(total, lastTotal)
			lastTotal = big.NewInt(0).Set(total)
			matchAddress, err := tool.FileCountLine(matchFile)
			if err != nil {
				log.Println(err)
			}
			findAddress, err := tool.FileCountLine(findFile)
			if err != nil {
				log.Println(err)
			}
			text := getNotifyText(total, speed, matchAddress, findAddress)
			log.Println(text)
			tool.AppendFile(speedFile, text)
			tool.SendMsgText(text)
		case count := <-msg:
			total = tool.BigIntAdd(total, count)
			tool.WriteFile(totalFile, total.String())
		}
	}
}

func getInitTotal() (*big.Int, *big.Int) {
	totalStr := tool.ReadFile(totalFile)
	totalStr = strings.TrimSpace(totalStr)
	total, ok := new(big.Int).SetString(totalStr, 10)
	if !ok {
		total = big.NewInt(-1)
	}
	lastTotal := total
	return total, lastTotal
}

func getNotifyText(total *big.Int, speed *big.Int, matchAddrs int, findAddrs int) string {
	dataStr := tool.FormatInt(int64(GetBloomLength()))
	totalStr := tool.FormatBigInt(*total)
	speedStr := tool.FormatBigInt(*speed)
	matchAddrsStr := tool.FormatInt(int64(matchAddrs))
	findAddrsStr := tool.FormatInt(int64(findAddrs))
	ipStr := tool.GetOutboundIP().String()
	text := fmt.Sprintf(""+
		"[ETH Collision Match Address]\n"+
		"Target: %s\n"+
		"Total: %s\n"+
		"Speed: %s\n"+
		"Matchs: %s\n"+
		"Finds: %s\n"+
		"IP: %s\n",
		dataStr, totalStr, speedStr, matchAddrsStr, findAddrsStr, ipStr)
	return text
}

func generateAccountJob(msg chan *big.Int) {
	var count int64 = 0
	tick := time.Tick(submitTime * time.Second)
	for {
		select {
		case <-tick:
			msg <- big.NewInt(count)
			count = 0
		default:
			generateAccount()
			count++
		}
	}
}

func generateAccount() {
	key, err := crypto.GenerateKey()
	if err != nil {
		log.Println(err)
	}
	address := crypto.PubkeyToAddress(key.PublicKey).Hex()
	checkAccount(key, address)
}

func checkAccount(key *ecdsa.PrivateKey, address string) {
	if checkAddressInBloom(address) {
		go func() {
			check := checkBalanceInEthScan(address)
			if check {
				handleFoundAddress(key, address, matchFile)
			}
		}()
	}
	if checkAddressInRules(address) {
		handleFoundAddress(key, address, findFile)
	}
}

func handleFoundAddress(key *ecdsa.PrivateKey, address string, filename string) {
	if key == nil {
		log.Println("key is nil")
		return
	}
	privateKey := hex.EncodeToString(key.D.Bytes())
	log.Println("Found: ", privateKey, address)
	text := fmt.Sprintf("%s,%s\n", privateKey, address)
	tool.AppendFile(filename, text)
	tool.SendMsgText(text)
}

func checkAddressInBloom(address string) bool {
	address = address[2:]
	ok := CheckDataInBloom(address)
	if ok {
		return true
	}
	return false
}

func checkBalanceInEthScan(address string) bool {
	balance, err := tool.GetBalanceFromEthScan(address)
	log.Println("Get balance from ethScan: ", balance, " err: ", err, " address: ", address, "")
	if err != nil {
		// notice err will return true
		return true
	}
	if balance.Cmp(big.NewInt(0)) > 0 {
		return true
	}
	return false
}

func checkAddressInRules(address string) bool {
	if strings.HasPrefix(address, "0x000000000") || strings.HasPrefix(address, "000000000") {
		return true
	}
	return false
}
