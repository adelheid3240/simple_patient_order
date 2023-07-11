/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"simplepatientorder/config"
	"simplepatientorder/internal/controller"
	"simplepatientorder/internal/handler"
	mongodb "simplepatientorder/internal/mongo"
	"simplepatientorder/internal/repository"
	"simplepatientorder/internal/server"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "list patients and crud patient orders",
	Short: "run server",
	Long:  `run server which provides crud operations for patients and orders`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.New()

		mongoClient := mongodb.GetMgoCli(cfg)
		patientRepo := repository.NewPatient(mongoClient)

		patientCtrl := controller.NewPatient(cfg, patientRepo)

		patientHandler := handler.NewPatient(patientCtrl)

		server.Run(cfg, patientHandler)
	},
}

func Execute() {
	rootCmd.Execute()
}
