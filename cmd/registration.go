package cmd

import "github.com/spf13/cobra"

var registration = &cobra.Command{
	Use:   "registration",
	Short: "Contain various subcommands for registrations managements",
}

func init() {
	rootCmd.AddCommand(registration)
}
