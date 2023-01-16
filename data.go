package main

import (
	"bufio"
	"github.com/bits-and-blooms/bloom/v3"
	"log"
	"os"
)

var fileList = []string{
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
}

var bloomFilter = bloom.NewWithEstimates(200000000, 0.000000001)

func init() {
	filename := "../eth-address-all/model.bin"
	file, err := os.Open(filename)
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

func ReadFromFile() {
	for _, filename := range fileList {
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

func CheckData(key string) bool {
	return bloomFilter.TestString(key)
}

func WriteTo() {
	filename := "../eth-address-all/model.bin"
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Println(err)
		return
	}
	to, err := bloomFilter.WriteTo(file)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(to)
}

func ReadFrom() {

}

func Rate() {
	rate := bloom.EstimateFalsePositiveRate(1000000000, 128, 200000000)
	log.Println(rate)
}
