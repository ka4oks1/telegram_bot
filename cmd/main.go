package main

import (
	"fmt"
	"main/internal/adapters"
	"net/http"
	"os"
	"strings"
)

// main bot file

func main() {

	adapters.Hello()

	resp, err := http.Get("https://zenquotes.io/api/random")

	if err != nil {
		fmt.Println("Error while server response waiting")
	}

	defer resp.Body.Close()

	for {

		byteSet := make([]byte, 1014)
		bytesRead, err := resp.Body.Read(byteSet)

		str := fmt.Sprint(string(byteSet[:bytesRead]))

		startQuote := "\"q\":\""

		endQuoteStartAuthor := "\",\"a\":\""

		endAuthor := "\",\"h\":\""

		sQindex := strings.Index(string(byteSet[:bytesRead]), startQuote)

		eQsAindex := strings.Index(string(byteSet[:bytesRead]), endQuoteStartAuthor)

		eAindex := strings.Index(string(byteSet[:bytesRead]), endAuthor)

		//	fmt.Println(string(byteSet[:bytesRead]))

		file, fileErr := os.OpenFile("../HTTP_req.txt", os.O_RDWR, 0644)

		if fileErr != nil {
			fmt.Println("Error when opening file")
		}

		defer file.Close()

		if bytesRead == 0 || err != nil {
			break
		} else {

			quote := string(byteSet[sQindex+len(startQuote) : eQsAindex])
			author := string(byteSet[eQsAindex+len(endQuoteStartAuthor) : eAindex])

			fmt.Printf("\"%s\"\n", quote)
			fmt.Printf("Author is %s", author)

			_, writeErr := file.WriteString(str)
			if writeErr != nil {
				fmt.Println("Error while writing in file")
			}
		}
	}

}
