package main

import (
	"log"
	"os"
)

var data = ""

func InitData() {
	readFile, err := os.ReadFile("address.txt")
	if err != nil {
		log.Println(err)
		return
	}
	data = string(readFile)
}
