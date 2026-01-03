package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("message.txt")
	if err != nil {
		log.Println(err)
		return
	}

	var a = make([]byte, 8)
	var line string

	for {
		isEOF, err := file.Read(a)
		if err != nil {

			log.Println(err)
			return
		}

		if isEOF == 0 {
			return
		}

		if strings.Contains(string(a), "\n") {
			str := strings.Split(string(a), "\n")
			line = line + strings.Join(str[:1], "")
			log.Println("read:", line)

			line = ""
			line = line + strings.Join(str[1:], "")
			continue
		}

		line = line + string(a)
	}
}
