/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	service "VoidDrop_V1/Service"
	view "VoidDrop_V1/View"
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// reciveCmd represents the recive command
var reciveCmd = &cobra.Command{
	Use:   "recive",
	Short: "Utilizado para receber o arquivo",
	Long: ``,
	Run: executeRecive,
}

var (
	codeRoom string
)

func init() {
	rootCmd.AddCommand(reciveCmd)
	reciveCmd.Flags().StringVarP(&codeRoom, "room", "r", "", "Informar codigo da sala criada")
}



func executeRecive(cmd *cobra.Command, args []string) {
	view.ClearTerminal()
	view.ShowLogo()


	if (codeRoom == "") {
		color.Red("┌──────────────────────────────────────────────┐")
		color.Red("│   Você deve informar um código da sala!       │")
		color.Red("├──────────────────────────────────────────────┤")
		color.Red("│ Use o parâmetro '-r <codigo>' para informar.  │")
		color.Red("└──────────────────────────────────────────────┘")
		return 
	}

	connection := service.WebSocketConnection{}

	connection.Connect("127.0.0.1:8080", "VoidDrop")
	defer connection.Close()

	connection.SendJoin(codeRoom)

	go connection.ReciveMensage(service.Reciver)

	fmt.Scanln()
}
