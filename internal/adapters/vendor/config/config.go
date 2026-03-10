package config

import (
	"errors"
	"os"
	"time"
)

type Saver interface {
	Initialization(path string, max int64) error
	Save(info string) error
	Load()
}

type FileSaver struct {
	Saver
	Path        string
	MaxWeight   int64
	Weight      int64
	quotesCount uint64
	lastQuote   time.Time
}

func (fs *FileSaver) Initialization(path string, max int64) error {

	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)

	if err != nil {
		return err
	}
	if err := file.Close(); err != nil {

		return err
	}

	fs.Path = path

	fs.MaxWeight = max
	return nil
}

func (fs *FileSaver) Save(info string) error {

	if fs.Weight >= fs.MaxWeight {

		return errors.New("big count quotes")
	}

	err := os.WriteFile(fs.Path, []byte(info), os.FileMode(os.O_RDWR|os.O_APPEND|os.O_CREATE))
	if err != nil {

		return err
	}

	return nil
}
