package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"client"
	"server"
	"ui"
	"util"
)

var (
	masterMode = flag.Bool("m", false, "master flag")
	username   = flag.String("u", "<<undefined>>", "username")
)

var (
	gui             *ui.UI
	textBoxMessages = make(chan string)
	clientMessages  = make(chan util.Message)
	s               *server.Server
)

// shutdown stops all the running services and terminates the process.
func shutdown() {
	log.Println("My Shutting down..")

	s.Stop()

	os.Exit(0)
}

// writeToChatView writes the passed message to the corresponding UI view.
func writeToChatView(msg util.Message) {
	gui.WriteToView(ui.ChatView, fmt.Sprintf("[%s] %s: %s", util.ParseTime(msg.Timestamp), msg.User, msg.Contents))
}

type myhandler func(util.Message) error

func processMyMessages(h myhandler) error {
	for msgStr := range textBoxMessages {
		msg := util.Message{
			User:      *username,
			Contents:  msgStr,
			Timestamp: time.Now(),
		}

		err := h(msg)
		if err != nil {
			return err
		}
	}

	return nil
}

func processClientMessages() {
	for msg := range clientMessages {
		writeToChatView(msg)
	}
}

func main() {
	flag.Parse()

	var err error

	// starting the UI and passing the shutdown function which will be kicked off
	// on the Ctrl+C event the channel the messages will be written to.
	gui, err = ui.DeployGUI(shutdown, textBoxMessages)
	if err != nil {
		log.Fatal(err)
	}

	if *masterMode {
		s = server.New("/", 8081, clientMessages)
		serverErr := s.Start()
		if serverErr != nil {
			log.Fatal(serverErr)
		}
		go processClientMessages()
	} else {

	}

	go processMyMessages(func(msg util.Message) error {
		writeToChatView(msg)

		if !*masterMode {
			client.Send(msg, "127.0.0.1", 8081, "/msg")
		}

		return nil
	})

	// holding the process with an empty select statement.
	select {}
}
