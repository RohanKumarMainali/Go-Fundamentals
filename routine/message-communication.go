package routine

import "fmt"

type Message struct {
	From    string
	Payload string
}

type Server struct {
	msgch chan Message
}

func (s *Server) StartAndListen() {
	for {
		// block till message comes
		msgch := <-s.msgch
		fmt.Printf("server is listening for %s ", msgch.From)
	}
}

func SendMessage(msgch chan Message) {
	msg := Message{
		From:    "rohan",
		Payload: "testing",
	}
	msgch <- msg
}

func MessageCommunicationMain() {
	s := &Server{
		msgch: make(chan Message),
	}

	go s.StartAndListen()
	SendMessage(s.msgch)
}
