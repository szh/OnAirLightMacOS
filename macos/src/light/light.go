package light

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"onair/mac/util"
	"strconv"
)

func TurnOn() {
	setState(true)
}

func TurnOff() {
	setState(false)
}

func setState(value bool) {
	url := fmt.Sprintf("http://%s:%s/setstate", util.Config.PiIp, util.Config.PiPort)

	content := bytes.NewBuffer([]byte(strconv.FormatBool(value)))
	req, err := http.NewRequest("PUT", url, content)
	if err != nil {
		log.Println("Error:", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error:", err)
	}
	defer resp.Body.Close()
}
