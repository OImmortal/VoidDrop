package models

type Request struct {
	Action  int               `json:"action"`
	Room    string            `json:"room"`
	Payload map[string]string `json:"payload"`
}