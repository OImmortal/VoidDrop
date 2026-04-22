package service

import (
	models "VoidDrop_V1/Models"
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

func (w *WebSocketConnection) Connect(host string, path string) {
	w.websocketUrl = url.URL{Scheme: "ws", Host: host, Path: path}
	w.conn, _, w.err = websocket.DefaultDialer.Dial(w.websocketUrl.String(), nil)

	if w.err != nil {
		log.Fatalln("Erro ao se conectar ao servidor: ", w.err)
	}
}

func (w WebSocketConnection) ReciveMensage() {
	for {

		var response models.Request

		err := w.conn.ReadJSON(&response)

		models.ReturnLogic(response)

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

func (w WebSocketConnection) SendMenssage(req models.Request) {
	if w.conn != nil {
		err := w.conn.WriteJSON(req)
		if err != nil {
			log.Fatalln("Erro ao enviar mensagem: ", err)
		}
	}
}
