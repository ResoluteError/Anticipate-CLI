package cmd

import (
	"fmt"
	"sort"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var removeCmd = &cobra.Command{
	Use:     "remove [date]",
	Short:   "Remove a stored date",
	Aliases: []string{"rm"},
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		events := viper.GetStringMapString("events")
		if len(events) == 0 {
			cobra.CheckErr("No events stored.")
		}

		var dateStr string

		if len(args) == 1 {
			dateStr = args[0]
		} else {
			var err error
			dateStr, err = promptForDate(events)
			cobra.CheckErr(err)
		}

		if _, exists := events[dateStr]; !exists {
			cobra.CheckErr("Date not found.")
		}

		delete(events, dateStr)

		viper.Set("events", events)
		viper.WriteConfig()

		fmt.Printf("Removed event on %s.\n", dateStr)

	}}

func promptForDate(events map[string]string) (string, error) {

	dates := make([]string, 0, len(events))

	for dateStr := range events {
		dates = append(dates, dateStr)
	}

	sort.Strings(dates)

	prompt := promptui.Select{
		Label: "Select a date to remove",
		Items: dates,
	}

	_, result, err := prompt.Run()

	return result, err

}

func init() {
	rootCmd.AddCommand(removeCmd)
}
