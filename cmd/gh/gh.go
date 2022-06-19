/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package gh

import (
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

type ConfigYaml struct {
	Tokens struct {
		Github string `yaml:"github"`
	} `yaml:"tokens"`
}

var Token = getToken()

// ghCmd represents the gh command
var GhCmd = &cobra.Command{
	Use:   "gh",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
