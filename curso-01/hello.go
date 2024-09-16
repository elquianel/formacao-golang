package main

import (
	"fmt"
)

func main() {
	//em go, você pode usar esse := como operador de atribuição para
	//inferir o tipo e dizer implicitamente que isso é uma variável
	//muito usado (declaração curta de variaveis)
	nome := "Elquiane"
	var versao float32 = 1.1

	fmt.Println("Hello world,", nome)
	fmt.Println("Esse programa está na versão: ", versao)

	// fmt.Println("O tipo da variavel nome é:", reflect.TypeOf(nome))
	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- Sair do programa")

	var comando int
	//captura o input do usuario
	//também diz que variavel vai armazenar o valor que recebe e seu tipo
	// %d -> representa um inteiro -> é o modificador
	//& -> endereço da variavel comando -> ponteiro
	fmt.Scanf("%d", &comando)
	//função da qual ele já identifica o tipo da variavel
	// fmt.Scan(&comando)

	fmt.Println("O ponteiro da minha variavel comando é:", &comando)
	fmt.Println("O comando escolhido foi:", comando)
}
