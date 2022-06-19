/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package timer

import (
	"fmt"
	"time"

	"github.com/golang-module/carbon/v2"
	"github.com/martinlindhe/notify"

	tm "github.com/buger/goterm"

	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
)

// timerCmd represents the timer command
var pomodoroCmd = &cobra.Command{
	Use:   "pom",
	Short: "First cmd",
	Long:  `first cmd.`,
	Run: func(cmd *cobra.Command, args []string) {
		tm.Clear()
		notify.Notify("O", "Pomodoro", "Pomodoro timer started", "")
		fmt.Println("timer called")
		pomodoro()
	},
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

func init() {
	TimerCmd.AddCommand(pomodoroCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// timerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// timerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
