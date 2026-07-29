package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	canal := converge(trabalho("A"), trabalho("B"), trabalho("C"))

	// Recebendo continuamente do canal, mas sem limite de mensagens
	// for v := range canal {
	// 	fmt.Println(v)
	// }
	// Limitando a quantidade de mensagens recebidas do canal
	for x := 0; x < 10; x++ {
		fmt.Println(<-canal)
	}
}

func trabalho(text string) chan string {
	canal := make(chan string)
	go func(s string, c chan string) {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("Função %v diz: %v", s, i)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
		}
	}(text, canal)
	return canal
}

func converge(canais ...chan string) chan string {
	canalConverge := make(chan string)
	for _, c := range canais {
		go func(canal chan string) {
			for v := range canal {
				canalConverge <- v
			}
		}(c)
	}
	return canalConverge
}
