package main

import (
	"adapters"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
)

// main bot file

const (
	startQuote          = "\"q\":\""
	endQuoteStartAuthor = "\",\"a\":\""
	endAuthor           = "\",\"h\":\""
)

const (
	logFileName = "generalLog.txt"
)

func findParsingIndexes(byteSet []byte, bytesRead int) (int, int, int) {

	sQindex := strings.Index(string(byteSet[:bytesRead]), startQuote)
	eQsAindex := strings.Index(string(byteSet[:bytesRead]), endQuoteStartAuthor)
	eAindex := strings.Index(string(byteSet[:bytesRead]), endAuthor)

	return sQindex, eQsAindex, eAindex
}

func getQuoteInfo(byteSet []byte, startQuoteInd int, endQuoteStartInd int, endAuthorInd int) (string, string) {

	quote := string(byteSet[startQuoteInd+len(startQuote) : endQuoteStartInd])
	author := string(byteSet[endQuoteStartInd+len(endQuoteStartAuthor) : endAuthorInd])

	return quote, author
}

func mainInitialization(logSourceName string) error {
	//	collect all errors
	var err error
	errTelegramAdap := adapters.TelegramAdapterInitialization(logSourceName)
	err = errors.Join(errTelegramAdap)
	return err
}

func main() {
	//add database adding info about authors

	initErrors := mainInitialization(logFileName)

	if initErrors != nil {
		panic("initialization errors")
	}

	resp, err := http.Get("https://zenquotes.io/api/random")

	if err != nil {
		fmt.Println("Error while server response waiting")
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	for {

		byteSet := make([]byte, 1024)
		bytesRead, err := resp.Body.Read(byteSet)

		str := fmt.Sprint(string(byteSet[:bytesRead]))

		sQindex, eQsAindex, eAindex := findParsingIndexes(byteSet, bytesRead)

		file, fileErr := os.OpenFile("HTTP_resp.txt", os.O_RDWR, 0644)

		if fileErr != nil {
			fmt.Println(fileErr)
			fmt.Println("Error when opening file")
		}

		if bytesRead == 0 || err != nil {
			break
		} else {

			quote, author := getQuoteInfo(byteSet, sQindex, eQsAindex, eAindex)

			fmt.Printf("\"%s\"\n", quote)
			fmt.Printf("Author is %s\n", author)

			_, writeErr := file.WriteString(str)
			if writeErr != nil {
				fmt.Println("Error while writing in file")
			}
		}

		errFileClosing := file.Close()

		if errFileClosing != nil {

			fmt.Println("error when close file:", errFileClosing)
		}

	}
}
