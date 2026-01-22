package main

import (
	"fmt"
	"net/http"
	"os"
)

// main bot file

func main() {

	resp, err := http.Get("https://google.com")

	if err != nil {

		fmt.Println("Something went wrong")

	}

	defer resp.Body.Close()

	for true {

		bs := make([]byte, 1014)
		n, err := resp.Body.Read(bs)

		str := fmt.Sprintln(string(bs[:n]))

		fmt.Println(string(bs[:n]))

		file, fileErr := os.OpenFile("../HTTP_req.txt", os.O_RDWR, 0644)
		file.WriteString(str + "\n")

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
