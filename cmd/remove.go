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
			fmt.Println("No events to remove.")
			return
		}

		var dateStr string

		if len(args) == 1 {
			dateStr = args[0]
		} else {
			dateStr = promptForDate(events)
			if dateStr == "" {
				fmt.Println("No date selected.")
				return
			}
		}

		if _, exists := events[dateStr]; !exists {
			fmt.Println("Date not found.")
			return
		}

		delete(events, dateStr)
		viper.Set("events", events)
		viper.WriteConfig()

		fmt.Printf("Removed event on %s.\n", dateStr)
	},
}

func promptForDate(events map[string]string) string {
	dates := make([]string, 0, len(events))
	for dateStr := range events {
		dates = append(dates, dateStr)
	}

	sort.Strings(dates)

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}:",
		Active:   "> {{ . | cyan }}",
		Inactive: "  {{ . | cyan }}",
		Selected: "> {{ . | green }}",
	}

	prompt := promptui.Select{
		Label:     "Select date to remove",
		Items:     dates,
		Templates: templates,
	}

	_, result, err := prompt.Run()
	if err != nil {
		return ""
	}

	return result
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
