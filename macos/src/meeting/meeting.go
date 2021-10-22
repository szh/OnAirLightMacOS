package meeting

import (
	"log"
	"os"
	"os/exec"

	mediaDevices "github.com/antonfisher/go-media-devices-state"
)

func GettMeetingState() bool {
	isCameraOn, err := mediaDevices.IsCameraOn()
	if err != nil {
		log.Println("Error:", err)
	}

	isMicOn, err := mediaDevices.IsMicrophoneOn()
	if err != nil {
		log.Println("Error:", err)
	}

	return isCameraOn || isMicOn
}

// mediaDevices module seems to cache the value, rendering it unusable to find state changes within a long-running process
// so we execute a second process here to discover state changes
func InvokeGetMeetingState() bool {
	selfExecutable := os.Args[0]
	cmd := exec.Command(selfExecutable, "--get-meeting-state")
	err := cmd.Run()

	return err == nil
}
