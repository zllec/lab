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
			log.Println("end of file")
			continue
		}

		if strings.Contains(string(a), "\n") {
			str := strings.Split(string(a), "\n")
			lastPart := str[:len(str)-1]

			line = line + strings.Join(str[0:len(str)-1], "")
			log.Println(line)

			line = ""
			line = line + strings.Join(lastPart, "")
			continue
		}

		line = line + string(a)

		// line = line + string(a)
		//
		// if strings.Contains(string(a), "\n") {
		// 	fmt.Printf("read: %s", line)
		// 	line = ""
		// 	continue
		// }
	}

}
