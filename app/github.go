package main

import (
	"os"
	"fmt"
	"regexp"
	"strings"
	"strconv"
	"context"

	"github.com/google/go-github/v30/github"
)

var (
	reEmoji = regexp.MustCompile(`:.+:`)
	// ToDo: re factoring
	ctx = context.Background()
)

func Action(client *github.Client, runId string, event github.PushEvent) error {
	message := *event.HeadCommit.Message
	emoji, contains := parseMessage(message)
	// If message dose not contain emoji, ignore this commit.
	if !contains {
		return nil
	}

	switch emoji {
	case "imp":
		return makeIssue(client, event)
	case "no_entry":
		return cancelWorkflowAction(client, runId, event)
	case "mega":
		return sendNotifyToSlack(client, event)
	}
	return nil
}

func parseMessage(message string) (string, bool) {
	res := reEmoji.FindAllStringSubmatch(message, -1)
	if len(res) != 0 {
		return strings.Replace(res[0][0], ":", "", -1), true
	}
	return "", false
}

// Note
// pre-condition: parseMessage should return true
func removeEmoji(message string) string {
	emoji, _ := parseMessage(message)
	return strings.Replace(message, ":"+emoji+":", "", 1)
}

func makeIssue(client *github.Client, event github.PushEvent) error {
	comment := removeEmoji(*event.HeadCommit.Message)
	title := fmt.Sprintf("[Record] %s", comment)
	content :=
		fmt.Sprintf("\n\nI(%s) make Technical debt in %s\n\n ----- \n%s\n\n ----- \n I give you my word that I clear my debts.\n\t\t　　　　　　　　　　Date: %s",
			*event.HeadCommit.Author.Name, *event.HeadCommit.URL,
			comment,
			event.HeadCommit.Timestamp.Format("2006/Jan/2 15:04:05"))
	assignees := []string{ *event.HeadCommit.Author.Login }
	issue := new(github.IssueRequest)
	issue.Title = &title
	issue.Body = &content
	issue.Assignees = &assignees
	if _, _, err := client.Issues.Create(ctx, 
		*event.Repo.Owner.Name, 
		*event.Repo.Name, issue); err != nil {
		return fmt.Errorf("to create an issue is failed err:%s", err)
	}
	return nil
}

func cancelWorkflowAction(client *github.Client,
	runId string, event github.PushEvent) error{
	id, err := strconv.Atoi(runId)
	if err != nil{
		return fmt.Errorf("to convert runId is failed err:%s", err)
	}
	_, err = client.Actions.CancelWorkflowRunByID(ctx, 
		*event.Repo.Owner.Name, 
		*event.Repo.Name, int64(id))
	return err
}

func sendNotifyToSlack(client *github.Client,event github.PushEvent) error{
	fmt.Println("[%s]\n",os.Getenv("INPUT_SLACK_WEBHOOK_URL"))
	if url := os.Getenv("INPUT_SLACK_WEBHOOK_URL"); url != "" {
		fmt.Println("url is %s\n", url)
	}else{
		fmt.Println("pass")
	}
	return nil
}