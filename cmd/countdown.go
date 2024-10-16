package cmd

import (
	"fmt"
	"sort"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var countdownCmd = &cobra.Command{
	Use:     "countdown",
	Short:   "List all future dates and countdown",
	Aliases: []string{"cd", "ls", "list"},
	Run: func(cmd *cobra.Command, args []string) {

		events := viper.GetStringMapString("events")

		if len(events) == 0 {
			cobra.CheckErr("No events found")
		}

		dates := make([]string, 0, len(events))
		for dateStr := range events {
			dates = append(dates, dateStr)
		}

		sort.Strings(dates)

		now := time.Now()

		for _, dateStr := range dates {
			date, err := time.Parse("2006-01-02", dateStr)
			if err != nil {
				continue
			}

			if date.Before(now) {
				continue
			}

			daysUntil := int(date.Sub(now).Hours() / 24)
			fmt.Printf("%s (%s): %d days remaining\n", dateStr, events[dateStr], daysUntil)
		}

	}}

func init() {
	rootCmd.AddCommand(countdownCmd)
}
