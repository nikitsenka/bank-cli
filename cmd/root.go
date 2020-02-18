package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Short:   "Command line interface for ODIS service interactions",
	Example: "odis create-registration -n registration-name",
	Long: `

                    Nike ED&A EMEA

           ____     ______      _____    _____
          / __ \   (_  __ \    (_   _)  / ____\
         / /  \ \    ) ) \ \     | |   ( (___
        ( ()  () )  ( (   ) )    | |    \___ \
        ( ()  () )   ) )  ) )    | |        ) )
         \ \__/ /   / /__/ /    _| |__  ___/ /
          \____/   (______/    /_____( /____/

                https://odis.nike.com/
              Slack: #help-onboarding-odis`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
}
