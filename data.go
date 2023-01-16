package main

import (
	"bufio"
	"log"
	"os"
)

var data = map[string]string{}

func init() {
	file, err := os.Open("address.txt")
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data[scanner.Text()] = ""
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
