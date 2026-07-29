package cachorro

/*
Idade recebe a idade em anos humanos e retorna a idade equivalente em anos caninos.

Exemplo de uso:

	idadeHumana := 5
	idadeCanina := cachorro.Idade(idadeHumana)
	fmt.Printf("A idade de %d anos humanos é equivalente a %d anos caninos.\n", idadeHumana, idadeCanina)
*/
func Idade(anosHumanos int) int {
	return anosHumanos * 7
}
