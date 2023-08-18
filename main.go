package main

import (
	"fmt"
	"log"
	"main/config"
	"os"

	env "github.com/joho/godotenv"
)

func main() {
	err := env.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file. Err: %s", err)
	}

	jiraClient, err := config.NewJiraClient()
	if err != nil {
		log.Fatalf("Error creating Jira client. Err: %s", err)
	}

	issue, res, err := jiraClient.Issue.Get(os.Getenv("JIRA_TEST_ISSUE_ID"), nil)
	if err != nil {
		log.Fatalf("Error loading issue! Err: %s, Res: %s", err, res.Request.URL)
	}

	fmt.Printf("%s: %+v\n", issue.Key, issue.Fields.Summary)
	fmt.Printf("Priority: %s\n", issue.Fields.Priority.Name)
}
