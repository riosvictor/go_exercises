package main

import (
	"fmt"
	"runtime"
)

/*
Crie um programa que demonstra seu OS e ARCH.
Rode-o com os seguintes comandos:
go run
	> go run exerc-9.6/main.go
go build
	> go build exerc-9.6/main.go
go install
  > go mod init exerc-9.6
	> go install .
	> ~/go/bin/exerc-9.6
*/

func main() {
	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("ARCH:", runtime.GOARCH)
}
