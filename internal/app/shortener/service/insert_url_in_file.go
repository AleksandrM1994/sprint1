package service

import (
	"encoding/json"
	"os"
)

type URLInfo struct {
	UUID        int64  `json:"uuid"`
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

func (s *ServiceImpl) InsertURLInFile(URLInfo *URLInfo) error {
	data, err := json.Marshal(URLInfo)
	if err != nil {
		return err
	}
	// добавляем перенос строки
	data = append(data, '\n')

	// открываем файл для записи в конец
	file, err := os.OpenFile(s.cfg.FileStoragePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	_, errWrite := file.Write(data)
	if errWrite != nil {
		return errWrite
	}

	file.Close()

	return nil
}
