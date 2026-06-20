package storage

import (
	"fmt"
	"os"
)

// AppendMessage добавляет сообщение в лог топика (append-only)
func AppendMessage(topic string, msg string) error {
	// создаём папку storage если её нет
	err := os.MkdirAll("storage", 0755)
	if err != nil {
		return err
	}

	filePath := fmt.Sprintf("storage/%s.log", topic)

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(msg + "\n")
	return err
}
