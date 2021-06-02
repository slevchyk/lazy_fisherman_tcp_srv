package main

import (
	"io"
	"net"
)

func main() {

	l, err := net.Listen("tcp", ":8002")
	if err != nil {
		panic(err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			panic(err)
		}

		for {
			bs := make([]byte, 1024)
			n, err := conn.Read(bs)
			if err != nil {
				break
			}

			request := string(bs[:n])

			if request == "close" {
				conn.Close()
				break
			}

			answer := request + "-x\n"

			io.WriteString(conn, answer)
		}

		conn.Close()
	}
}
