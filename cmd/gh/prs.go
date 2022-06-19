/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package gh

import (
	"fmt"

	tm "github.com/buger/goterm"
	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
)

// ghCmd represents the gh command
var prsCmd = &cobra.Command{
	Use:   "prs",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		listOpenPrs()
	},
}

func listOpenPrs() {
	query := "is:open+is:pr+author:jcpsimmons+archived:false"

	opts := &github.SearchOptions{
		Sort: "stars",
		ListOptions: github.ListOptions{
			PerPage: 100,
		},
	}

	result, _, err := client.Search.Issues(ctx, query, opts)

	if err != nil {
		fmt.Println(err)
	}

	prMessage := fmt.Sprintf("You have %d open PRs", result.GetTotal())
	tm.MoveCursorDown(2)
	tm.Println(tm.Background(tm.Color(tm.Bold(prMessage), tm.BLACK), tm.WHITE))
	tm.MoveCursorDown(2)
	tm.Flush()

	for _, pr := range result.Issues {
		tm.Println(tm.Bold(pr.GetTitle()))
		tm.Println(pr.GetHTMLURL())
		tm.MoveCursorDown(1)
		tm.Flush()
	}
}

func init() {
	GhCmd.AddCommand(prsCmd)

}
