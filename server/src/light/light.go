package light

import (
	"log"
	"onair/server/util"

	rpio "github.com/stianeikeland/go-rpio/v4"
)

type RGBStruct struct {
	PinRed   rpio.Pin
	PinGreen rpio.Pin
}

func (RGBStruct) AllPins() []rpio.Pin {
	return []rpio.Pin{RGBLight.PinRed, RGBLight.PinGreen}
}

var RGBLight RGBStruct

func SetupRGB() {
	if util.Config.TestMode {
		log.Println("Test mode: Skipping GPIO access")
		return
	}

	err := rpio.Open()
	if err != nil {
		log.Fatal("Cannot access GPIO:", err)
	}

	RGBLight = RGBStruct{
		PinRed:   rpio.Pin(util.Config.PinRed),
		PinGreen: rpio.Pin(util.Config.PinGreen),
	}
	// Set pins to output mode
	for _, pin := range RGBLight.AllPins() {
		pin.Output()
	}
}

func SetState(busy bool) {
	if util.Config.TestMode {
		log.Println("Test mode: Skipping GPIO access")
		return
	}

	if busy {
		RGBLight.PinRed.High()
		RGBLight.PinGreen.Low()
	} else {
		RGBLight.PinRed.Low()
		RGBLight.PinGreen.High()
	}
}

func Close() {
	rpio.Close()
}
