package tool

import (
	"bytes"
	"io"
	"log"
	"os"
)

func WriteFile(filename string, content string) {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	file.WriteString(content)
	file.Sync()
}

func AppendFile(filename string, text string) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	_, err = f.WriteString(text)
	if err != nil {
		log.Println(err)
	}
	f.Sync()
}

func FileCountLine(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}
	for {
		c, err := file.Read(buf)
		count += bytes.Count(buf[:c], lineSep)
		switch {
		case err == io.EOF:
			return count, nil
		case err != nil:
			return 0, err
		}
	}
}

func ReadFile(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	buf := make([]byte, 1024)
	var content string
	for {
		n, err := file.Read(buf)
		if n == 0 {
			break
		}
		if err != nil && err != io.EOF {
			log.Println(err)
		}
		content += string(buf[:n])
	}
	return content
}
