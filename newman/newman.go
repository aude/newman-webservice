package newman

import (
	"log"
	"fmt"
	"os"
	"os/exec"
	"io/ioutil"
)

var baseURL = "https://api.getpostman.com/"

func CLIReport(collectionID, environmentID, apiKey string) (out string, err error) {
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

func JUnitReport(collectionID, environmentID, apiKey string) (out string, err error) {
	// create tempfile
	file, err := ioutil.TempFile(os.TempDir(), "newman-webservice.junit-report.*.xml")
	if err != nil {
		return "", fmt.Errorf("error creating tempfile: %s", err)
	}
	// close tempfile, Newman should write to it now
	if err := file.Close(); err != nil {
		return "", fmt.Errorf("error closing tempfile: %s", err)
	}
	defer os.Remove(file.Name())

	args := []string{
		"run",
		fmt.Sprintf("%scollections/%s?apikey=%s", baseURL, collectionID, apiKey),
		"--environment",
		fmt.Sprintf("%senvironments/%s?apikey=%s", baseURL, environmentID, apiKey),
		"--reporters",
		"junit",
		"--reporter-junit-export",
		file.Name(),
		"--suppress-exit-code",
	}

	raw, err := invokeNewman(args)
	if err != nil {
		return "", fmt.Errorf("%s\n\n%s", err, raw)
	}

	// read tempfile
	content, err := ioutil.ReadFile(file.Name())
	if err != nil {
		return "", fmt.Errorf("error reading Newman output file (%s): %s", file.Name(), err)
	}

	// expect that output is text
	out = string(content)

	return out, nil
}

func invokeNewman(args []string) (out string, err error) {
	log.Println("running:", "newman", args)
	raw, err := exec.Command("newman", args...).Output()
	// expect that output is text
	out = string(raw)
	return
}
