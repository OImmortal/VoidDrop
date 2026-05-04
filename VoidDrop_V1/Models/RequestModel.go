package models

import "github.com/pion/webrtc/v3"

type Request struct {
	Action  int    `json:"action"`
	Room    string `json:"room"`
	Payload webrtc.SessionDescription `json:"payload"`
	IceCandidate webrtc.ICECandidateInit `json:"icecandidate"`
}