package main

import (
	"log"
	"onair/server/light"
	"onair/server/server"
	"onair/server/util"
	"os"
	"path/filepath"
)

func main() {
	ex, err := os.Executable()
	if err != nil {
		log.Fatal("Cannot get path to executable:", err)
	}
	exPath := filepath.Dir(ex)

	_, err = util.LoadConfig(exPath)
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	light.SetupRGB()
	light.SetState(false)
	defer light.Close()
	server.Start()
}
