package command

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Short:   "Command line interface for BANK service interactions",
	Example: "bank create client --deposit 100",
	Long: `
		_____   ____   __  _  __  __ 
		| () ) / () \ |  \| ||  |/  /
		|_()_)/__/\__\|_|\__||__|\__\

	https://github.com/nikitsenka/bank-cli `,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
}
