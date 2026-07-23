package storage

import (
	"io"
	"mime/multipart"
	"os"
)

func CreateUserFolder(id string) error {
	return os.Mkdir("files/"+id, 0755)
}

func SaveFile(userID string, filename string, src multipart.File) error {
	dst, err := os.Create("files/" + userID + "/" + filename)
	if err != nil {
		dst.Close()
		return err
	}

	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		os.Remove("files/" + userID + "/" + filename)
		return err
	}

	return nil
}
