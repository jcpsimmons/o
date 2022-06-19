/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package gh

import (
	"fmt"

	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
)

// ghCmd represents the gh command
var meCmd = &cobra.Command{
	Use:   "me",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// have an arg to not mark them read, should mark read by default
		showMentions()
	},
}

func showMentions() {

	opts := &github.NotificationListOptions{
		All: true,
	}

	result, _, err := client.Activity.ListNotifications(ctx, opts)

	if err != nil {
		fmt.Println(err)
	}

	mentions := make([]*github.Notification, 0)

	for _, notif := range result {
		if notif.GetReason() == "mention" {
			mentions = append(mentions, notif)
		}
	}

	// now mentions are in a new slice, we can format,print,mark unread

}

func init() {
	GhCmd.AddCommand(meCmd)

}
