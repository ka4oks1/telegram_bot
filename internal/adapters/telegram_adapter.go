package adapters

import (
	"fmt"
	"log"
	"os"
	"telegram_bot/config"
)

type Adapter interface {
	Initialization(logSource string) error
}
type TgAdapter struct {
	saver config.FileSaver
}

func (tg *TgAdapter) Initialization(saverPath string, fileCap int64) error {

	file, err := os.OpenFile(saverPath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)

	tg.saver = config.FileSaver{
		Saver:     nil,
		Path:      saverPath,
		MaxWeight: fileCap,
		Weight:    0,
	}

	if err != nil {

		log.Println("any errors with creation/opening tg_log file")

		err = fmt.Errorf("problems with open tg log file:%w", err)

	}

	log.SetOutput(file)

	return err

}
