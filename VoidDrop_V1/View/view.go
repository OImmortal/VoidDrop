package view

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var logo string = `
██╗   ██╗ ██████╗ ██╗██████╗ ██████╗ ██████╗  ██████╗ ██████╗ 
██║   ██║██╔═══██╗██║██╔══██╗██╔══██╗██╔══██╗██╔═══██╗██╔══██╗
██║   ██║██║   ██║██║██║  ██║██║  ██║██████╔╝██║   ██║██████╔╝
╚██╗ ██╔╝██║   ██║██║██║  ██║██║  ██║██╔══██╗██║   ██║██╔═══╝ 
 ╚████╔╝ ╚██████╔╝██║██████╔╝██████╔╝██║  ██║╚██████╔╝██║     
  ╚═══╝   ╚═════╝ ╚═╝╚═════╝ ╚═════╝ ╚═╝  ╚═╝ ╚═════╝ ╚═╝     
                                                              
`

func ShowLogo() {
	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#fab636")).
		Bold(true)

	fmt.Println(style.Render(logo))
}