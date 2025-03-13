package helpers

import (
	"fmt"
	"io"
	"os"
)

func ReadFile(path string) (string, error) {
	file, errOpen := os.Open(path)
	if errOpen != nil {
		return "", fmt.Errorf("error opening file: %w", errOpen)
	}
	defer file.Close()

	var data []byte
	buf := make([]byte, 64)

	for {
		n, errRead := file.Read(buf)
		if errRead != nil {
			if errRead == io.EOF {
				break
			}
			return "", fmt.Errorf("error reading file: %w", errRead)
		}
		data = append(data, buf[:n]...)
	}

	return string(data), nil
}
