package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("message.txt")
	if err != nil {
		log.Println(err)
		return
	}

	var a = make([]byte, 8)

	for {
		isEOF, err := file.Read(a)
		if err != nil {
			log.Println(err)
			return
		}

		if isEOF == 0 {
			log.Println("end of file")
		}

		fmt.Printf("read: %s\n", a)
	}

}
