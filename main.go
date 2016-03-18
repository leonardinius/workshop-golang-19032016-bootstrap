package main

import (
	"fmt"
	"log"
	"os"

	"github.com/TransactPRO/workshop-golang-19032016-bootstrap/ui"
	"github.com/TransactPRO/workshop-golang-19032016-bootstrap/util"
)

var (
	gui             *ui.UI
	textBoxMessages = make(chan string)
)

// shutdown stops all the running services and terminates the process.
func shutdown() {
	log.Println("Shutting down..")

	os.Exit(0)
}

// writeToChatView writes the passed message to the corresponding UI view.
func writeToChatView(msg util.Message) {
	gui.WriteToView(ui.ChatView, fmt.Sprintf("[%s] %s: %s", util.ParseTime(msg.Timestamp), msg.User, msg.Contents))
}

func main() {
	var err error

	// starting the UI and passing the shutdown function which will be kicked off
	// on the Ctrl+C event the channel the messages will be written to.
	gui, err = ui.DeployGUI(shutdown, textBoxMessages)
	if err != nil {
		log.Fatal(err)
	}

	// an awesome stuff to implement.

	// holding the process with an empty select statement.
	select {}
}
