/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"simplepatientorder/config"
	"simplepatientorder/internal/server"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "list patients and crud patient orders",
	Short: "run server",
	Long:  `run server which provides crud operations for patients and orders`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.New()

		server.Run(cfg)
	},
}

func Execute() {
	rootCmd.Execute()
}
