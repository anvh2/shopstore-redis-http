package cmd

import (
	"fmt"
	"net/http"
	"shopstore/src/server/router"

	"github.com/spf13/cobra"
)

var version = "v.0.0.1"
var port = 8000

func init() {
	serviceWithPortCmd.Flags().IntVarP(&port, "port", "p", 0, "specify port to start server")

	RootCmd.AddCommand(versionCmd)
	RootCmd.AddCommand(serviceCmd)
	RootCmd.AddCommand(serviceWithPortCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Number version of service",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Version --%v", version)
		fmt.Println()
	},
}

var serviceCmd = &cobra.Command{
	Use:   "httpservice",
	Short: "Http Service",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("service is runing on http://localhost:8000")
		http.ListenAndServe(":8000", router.NewProductRouter())
	},
}

var serviceWithPortCmd = &cobra.Command{
	Use:   "httpserviceport",
	Short: "Run Http Service with specify port",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(port)

	},
}
