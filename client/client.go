package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func startClient(address string, port int) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", address, port))
	if err != nil {
		fmt.Println("Erro ao conectar:", err)
		return
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Digite uma mensagem: ")
	message, _ := reader.ReadString('\n')

	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Erro ao enviar mensagem:", err)
		return
	}
}

func main() {
	const destination = "localhost"
	const port = 6000
	startClient(destination, port)
}
