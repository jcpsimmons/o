/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package timer

import (
	"os"

	"github.com/spf13/cobra"
)

// TimerCmd represents the timer command
var TimerCmd = &cobra.Command{
	Use:   "timer",
	Short: "Timing functions",
	Long:  `Defaults to pomodoro - a timer that counts down to a break timer. It will send notifications as the periods switch.`,
}

func Execute() {
	err := TimerCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
