package adapters

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"os"
)

func TelegramAdapterInitialization(logSource string) error {

	file, err := os.OpenFile(logSource, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)

	if err != nil {

		log.Println("any errors with creation/opening tg_log file")

		err = fmt.Errorf("problems with open tg log file:%w", err)

	}

	log.SetOutput(file)

	return err

}

func TelegramAdapter(logSource string) {

}
