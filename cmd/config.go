/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"simplepatientorder/config"

	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "preview config",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.New()
		log.Printf("config.Config: %+v\n", cfg)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
