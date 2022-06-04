package main

import (
	"net"
	"fmt"
	"os"
	"io"
	"os/signal" //https://pkg.go.dev/os/signal
	"runtime/pprof"
	. "github.com/metinagaoglu/GolangTCPServerExample/exampleProtocol"
)

func main()  {

	f, err := os.Create("cpu.out")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	//TODO: use enviroment for port
	ls, err := net.Listen("tcp", ":7000")
	if err != nil {
		panic(err)
	}

	fmt.Println("Server started")

	// Concurrent
	go func(){
		for {
			conn, err := ls.Accept()
			if err != nil {
				panic(err)
			}
			go handler(conn)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

}

func handler(conn net.Conn) {
	fmt.Println("Connection accepted", conn.RemoteAddr().String())

	for {
		// It reads n bytes of any packet from the Client.
		buf := make([]byte, 1024) // Read it with an n-byte buffer
		// The stream can come from 1500 to 1500. We keep everything in the array and record it. (Buffer)
		_, err := conn.Read(buf) //It reads data from the Ethernet card from the ring buffer and gives you as much as you can get.

		if err != nil {
			if err != io.EOF {
				fmt.Println("Error reading:", err)
				panic(err)
				conn.Close()
			}
			break
		}

		fmt.Println("Received: ", string(buf))
		//appendFile("/var/log/gotcp_access.txt", buf)

		mtype, _, msg := ReadMessage(buf)

		fmt.Printf("Message type: %d, message: %s\n", mtype, msg)

		_ , err = conn.Write(buf); // Like echo server , we send back what we can.
		if err != nil {
			fmt.Println("Error writing to client")
			panic(err)
			conn.Close()
			break
		}

	}
}

func appendFile(filename string, data []byte) error {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}

	defer f.Close()

	if _, err = f.Write(data); err != nil {
		return err
	}

	return nil
}