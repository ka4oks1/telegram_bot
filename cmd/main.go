package main

import (
	"fmt"
	"net/http"
	"os"
)

// main bot file

func main() {

	resp, err := http.Get("https://zenquotes.io/api/random")

	if err != nil {

		fmt.Println("Something went wrong")

	}

	defer resp.Body.Close()

	//var finalstr string

	for true {

		bs := make([]byte, 1014)
		n, err := resp.Body.Read(bs)

		str := fmt.Sprint(string(bs[:n]))

		startQuote := []rune{'"', 'q', '"'}

		endQuote := []rune{'"', 'a', '"'}

		currStack := make([]rune, 3, 3)

		var quote string

		for _, v := range str {

			if currStack[0] == startQuote[0] && currStack[1] == startQuote[1] && currStack[2] == startQuote[2] {
				quote += string(v)
			}

			if currStack[0] == endQuote[0] && currStack[1] == endQuote[1] && currStack[2] == endQuote[2] {
				break
			}

		}

		fmt.Println(string(bs[:n]))

		file, fileErr := os.OpenFile("../HTTP_req.txt", os.O_RDWR, 0644)

		file.WriteString(str)

		if fileErr != nil {
			fmt.Println("fileErr")
		}

		defer file.Close()
		if n == 0 || err != nil {
			break
		}
	}
	//fmt.Println("Hey man i am bot")

	//client := &http.Client{}

}
