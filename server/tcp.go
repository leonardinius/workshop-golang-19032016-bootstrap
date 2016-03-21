package server

import (
	"bufio"
	"encoding/json"
	"log"
	"net"
	"strconv"
	"strings"
	// "net/tcp"
)

type Listener struct {
	tcpListener       *net.TCPListener
	activeConnecitons map[string]*net.TCPConn
	cn                ConnectionNotifier
	username          string
}

type ConnectionNotifier func(string, bool)

func NewListener(port int, username string, cn ConnectionNotifier) (l *Listener, err error) {
	tcpAddr, addrErr := net.ResolveTCPAddr("tcp4", ":"+strconv.Itoa(port))
	if addrErr != nil {
		return
	}

	var tcpListener *net.TCPListener
	tcpListener, err = net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		return
	}

	l = &Listener {
		tcpListener:       tcpListener,
		activeConnecitons: make(map[string]*net.TCPConn),
		cn:                cn,
		username:          username,
	}

	return
}

func (l *Listener) getActiveUsers() (data []byte, err error) {
	users := []string{l.username}
	for user := range l.activeConnecitons {
		users = append(users, user)
	}

	data, err = json.Marshal(users)

	return
}

func (l *Listener) Start() {
	go func() {
		for {
			tcpConn, connErr := l.tcpListener.AcceptTCP()
			if connErr != nil {
				continue
			}

			result, resultErr := bufio.NewReader(tcpConn).ReadString('\n')
			if resultErr != nil {
				log.Println("failed to establish TCP connection")
				continue
			}

			username := strings.Replace(result, "\n", "", -1)
			l.cn(username, true)
			if _, e := l.activeConnecitons[username]; e {
				continue
			}
			l.activeConnecitons[username] = tcpConn
		}
	}()
}
