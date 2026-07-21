package storage

import "os"

func CreateUserFolder(id string) error {
	return os.Mkdir("files/"+id, 0755)
}
