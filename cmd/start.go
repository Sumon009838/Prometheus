/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/Sumon009838/Prometheus/api"
	"github.com/spf13/cobra"
)

var Port string

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Prometheus",
	Long:  "run a a http server and expose the number of times it called",
	Run: func(cmd *cobra.Command, args []string) {
		api.RunServer(Port)
	},
}

func init() {
	fmt.Println("inside start.go")
	startCmd.Flags().StringVarP(&Port, "port", "p", ":9000", "http server port")
	rootCmd.AddCommand(startCmd)
}
