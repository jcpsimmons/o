/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package gh

import (
	"context"
	"io/ioutil"
	"os"

	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
	"gopkg.in/yaml.v3"
)

type ConfigYaml struct {
	Tokens struct {
		Github string `yaml:"github"`
	} `yaml:"tokens"`
}

var GhCmd = &cobra.Command{
	Use:   "gh",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

var ctx = context.Background()
var client = generateClient()

func generateClient() *github.Client {
	token := getToken()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	return client
}

func getToken() string {
	dirname, _ := os.UserHomeDir()
	yfile, _ := ioutil.ReadFile(dirname + "/.o.yml")
	var config ConfigYaml
	yaml.Unmarshal(yfile, &config)
	return config.Tokens.Github
}

func Execute() {
	err := GhCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
