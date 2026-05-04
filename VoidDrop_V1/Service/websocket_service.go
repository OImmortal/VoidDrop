package service

import (
	models "VoidDrop_V1/Models"
	view "VoidDrop_V1/View"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
	"github.com/pion/webrtc/v3"
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

var (
	config *webrtc.Configuration
	peerConnection *webrtc.PeerConnection
	genericErro error
	roomCodeVar *string
)

func (w *WebSocketConnection) Connect(host string, path string) {
	w.websocketUrl = url.URL{Scheme: "ws", Host: host, Path: path}
	w.conn, _, w.err = websocket.DefaultDialer.Dial(w.websocketUrl.String(), nil)

	if w.err != nil {
		log.Fatalln("Erro ao se conectar ao servidor: ", w.err)
	}

	config = &webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{"stun:stun.l.google.com:19302"},
			},
		},
	}

	peerConnection, genericErro = webrtc.NewPeerConnection(*config)

	if genericErro != nil {
		log.Fatal("Erro ao se conectar com o ICEServer", genericErro)
	}

	peerConnection.OnICECandidate(func(c *webrtc.ICECandidate) {
		if c == nil {
			return
		}

		cadidateJson := c.ToJSON()

		if roomCodeVar != nil {
			
			req := models.Request{
				Action: IceCandidate,
				Room: *roomCodeVar,
				IceCandidate: cadidateJson,
			}
			w.sendMenssage(req)
			// fmt.Println("Enviando Candidate")
		}
	})

	peerConnection.OnConnectionStateChange(func(s webrtc.PeerConnectionState) {
		log.Printf("Estado da conexão WebRTC mudou para: %s\n", s.String())

		if s == webrtc.PeerConnectionStateConnected {
			log.Println("🚀 SUCESSO! AS MÁQUINAS ESTÃO CONECTADAS DIRETAMENTE (P2P)!")
		} else if s == webrtc.PeerConnectionStateFailed || s == webrtc.PeerConnectionStateClosed {
			log.Println("❌ A conexão falhou ou foi encerrada.")
		}
	})

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
		Payload: webrtc.SessionDescription{},
	}

	w.sendMenssage(req)
}
func (w WebSocketConnection) SendJoin(codeRoom string) {

	roomCodeVar = &codeRoom

	req := models.Request{
		Action: Join,
		Room: codeRoom,
		Payload: webrtc.SessionDescription{},
	}

	w.sendMenssage(req)
}
func (w WebSocketConnection) SendOffer(codeRoom string) {
	if peerConnection != nil {
		offer, genericErro := peerConnection.CreateOffer(nil)
		if genericErro != nil {
			log.Fatal("Erro ao criar offer", genericErro)
		}

		peerConnection.SetLocalDescription(offer)

		req := models.Request{
			Action: Offer,
			Room: codeRoom,
			Payload: offer,
		}

		w.sendMenssage(req)
	}
}
func (w WebSocketConnection) SendAnswer(codeRoom string) {
	if peerConnection != nil {
		answer, genericErro := peerConnection.CreateAnswer(nil)
		if genericErro != nil {
			log.Fatal("Erro ao criar Answer", genericErro)
		}

		peerConnection.SetLocalDescription(answer)

		req := models.Request{
			Action: Answer,
			Room: codeRoom,
			Payload: answer,
		}

		w.sendMenssage(req)
	}
}
func (w WebSocketConnection) SendIceCandidate() {

}

func (w WebSocketConnection) returnLogic(response models.Request, typeUser int) {

	if response.Action == Close {
		fmt.Println("Servidor foi fechado")
		return
	}

	if response.Action == Create {
		view.ReciveCreateAction(response)

		if response.Room != "" {
			roomCodeVar = &response.Room
		}

		return
	}

	if response.Action == Join {
		view.ReciveJoinAction(response, typeUser)

		if typeUser == Sender {

			if roomCodeVar != nil {
				w.SendOffer(*roomCodeVar)
			}
		}

		return
	}

	if response.Action == Offer {
		if typeUser == Reciver {
			peerConnection.SetRemoteDescription(response.Payload)
			w.SendAnswer(response.Room)
		}
	}

	if response.Action == Answer {
		if typeUser == Sender {
			peerConnection.SetRemoteDescription(response.Payload)
		}
	}

	if response.Action == IceCandidate {
		err := peerConnection.AddICECandidate(response.IceCandidate)
        if err != nil {
			log.Println("Erro ao adicionar ICE Candidate remoto:", err)
        }
		// fmt.Println("Adicionando Ice")
	}
}



