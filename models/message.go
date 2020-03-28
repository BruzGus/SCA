package models

//Message para el manejo de los mensajes HTTP
type Message struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}
