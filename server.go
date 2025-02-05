package main

import (
	"io"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

type Server struct {
	listenAddr string
	router     *http.ServeMux
}

func (s *Server) HandleWebSocket(ws *websocket.Conn) {
	log.Printf("New incoming websocket connection %s", ws.RemoteAddr())

	// remember to add mutex
	// s.ws[conn] = ws.RemoteAddr()
	s.readFromWs(ws)
}
func (s *Server) readFromWs(ws *websocket.Conn) {
	buff := make([]byte, 1024)
	for {
		n, err := ws.Read(buff)
		if err != nil {
			if err == io.EOF {
				// means conn on other side close itself
				break
			}
			log.Println("read error: ", err)
			continue
		}
		msg := buff[:n]
		log.Println("msg received: ", string(msg))
		ws.Write([]byte("Thank you for the message!"))

	}
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
	}
}

func (s *Server) Start() {
	// router := http.NewServeMux()

	// router.Handle("/ws", websocket.Handler(s.HandleWebSocket))

	// user := NewUserHandlers()

	// router.HandleFunc("POST /register", handler(user.handleRegister))
	// router.HandleFunc("POST /login", handler(user.handleLogin))
	// // onlyAdmin(handler(admin.handle))

	// log.Printf("Listening on %s\n", listenAddr)

	// http.ListenAndServe(listenAddr, router)
	//

}
