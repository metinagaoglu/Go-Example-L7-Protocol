package main

import(
	"net"
	"fmt"
	"io"
	. "github.com/metinagaoglu/GolangTCPServerExample/exampleProtocol"
)

func main() {

	conn, err := net.Dial("tcp4", ":7000")
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	/*
	* Wait server response
	*/
	go func () {
		for {
			buf := make([]byte, 1024)
			_, err := conn.Read(buf)
			if err != nil {
				if err != io.EOF {
					fmt.Println("Error reading:", err)
					panic(err)
					conn.Close()
				}
				break
			}
	
			fmt.Println("Received: ", string(buf))
		}	
	}()

	/**
	* Send 10 request to server
	*/
	for i := 0; i < 10; i++ {
		data := CreateMessage(MessageTypeText, "This is my message")
		fmt.Println(data)
		_, err = conn.Write(data)
		if err != nil {
			fmt.Println("Error writing to client")
			panic(err)
		}
	
	}

	/**
	* Infinite loop handler
	*/
	for {
		select {}
	}

}