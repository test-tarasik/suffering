package internal

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

func StartServer() error {
	router := mux.NewRouter()

	router.Path("/register").Methods("POST").HandlerFunc(UserHandler)
	router.Path("/upload").Methods("POST").HandlerFunc(UploadFiles)

	if err := http.ListenAndServe(":9091", router); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}

		return err
	}

	return nil
}
