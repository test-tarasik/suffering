package storage

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

func CreateUserFolder(id string) error {
	return os.Mkdir("files/"+id, 0755)
}

func SaveFile(userID string, filename string, src multipart.File) error {
	path := "files/" + userID + "/" + filename

	ext := filepath.Ext(filename)
	name := strings.TrimSuffix(filename, ext)

	_, err := os.Stat(path)

	if err == nil {
		for i := 1; ; i++ {
			candidate := fmt.Sprintf("%s(%d)%s", name, i, ext)
			candidatePath := "files/" + userID + "/" + candidate

			_, err := os.Stat(candidatePath)
			if os.IsNotExist(err) {
				path = candidatePath
				break
			}

			if err != nil {
				return err
			}

		}
	}

	dst, err := os.Create(path)
	if err != nil {
		return err
	}

	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		os.Remove(path)
		return err
	}

	return nil
}
