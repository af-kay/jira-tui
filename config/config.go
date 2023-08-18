package config

import (
	"log"
	"os"

	"github.com/andygrunwald/go-jira"
)

type JiraConfig struct {
	Url      string
	Username string
	Password string
}

func NewJiraClient() (*jira.Client, error) {
    jiraCfg := New()

    tp := jira.BasicAuthTransport{
		Username: jiraCfg.Username,
		Password: jiraCfg.Password,
	}

	return jira.NewClient(tp.Client(), jiraCfg.Url)
}

func New() *JiraConfig {
	keyJiraUrl := "JIRA_URL"
	keyJiraUsername := "JIRA_USERNAME"
	keyJiraPassword := "JIRA_PASSWORD"
	errorFormat := "Missing %s env variable! Make sure you have loaded .env file or variables by yourself"

	url, isKeyExist := os.LookupEnv(keyJiraUrl)
	if !isKeyExist {
		log.Fatalf(errorFormat, keyJiraUrl)
	}

	username, isKeyExist := os.LookupEnv(keyJiraUsername)
	if !isKeyExist {
		log.Fatalf(errorFormat, keyJiraUsername)
	}

	password, isKeyExist := os.LookupEnv(keyJiraPassword)
	if !isKeyExist {
		log.Fatalf(errorFormat, keyJiraPassword)
	}

	return &JiraConfig{
		Username: username,
		Url:      url,
		Password: password,
	}
}
