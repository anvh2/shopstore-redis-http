package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

//RootCmd -
var RootCmd = &cobra.Command{
	Use:   "HttpService",
	Short: "HttpService command",
}

//Execute -
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
