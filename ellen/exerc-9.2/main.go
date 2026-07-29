package main

import (
	"fmt"
)

/*
Esse exercício vai reforçar seus conhecimentos sobre conjuntos de métodos.
Crie um tipo para um struct chamado "pessoa"
Crie um método "falar" para este tipo que tenha um receiver ponteiro (*pessoa)
Crie uma interface, "humanos", que seja implementada por tipos com o método "falar"
Crie uma função "dizerAlgumaCoisa" cujo parâmetro seja do tipo "humanos" e que chame o método "falar"
Demonstre no seu código:
Que você pode utilizar um valor do tipo "*pessoa" na função "dizerAlgumaCoisa"
Que você não pode utilizar um valor do tipo "pessoa" na função "dizerAlgumaCoisa"
Se precisar de dicas, veja: https://gobyexample.com/interfaces
*/

type Pessoa struct {
	nome  string
	idade int
}

func (p *Pessoa) falar() {
	fmt.Printf("Olá, meu nome é %s e tenho %d anos.\n", p.nome, p.idade)
}

type Humanos interface {
	falar()
}

func dizerAlgumaCoisa(h Humanos) {
	h.falar()
}

func main() {
	p := Pessoa{nome: "Paulo", idade: 30}
	p.falar()    // atalho para "(&p).falar()"
	(&p).falar() // forma mais correta de chamar o método com receiver ponteiro

	dizerAlgumaCoisa(&p) // Passando um ponteiro para a função dizerAlgumaCoisa
	// dizerAlgumaCoisa(p) // Isso não compilará, pois p é do tipo Pessoa, não *Pessoa
}
