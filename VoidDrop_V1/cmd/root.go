/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "VoidDrop_V1",
	Short: "Cliente CLI P2P para transferência segura de arquivos massivos via WebRTC, com travessia de NAT por UDP hole punching.",
	Long: `VoidDrop é a engine de transferência do ecossistema 
VoidDrop: uma aplicação de linha de comando peer-to-peer
voltada a arquivos muito grandes (ordem de dezenas de GB) sem carregar o conteúdo inteiro em RAM — leitura e escrita
ocorrem em fluxos por chunks, com suporte a retomada após queda de sessão.

Contrato de mensagem: {"action":"...","room":"...","payload":...} com payload adiado (json.RawMessage no Go) para
SDP ou candidatos ICE conforme a fase da negociação.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
