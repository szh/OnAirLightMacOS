package main

import (
	"log"
	"onair/mac/agent"
	"onair/mac/meeting"
	"onair/mac/util"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

func main() {
	handleGetMeetingState()

	ex, err := os.Executable()
	if err != nil {
		log.Fatal("Cannot get path to executable:", err)
	}
	exPath := filepath.Dir(ex)

	_, err = util.LoadConfig(exPath)
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	startAgent()
}

func handleGetMeetingState() {
	for _, arg := range os.Args[1:] {
		if arg == "--get-meeting-state" {
			if meeting.GettMeetingState() {
				os.Exit(0)
			} else {
				os.Exit(1)
			}
		}
	}
}

func startAgent() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		os.Exit(0)
	}()

	agent.Start()
}
