### Desafio Ping Pong

Criada a pasta pingpong dentro da pasta src, aberto o VS Code, criado o arquivo pingpong.go com o seguinte código:

```go
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
```

Depois foi aberto o terminal e dados os comandos: “go build pingpong.go” e “go run pingpong.go”, exibindo na tela as mensagens:

“pong
ping
pong
ping
pong
pong
…”

Neste exemplo, você tem duas funções, ping() e pong(), que enviam repetidamente as palavras "ping" e "pong" para um canal de strings. Em seguida, no main(), você inicia as goroutines para executar as funções ping() e pong(). Por fim, um loop infinito recebe as palavras do canal e as imprime em alternância na saída.

A ordem exata das palavras "ping" e "pong" será determinada pela concorrência e pode variar a cada execução. O programa continuará imprimindo em alternância indefinidamente até que seja interrompido.

Foram adicionados dois atrasos usando a função Sleep() do pacote time. A função Sleep() pausa a execução por um determinado período de tempo, especificado em milissegundos.

No caso da função ping(), após enviar a palavra "ping" para o canal, há um atraso de 500 milissegundos usando time.Sleep(500 _ time.Millisecond). Já na função pong(), após enviar a palavra "pong", há um atraso de 300 milissegundos usando time.Sleep(300 _ time.Millisecond).

Esses atrasos podem ser ajustados conforme necessário para criar o ritmo desejado entre as palavras "ping" e "pong".
