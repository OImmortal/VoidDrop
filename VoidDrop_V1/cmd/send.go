/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
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
	Run: executeSend,
}

var (
	urlFile string
)



func init() {
	rootCmd.AddCommand(sendCmd)
	sendCmd.Flags().StringVarP(&urlFile, "file", "f", "", "Informar o camnho do arquivo a ser enviado")
}

func executeSend(cmd *cobra.Command, args []string) {
	view.ClearTerminal()
	view.ShowLogo()

	connection := service.WebSocketConnection{}

	connection.Connect("127.0.0.1:8080", "VoidDrop")
	defer connection.Close()

	connection.SendCreate()

	go connection.ReciveMensage(service.Sender)

	fmt.Scanln()
}

