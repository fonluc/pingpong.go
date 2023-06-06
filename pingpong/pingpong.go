package main

import (
	"fmt"
	"time"
)

func ping(c chan string) {
	for {
		c <- "ping"                        // Envia a palavra "ping" para o canal
		time.Sleep(500 * time.Millisecond) // Aguarda 500 milissegundos
	}
}

func pong(c chan string) {
	for {
		c <- "pong"                        // Envia a palavra "pong" para o canal
		time.Sleep(300 * time.Millisecond) // Aguarda 300 milissegundos
	}
}

func main() {
	c := make(chan string) // Cria um canal de strings

	go ping(c) // Inicia uma goroutine para executar a função ping()
	go pong(c) // Inicia uma goroutine para executar a função pong()

	for {
		fmt.Println(<-c) // Recebe uma palavra do canal e imprime na saída
	}
}
