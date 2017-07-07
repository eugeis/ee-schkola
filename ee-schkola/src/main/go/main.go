package main

import (
	"fmt"
	"ee/schkola"
	"net/http"
)

func main() {

}

func StartChat() {
	fmt.Println("Starting application...")
	manager := schkola.NewClientManager()
	go manager.Start()
	http.HandleFunc("/ws", manager.Handler)
	http.ListenAndServe(":12345", nil)
}

func StartMailer() {
	mailer := schkola.Mailer{}
	mailer.SendRegistration()
}
