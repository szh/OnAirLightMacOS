package server

import (
	"io"
	"log"
	"net/http"
	"onair/server/light"
	"onair/server/util"
	"strconv"
)

func Start() {
	http.HandleFunc("/setstate", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		body, err := io.ReadAll(r.Body)
		if checkError(w, err) {
			return
		}

		value, err := strconv.ParseBool(string(body))
		if checkError(w, err) {
			return
		}

		log.Printf("Received request to set busy status to " + strconv.FormatBool(value))
		light.SetState(value)
	})

	log.Fatal(http.ListenAndServe(":"+util.Config.Port, nil))
}

func checkError(w http.ResponseWriter, err error) bool {
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "Can't read request body", http.StatusBadRequest)
		return true
	}
	return false
}
