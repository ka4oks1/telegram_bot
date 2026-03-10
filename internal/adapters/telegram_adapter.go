package adapters

import (
	"config"
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"os"
)

type Adapter interface {
	Initialization(logSource string) error
}
type TgAdapter struct {
	saver config.FileSaver
}

func (tg *TgAdapter) Initialization(saverPath string, fileCap int64) error {

	file, err := os.OpenFile(saverPath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)

	//tg.saver = config.FileSaver{
	//	Path:      saverPath,
	//	MaxWeight: fileCap,
	//}

	bot, botErr := tgbotapi.NewBotAPI("asd")

	if err != nil {

		log.Println("any errors with creation/opening tg_log file")

		err = fmt.Errorf("problems with open tg log file:%w", err)

	}

	log.SetOutput(file)

	return err

}
