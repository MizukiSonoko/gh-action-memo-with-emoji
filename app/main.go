package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// See: https://help.github.com/en/actions/configuring-and-managing-workflows/using-environment-variables
	eventPath := os.Getenv("GITHUB_EVENT_PATH")
	if eventPath == "" {
		log.Fatal("GITHUB_EVENT_PATH is empty, but it's invalid")
	}

	payload, err := ioutil.ReadFile(eventPath)
	if err != nil {
		log.Fatalf("To read file(path:%s) is failed err:%s ", eventPath, err)
	}

	// ToDo: define event typp, this is not good
	var event map[string]interface{}
	err = json.Unmarshal(payload, &event)
	if err != nil {
		log.Fatalf("To unmarshal payload is failed err:%s", err)
	}

}
