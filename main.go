package main

import (
	"fmt"
	"log"

	"github.com/anthdm/hollywood/actor"
)

type SetState struct {
	value uint
}

type Handler struct {
	state uint
}

type ResetState struct{}

func newHandler() actor.Producer {
	return func() actor.Receiver {
		return &Handler{}
	}
}

func (h *Handler) Receive(c *actor.Context) {
	switch msg := c.Message().(type) {
	case ResetState:
		h.state = 0
		fmt.Println("handler resetado, estado:", h.state)
	case SetState:
		h.state = msg.value
		fmt.Println("handler recebeu novo estado:", h.state)
	case actor.Initialized:
		h.state = 10
		fmt.Println("handler inicializado, estado:", h.state)
	case actor.Started:
		fmt.Println("handle started")
	case actor.Stopped:
	}
}

func main() {
	e, err := actor.NewEngine(actor.NewEngineConfig())
	if err != nil {
		log.Fatal(err)
	}
	pid := e.Spawn(newHandler(), "handler")
	for i := 0; i < 10; i++ {
		go e.Send(pid, SetState{value: uint(i)})
	}
	e.Send(pid, ResetState{})
}
