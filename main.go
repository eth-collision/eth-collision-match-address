package main

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"github.com/eth-collision/eth-collision-tool"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"
	"strings"
	"sync"
	"time"
)

var totalFile = "total.txt"
var matchFile = "match.txt"
var findFile = "find.txt"
var speedFile = "speed.txt"

var locker = sync.Mutex{}

// second
const rollupTime time.Duration = 1 * 60 * 60
const submitTime time.Duration = 1 * 60
const goroutineNum = 128

func main() {
	msg := make(chan *big.Int)
	for i := 0; i < goroutineNum; i++ {
		go generateAccountJob(msg)
	}
	totalStr := tool.ReadFile(totalFile)
	totalStr = strings.TrimSpace(totalStr)
	total, ok := new(big.Int).SetString(totalStr, 10)
	if !ok {
		total = big.NewInt(-1)
	}
	lastTotal := total
	ticker, callback := tool.NewProxyTicker(rollupTime * time.Second)
	go callback()
	for {
		select {
		case <-ticker:
			speed := new(big.Int).Sub(total, lastTotal)
			lastTotal = total
			matchAddrs, err := tool.FileCountLine(matchFile)
			if err != nil {
				log.Println(err)
			}
			findAddrs, err := tool.FileCountLine(findFile)
			if err != nil {
				log.Println(err)
			}
			text := getNotifyText(total, speed, matchAddrs, findAddrs)
			log.Println(text)
			tool.AppendFile(speedFile, text)
			tool.SendMsgText(text)
		case count := <-msg:
			total = bigIntAddMutex(total, count)
			tool.WriteFile(totalFile, total.String())
		}
	}
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

func bigIntAddMutex(a, b *big.Int) *big.Int {
	locker.Lock()
	defer locker.Unlock()
	c := new(big.Int)
	c.Add(a, b)
	return c
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
		handleFoundAddress(key, address, matchFile)
	}
	if checkAddressInRules(address) {
		handleFoundAddress(key, address, findFile)
	}
}

func handleFoundAddress(key *ecdsa.PrivateKey, address string, filename string) {
	privateKey := hex.EncodeToString(key.D.Bytes())
	// print to output
	log.Println("Found: ", privateKey, address)
	// write to file
	text := fmt.Sprintf("%s,%s\n", privateKey, address)
	tool.AppendFile(filename, text)
	// send to telegram
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

func checkAddressInRules(address string) bool {
	if strings.HasPrefix(address, "0x000000000") || strings.HasPrefix(address, "000000000") {
		return true
	}
	return false
}
