package main

import (
	"encoding/hex"
	"fmt"
	"github.com/eth-collision/eth-collision-tool"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"
	"sync"
	"time"
)

var totalFile = "total.txt"
var accountsFile = "accounts.txt"
var speedFile = "speed.txt"

var locker = sync.Mutex{}

// second
var rollupTime time.Duration = 1 * 60 * 60
var submitTime time.Duration = 1 * 60

var datas [20]map[string]string

func main() {
	msg := make(chan *big.Int)
	for i := 0; i < 20; i++ {
		m := CopyMap(data)
		datas[i] = m
		go generateAccountJob(msg, i)
	}
	totalStr := tool.ReadFile(totalFile)
	n := new(big.Int)
	total, ok := n.SetString(totalStr, 10)
	if !ok {
		total = big.NewInt(0)
	}
	lastTotal := total
	tick := time.Tick(rollupTime * time.Second)
	for {
		select {
		case <-tick:
			speed := new(big.Int).Sub(total, lastTotal)
			lastTotal = total
			addresses, err := tool.FileCountLine(accountsFile)
			if err != nil {
				log.Println(err)
			}
			dataStr := tool.FormatInt(int64(len(data)))
			totalStr := tool.FormatBigInt(*total)
			speedStr := tool.FormatBigInt(*speed)
			addressesStr := tool.FormatInt(int64(addresses))
			text := fmt.Sprintf(""+
				"[ETH Collision Match Address]\n"+
				"Data : %s\n"+
				"Total: %s\n"+
				"Speed: %s\n"+
				"Addrs: %s\n",
				dataStr, totalStr, speedStr, addressesStr)
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

func generateAccountJob(msg chan *big.Int, dataId int) {
	count := big.NewInt(0)
	tick := time.Tick(submitTime * time.Second)
	for {
		select {
		case <-tick:
			msg <- count
			count = big.NewInt(0)
		default:
			generateAccount(dataId)
			count = count.Add(count, big.NewInt(1))
		}
	}
}

func generateAccount(dataId int) {
	key, err := crypto.GenerateKey()
	if err != nil {
		log.Println(err)
	}
	privateKey := hex.EncodeToString(key.D.Bytes())
	address := crypto.PubkeyToAddress(key.PublicKey).Hex()
	handleAccount(privateKey, address, dataId)
}

func handleAccount(privateKey string, address string, dataId int) {
	if checkAddress(address, dataId) {
		log.Println("Found: ", privateKey, address)
		text := fmt.Sprintf("%s,%s\n", privateKey, address)
		tool.AppendFile(accountsFile, text)
	}
}

func checkAddress(address string, dataId int) bool {
	_, ok := datas[dataId][address]
	if ok {
		return true
	}
	return false
}

func CopyMap(m map[string]string) map[string]string {
	n := map[string]string{}
	for k, v := range m {
		n[k] = v
	}
	return n
}