/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	models "VoidDrop_V1/Models"
	service "VoidDrop_V1/Service"
	view "VoidDrop_V1/View"
	"fmt"

	"github.com/spf13/cobra"
)

// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Utilizado para enviar o arquivo",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		view.ShowLogo()

		connection := service.WebSocketConnection{}

		connection.Connect("127.0.0.1:8080", "VoidDrop")
		defer connection.Close()

		req := models.Request{
			Action: models.Create,
		}

		connection.SendMenssage(req)

		go connection.ReciveMensage()

		fmt.Scanln()
	},
}

var (
	urlFile string
)

func init() {
	rootCmd.AddCommand(sendCmd)
	sendCmd.Flags().StringVarP(&urlFile, "file", "f", "", "Informar o camnho do arquivo a ser enviado")
}


