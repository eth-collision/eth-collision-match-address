package main

import (
	"bufio"
	"github.com/bits-and-blooms/bloom/v3"
	"log"
	"os"
	"strings"
	"sync"
)

const n = 1000000000
const fp = 0.00001

var bloomFilter = bloom.NewWithEstimates(n, fp)
var modelFile = "../eth-address-all/model.bin"

func LoadFromSourceFile() {
	var sourceFileList = []string{
		"../eth-address-all/180M/data0.txt",
		"../eth-address-all/180M/data1.txt",
		"../eth-address-all/180M/data2.txt",
		"../eth-address-all/180M/data3.txt",
		"../eth-address-all/180M/data4.txt",
		"../eth-address-all/180M/data5.txt",
		"../eth-address-all/180M/data6.txt",
		"../eth-address-all/180M/data7.txt",
		"../eth-address-all/180M/data8.txt",
		"../eth-address-all/180M/data9.txt",
		"../eth-address-all/180M/dataa.txt",
		"../eth-address-all/180M/datab.txt",
		"../eth-address-all/180M/datac.txt",
		"../eth-address-all/180M/datad.txt",
		"../eth-address-all/180M/datae.txt",
		"../eth-address-all/180M/dataf.txt",
		"../eth-address-all/130M/data0.txt",
		"../eth-address-all/130M/data1.txt",
		"../eth-address-all/130M/data2.txt",
		"../eth-address-all/130M/data3.txt",
		"../eth-address-all/130M/data4.txt",
		"../eth-address-all/130M/data5.txt",
		"../eth-address-all/130M/data6.txt",
		"../eth-address-all/130M/data7.txt",
		"../eth-address-all/130M/data8.txt",
		"../eth-address-all/130M/data9.txt",
		"../eth-address-all/130M/dataa.txt",
		"../eth-address-all/130M/datab.txt",
		"../eth-address-all/130M/datac.txt",
		"../eth-address-all/130M/datad.txt",
		"../eth-address-all/130M/datae.txt",
		"../eth-address-all/130M/dataf.txt",
		"../eth-address-all/89M/data0.txt",
		"../eth-address-all/89M/data1.txt",
		"../eth-address-all/89M/data2.txt",
		"../eth-address-all/89M/data3.txt",
		"../eth-address-all/89M/data4.txt",
		"../eth-address-all/89M/data5.txt",
		"../eth-address-all/89M/data6.txt",
		"../eth-address-all/89M/data7.txt",
		"../eth-address-all/89M/data8.txt",
		"../eth-address-all/89M/data9.txt",
		"../eth-address-all/89M/dataa.txt",
		"../eth-address-all/89M/datab.txt",
		"../eth-address-all/89M/datac.txt",
		"../eth-address-all/89M/datad.txt",
		"../eth-address-all/89M/datae.txt",
		"../eth-address-all/89M/dataf.txt",
		"../eth-address-top-list/address.txt",
	}
	group := sync.WaitGroup{}
	for _, filename := range sourceFileList {
		go func(filename string) {
			group.Add(1)
			log.Println("load file:", filename)
			file, err := os.Open(filename)
			if err != nil {
				log.Println(err)
				return
			}
			defer file.Close()
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				text := scanner.Text()
				bloomFilter.AddString(text)
				bloomFilter.AddString(strings.ToUpper(text))
				bloomFilter.AddString(strings.ToLower(text))
			}
			if err := scanner.Err(); err != nil {
				log.Println(err)
			}
			group.Done()
		}(filename)
	}
	group.Wait()
}

func VerifyFromFile() {
	var sourceFileList = []string{
		"./no-random.txt",
	}
	for _, filename := range sourceFileList {
		log.Println("load file:", filename)
		file, err := os.Open(filename)
		if err != nil {
			log.Println(err)
			continue
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			address := scanner.Text()[2:]
			if CheckDataInBloom(address) {
				log.Println("verify success", address)
			}
		}
		if err := scanner.Err(); err != nil {
			log.Println(err)
		}
	}
}

func CheckDataInBloom(key string) bool {
	a := bloomFilter.TestString(key)
	b := bloomFilter.TestString(strings.ToUpper(key))
	c := bloomFilter.TestString(strings.ToLower(key))
	return a || b || c
}

func GenerateModelFile() {
	LoadFromSourceFile()
	file, err := os.OpenFile(modelFile, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Println(err)
		return
	}
	to, err := bloomFilter.WriteTo(file)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("generate model success", to)
}

func LoadFromModelFile() {
	file, err := os.Open(modelFile)
	if err != nil {
		log.Println(err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println(err)
		}
	}(file)
	_, err = bloomFilter.ReadFrom(file)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Load model success", GetBloomLength(), "/", bloomFilter.Cap())
}

func GetBloomLength() uint {
	return bloomFilter.BitSet().Count()
}

func RealPositiveRate() float64 {
	m, k := bloom.EstimateParameters(n, fp)
	rate := bloom.EstimateFalsePositiveRate(m, k, n)
	return rate
}
