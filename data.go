package main

import (
	"bufio"
	"github.com/bits-and-blooms/bloom/v3"
	"log"
	"os"
)

var bloomFilter = bloom.NewWithEstimates(200000000, 0.000000001)
var modelFile = "../eth-address-all/model.bin"

func init() {
	LoadFromModelFile()
}

func LoadFromSourceFile() {
	var sourceFileList = []string{
		"../eth-address-all/data0.txt",
		"../eth-address-all/data1.txt",
		"../eth-address-all/data2.txt",
		"../eth-address-all/data3.txt",
		"../eth-address-all/data4.txt",
		"../eth-address-all/data5.txt",
		"../eth-address-all/data6.txt",
		"../eth-address-all/data7.txt",
		"../eth-address-all/data8.txt",
		"../eth-address-all/data9.txt",
		"../eth-address-all/dataa.txt",
		"../eth-address-all/datab.txt",
		"../eth-address-all/datac.txt",
		"../eth-address-all/datad.txt",
		"../eth-address-all/datae.txt",
		"../eth-address-all/dataf.txt",
		"../eth-address-top-list/address.txt",
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
			bloomFilter.AddString(scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			log.Println(err)
		}
	}
}

func CheckDataInBloom(key string) bool {
	return bloomFilter.TestString(key)
}

func GenerateModelFIle() {
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
	defer file.Close()
	_, err = bloomFilter.ReadFrom(file)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("load model success", bloomFilter.Cap())
}

func GetBloomLength() uint {
	return bloomFilter.BitSet().Len()
}
