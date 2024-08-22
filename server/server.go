package main

import (
	"bufio"
	"fmt"
	"net"
)

func startServer(address string, port int) {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", address, port))
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
		return
	}
	defer listener.Close()

	fmt.Printf("Servidor rodando.. %s:%d\n", address, port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Erro ao aceitar conexão:", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	data, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Erro ao ler dados:", err)
		return
	}

	fmt.Printf("Mensagem: %s", data)
}

func main() {
	const host = "localhost"
	const port = 6000

	startServer(host, port)
}
