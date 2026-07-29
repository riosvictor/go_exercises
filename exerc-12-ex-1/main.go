/*
Crie um package "cachorro".
Este package deverá exportar uma função Idade, que toma como parâmetro um número de anos e retorna a idade equivalente em anos caninos. (1 ano humano → 7 anos caninos)
Documente seu código com comentários, e utilize a função Idade na sua função main.
Rode seu programa para verificar se ele funciona.
Rode um local server com godoc e leia sua documentação
*/

package main

import (
	"exerc-12-ex-1/cachorro"
	"fmt"
)

func main() {
	idadeHumana := 5
	idadeCanina := cachorro.Idade(idadeHumana)

	fmt.Printf("A idade de %d anos humanos é equivalente a %d anos caninos.\n", idadeHumana, idadeCanina)
}
