package server

import (
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"

	"util"
)

type Server struct {
	endpoint string
	port     int
	listener net.Listener
	messages chan util.Message
}

func New(endpoint string, port int, messages chan util.Message) *Server {
	return &Server{
		endpoint: endpoint,
		port:     port,
		messages: messages,
	}
}

func (s *Server) handler(w http.ResponseWriter, r *http.Request) {
	returnHTTP := func(msg string, status int) {
		w.WriteHeader(status)
		w.Write([]byte(msg))
	}

	data, readErr := ioutil.ReadAll(r.Body)
	if readErr != nil {
		returnHTTP("invalid request", http.StatusBadRequest)
		return
	}

	var msg util.Message
	msgErr := json.Unmarshal(data, &msg)
	if msgErr != nil {
		returnHTTP(msgErr.Error(), http.StatusBadRequest)
		return
	}

	if msg.User == "" || msg.Contents == "" {
		returnHTTP("empty message", http.StatusBadRequest)
		return
	}

	s.messages <- msg
	returnHTTP("OK", http.StatusOK)
}

func (s *Server) Start() (err error) {
	http.HandleFunc(s.endpoint, s.handler)

	//var err error
	s.listener, err = net.Listen("tcp", ":"+strconv.Itoa(s.port))
	if err != nil {
		return
	}

	go http.Serve(s.listener, nil)

	return
}

//go docs test
func (s *Server) Stop() {
	s.listener.Close()
}
