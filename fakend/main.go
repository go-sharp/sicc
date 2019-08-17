package main

import (
	"fmt"
	"log"
	"net"

	"github.com/hypebeast/go-osc/osc"
)

func main() {
	addr := "127.0.0.1:8765"
	server := &osc.Server{Addr: addr}

	conn, err := net.ListenPacket("udp", addr)
	if err != nil {
		log.Fatalln("Couldn't listen: ", err)
	}
	defer conn.Close()

	go func() {
		for {
			packet, err := server.ReceivePacket(conn)
			if err != nil {
				log.Println("Failed to receive packet:", err)
				continue
			}

			if packet != nil {
				switch packet.(type) {
				default:
					fmt.Println("Unknow packet type!")

				case *osc.Message:
					fmt.Printf("-- OSC Message: ")
					osc.PrintMessage(packet.(*osc.Message))

				case *osc.Bundle:
					fmt.Println("-- OSC Bundle:")
					bundle := packet.(*osc.Bundle)
					for i, message := range bundle.Messages {
						fmt.Printf("  -- OSC Message #%d: ", i+1)
						osc.PrintMessage(message)
					}
				}
			}
		}
	}()

	log.Printf("Listening osc server on: %v", addr)
	log.Println("Press Ctrl-C to stop server")
	ch := make(chan struct{}, 0)
	<-ch
}
