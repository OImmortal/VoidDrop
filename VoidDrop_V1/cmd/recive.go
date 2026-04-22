/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// reciveCmd represents the recive command
var reciveCmd = &cobra.Command{
	Use:   "recive",
	Short: "Utilizado para receber o arquivo",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("recive called")
	},
}

func init() {
	rootCmd.AddCommand(reciveCmd)
}
