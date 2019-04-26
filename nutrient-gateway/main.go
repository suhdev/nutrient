package main

import (
	"fmt"
	"log"

	nats "github.com/nats-io/go-nats"
)

func main() {

	nc, _ := nats.Connect("nats://nutrient-nats:4222")
	ch := make(chan struct{})
	nc.Subscribe("box", func(m *nats.Msg) {
		fmt.Println(string(m.Data))
		log.Println(string(m.Data))
	})
	<-ch
}
