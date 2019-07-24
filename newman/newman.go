package newman

import (
	"log"
	"fmt"
	"os/exec"
)

var baseURL = "https://api.getpostman.com/"

func JUnitReport(collectionID, environmentID, apiKey string) (out string, err error) {
	args := []string{
		"run",
		fmt.Sprintf("%scollections/%s?apikey=%s", baseURL, collectionID, apiKey),
		"--environment",
		fmt.Sprintf("%senvironments/%s?apikey=%s", baseURL, environmentID, apiKey),
		"--reporters",
		"cli",
		"--suppress-exit-code",
	}

	raw, err := invokeNewman(args)
	if err != nil {
		return "", fmt.Errorf("%s\n\n%s", err, raw)
	}

	// expect that output is text
	out = raw

	return out, nil
}

func invokeNewman(args []string) (out string, err error) {
	log.Println("running:", "newman", args)
	raw, err := exec.Command("newman", args...).Output()
	// expect that output is text
	out = string(raw)
	return
}
