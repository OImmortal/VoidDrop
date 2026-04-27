package service

import (
	models "VoidDrop_V1/Models"
	view "VoidDrop_V1/View"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)


type WebSocketConnection struct {
	websocketUrl url.URL
	conn *websocket.Conn
	err error
}

const (
	Sender	= 0
	Reciver = 1
)

const (
	Create       int = 0
	Join             = 1
	Offer            = 2
	Answer           = 3
	IceCandidate     = 4
	Close 			 = 5
)



func (w *WebSocketConnection) Connect(host string, path string) {
	w.websocketUrl = url.URL{Scheme: "ws", Host: host, Path: path}
	w.conn, _, w.err = websocket.DefaultDialer.Dial(w.websocketUrl.String(), nil)

	if w.err != nil {
		log.Fatalln("Erro ao se conectar ao servidor: ", w.err)
	}
}

func (w WebSocketConnection) ReciveMensage(typeUser int) {
	for {

		var response models.Request

		err := w.conn.ReadJSON(&response)

		w.returnLogic(response, typeUser)

		if err != nil {
			log.Println("Erro ao ler mensagem: ", err)
		}
	}
}

func (w WebSocketConnection) Close() {
	if w.conn != nil {

		err := w.conn.WriteMessage(
			websocket.CloseMessage,
			websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Encerrando conexão"),
		)

		if err != nil {
			log.Println("Erro ao enviar mensagem de fechamento: ", err)
		}

		w.conn.SetReadDeadline(time.Now().Add(10 * time.Second))
		
		for {
			_, _, err := w.conn.ReadMessage()
			if err != nil {
				// Assim que der erro (o servidor fechou a porta ou deu timeout), nós saímos do loop
				break
			}
		}

		w.conn.Close()
	}
}

func (w WebSocketConnection) sendMenssage(req models.Request) {
	if w.conn != nil {
		err := w.conn.WriteJSON(req)
		if err != nil {
			log.Fatalln("Erro ao enviar mensagem: ", err)
		}
	}
}


func (w WebSocketConnection) SendCreate() {
	req := models.Request{
		Action: Create,
	}

	w.sendMenssage(req)
}
func (w WebSocketConnection) SendJoin(codeRoom string) {
	req := models.Request{
		Action: Join,
		Room: codeRoom,
	}

	w.sendMenssage(req)
}
func (w WebSocketConnection) SendOffer() {}
func (w WebSocketConnection) SendAnswer() {}
func (w WebSocketConnection) SendIceCandidate() {}

func (w WebSocketConnection) returnLogic(response models.Request, typeUser int) {

	if response.Action == Close {
		fmt.Println("Servidor foi fechado")
		return
	}

	if response.Action == Create {
		view.ReciveCreateAction(response)
		return
	}

	if response.Action == Join {
		view.ReciveJoinAction(response, typeUser)
		
		if(typeUser == Sender) {
			w.SendOffer()
		}

		if (typeUser == Reciver) {
			w.SendAnswer()
		}

		return
	}
}



