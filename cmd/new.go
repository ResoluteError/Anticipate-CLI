package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var description string

var newCmd = &cobra.Command{
	Use:     "new [date]",
	Short:   "Store a future date with a description",
	Aliases: []string{"n"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dateStr := args[0]
		_, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			fmt.Println("Please enter the date in YYYY-MM-DD format.")
			return
		}

		events := viper.GetStringMap("events")
		if events == nil {
			events = make(map[string]interface{})
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
