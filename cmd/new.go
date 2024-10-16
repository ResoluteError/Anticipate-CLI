package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var description string

var newCmd = &cobra.Command{
	Use:     "new",
	Short:   "Create a new event to anticipate",
	Aliases: []string{"n", "a", "add"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dateStr := args[0]
		date, err := time.Parse("2006-01-02", dateStr)

		if err != nil {
			cobra.CheckErr("Invalid date format. Please use YYYY-MM-DD")
		}

		if date.Before(time.Now()) {
			cobra.CheckErr("Date cannot be in the past")
		}

		events := viper.GetStringMap("events")

		if events == nil {
			events = make(map[string]interface{})
		}

		if _, found := events[dateStr]; found {
			cobra.CheckErr("Event already exists on this date")
		}

		events[dateStr] = description
		viper.Set("events", events)
		viper.WriteConfig()
		fmt.Printf("Stored: %s - %s\n", dateStr, description)
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
	newCmd.Flags().StringVarP(&description, "description", "d", "", "Description of the event")
	newCmd.MarkFlagRequired("description")
}
