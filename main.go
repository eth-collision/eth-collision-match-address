package main

import (
	"encoding/hex"
	"fmt"
	"github.com/eth-collision/eth-collision-tool"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"
	"regexp"
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
var rollupTime time.Duration = 1 * 60 * 60
var submitTime time.Duration = 1 * 60

func main() {
	msg := make(chan *big.Int)
	for i := 0; i < 100; i++ {
		go generateAccountJob(msg)
	}
	totalStr := tool.ReadFile(totalFile)
	totalStr = strings.TrimSpace(totalStr)
	n := new(big.Int)
	total, ok := n.SetString(totalStr, 10)
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
			log.Println(text)
			tool.AppendFile(speedFile, text)
			tool.SendMsgText(text)
		case count := <-msg:
			total = bigIntAddMutex(total, count)
			tool.WriteFile(totalFile, total.String())
		}
	}
}

func bigIntAddMutex(a, b *big.Int) *big.Int {
	locker.Lock()
	defer locker.Unlock()
	c := new(big.Int)
	c.Add(a, b)
	return c
}

func generateAccountJob(msg chan *big.Int) {
	count := big.NewInt(0)
	tick := time.Tick(submitTime * time.Second)
	for {
		select {
		case <-tick:
			msg <- count
			count = big.NewInt(0)
		default:
			generateAccount()
			count = count.Add(count, big.NewInt(1))
		}
	}
}

func generateAccount() {
	key, err := crypto.GenerateKey()
	if err != nil {
		log.Println(err)
	}
	privateKey := hex.EncodeToString(key.D.Bytes())
	address := crypto.PubkeyToAddress(key.PublicKey).Hex()
	handleAccount(privateKey, address)
}

func handleAccount(privateKey string, address string) {
	if checkAddressInBloom(address) {
		log.Println("Found: ", privateKey, address)
		text := fmt.Sprintf("%s,%s\n", privateKey, address)
		tool.AppendFile(matchFile, text)
		tool.SendMsgText(text)
	}
	if checkAddressInRules(address) {
		log.Println("Found: ", privateKey, address)
		text := fmt.Sprintf("%s,%s\n", privateKey, address)
		tool.AppendFile(findFile, text)
		tool.SendMsgText(text)
	}
}

func checkAddressInBloom(address string) bool {
	address = address[2:]
	ok := CheckDataInBloom(address)
	if ok {
		return true
	}
	return false
}

var re = regexp.MustCompile(`0x000000000`)

func checkAddressInRules(address string) bool {
	if re.MatchString(address) {
		return true
	}
	return false
}
