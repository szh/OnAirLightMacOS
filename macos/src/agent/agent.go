package agent

import (
	"log"
	"onair/mac/light"
	"onair/mac/meeting"
	"time"
)

var isInMeetingPrev = false

func Start() {
	for {
		isInMeeting := meeting.InvokeGetMeetingState()
		if isInMeeting != isInMeetingPrev {
			log.Printf("In a meeting:%t", isInMeeting)
			isInMeetingPrev = isInMeeting
			if isInMeeting {
				light.TurnOn()
			} else {
				light.TurnOff()
			}
		}

		time.Sleep(1 * time.Second)
	}
}
