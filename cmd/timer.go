/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"time"

	tm "github.com/buger/goterm"
	"github.com/golang-module/carbon/v2"
	"github.com/martinlindhe/notify"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
)

// timerCmd represents the timer command
var timerCmd = &cobra.Command{
	Use:   "timer",
	Short: "Timing functions",
	Long:  `Defaults to pomodoro - a timer that counts down to a break timer. It will send notifications as the periods switch.`,
	Run: func(cmd *cobra.Command, args []string) {
		tm.Clear()
		notify.Notify("O", "Pomodoro", "Pomodoro timer started", "")
		fmt.Println("timer called")
		pomodoro()
	},
}

func init() {
	rootCmd.AddCommand(timerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// timerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// timerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func pomodoro() {
	isWork := true

	for {
		minutes := 5

		if isWork {
			minutes = 25
		}

		deadline := carbon.Now().AddMinutes(minutes)
		bar := progressbar.Default(int64(minutes) * 60)

		for range time.Tick(1 * time.Second) {
			tm.MoveCursor(1, 1)

			printPomodoroHeader(isWork)

			humanDiff := deadline.DiffForHumans()
			secondsDiff := deadline.DiffAbsInSeconds(carbon.Now())

			fmt.Printf("next break: %s\n", humanDiff)
			bar.Set((minutes * 60) - int(secondsDiff))
			tm.Flush()
		}

		isWork = !isWork
	}

}

func printPomodoroHeader(isWork bool) {
	if isWork {
		tm.Println(tm.Background(tm.Color(tm.Bold("Time to work!"), tm.WHITE), tm.RED))
		return
	}

	tm.Println(tm.Background(tm.Color(tm.Bold("Take a break"), tm.BLACK), tm.GREEN))
}
