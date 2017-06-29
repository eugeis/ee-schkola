package main

import (
	"fmt"
	"net/http"
	"schkola/schkola"
)

func main() {
	//startChat()
	startMailer()
}

func startChat() {
	fmt.Println("Starting application...")
	manager := schkola.NewClientManager()
	go manager.Start()
	http.HandleFunc("/ws", manager.Handler)
	http.ListenAndServe(":12345", nil)
}

func startMailer() {
	mailer := schkola.Mailer{}
	mailer.SendRegistration()
}
