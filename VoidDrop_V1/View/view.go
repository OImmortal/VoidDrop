package view

import (
	models "VoidDrop_V1/Models"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/briandowns/spinner"
	"github.com/charmbracelet/lipgloss"
	"github.com/fatih/color"
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


func ClearTerminal() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}

var loading *spinner.Spinner

func ReciveCreateAction(response models.Request) {
	color.Green("┌──────────────────────────────────────────────┐")
	color.Green("│ Sala criada com sucesso!                     │")
	color.Green("├──────────────────────────────────────────────┤")
	color.Green("│ Código da sala: %s              │", response.Room)
	color.Green("└──────────────────────────────────────────────┘")

	loading = spinner.New(spinner.CharSets[40], 100*time.Millisecond)
	loading.Suffix = " Aguardando outro participante entrar na sala..."
	loading.Color("cyan")
	loading.Start()
}

func ReciveJoinAction(response models.Request, typeUser int) {

	if (loading != nil) {
		loading.Stop()
	} 
	
	ClearTerminal()

	ShowLogo()

	color.Green("┌──────────────────────────────────────────────┐")
	color.Green("│           Sala encontrada com sucesso!       │")
	color.Green("├──────────────────────────────────────────────┤")
	color.Green("│      Conexão entre os usuários realizada!    │")
	color.Green("└──────────────────────────────────────────────┘")

	loading = spinner.New(spinner.CharSets[40], 100*time.Millisecond)
	loading.Suffix = " Estabelecendo conexão máquina a máquina..."
	loading.Color("yellow")
	loading.Start()

}