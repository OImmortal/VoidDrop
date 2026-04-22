package models

import (
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
)

type Request struct {
	Action  int               `json:"action"`
	Room    string            `json:"room"`
	Payload map[string]string `json:"payload"`
}

const (
	Create       int = 0
	Join             = 1
	Offer            = 2
	Answer           = 3
	IceCandidate     = 4
)

var loading *spinner.Spinner

func ReturnLogic(response Request) {
	if response.Action == Create {
		createAction(response)
	}

	if response.Action == Join {
		joinAction()
	}

}

func createAction(response Request) {
	color.Green("Sala criada com sucesso: %s", response.Room)
	loading = spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	loading.Suffix = " Esperando conexão"
	loading.Start()
}

func joinAction() {
	loading.Stop()
}