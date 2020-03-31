package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"fmt"

	"github.com/google/go-github/v30/github"
)

func main() {
	if eventType := os.Getenv("GITHUB_EVENT_NAME"); eventType != "push" {
		log.Fatalf("Unsupported eventType(%s), only 'push' is supported", eventType)
	}

	// See: https://help.github.com/en/actions/configuring-and-managing-workflows/using-environment-variables
	eventPath := os.Getenv("GITHUB_EVENT_PATH")
	if eventPath == "" {
		log.Fatal("GITHUB_EVENT_PATH is empty, but it's invalid")
	}

	payload, err := ioutil.ReadFile(eventPath)
	if err != nil {
		log.Fatalf("To read file(path:%s) is failed err:%s ", eventPath, err)
	}

	// https://github.com/google/go-github/blob/bf4e9281481bcbc811e5f8001c18cbe11613bffd/github/event_types.go
	var event github.PushEvent
	err = json.Unmarshal(payload, &event)
	if err != nil {
		log.Fatalf("To unmarshal payload is failed err:%s", err)
	}

	fmt.Printf("[test] commit is %s", event.HeadCommit.Message)
}
