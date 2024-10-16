package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// saudacao()

	//exemplo de função que retorna multiplos valores
	// nem sempre vou querer utilizar os multiplos valores que são retornados, então
	// para lidar com isso, basta usar o operador do proprio go
	// o _ -> operador em branco -> meio que é usado pra ignorar o valor que vc não quer usar
	// _, idade := devolveNomeEIdade()
	// fmt.Println(idade, "anos")

	// algumas questões aqui
	// não existe o WHILE no golang (sim, estranho)
	// quando passamos o for sem parametro, o golang entende que será infinito
	// você precisa parar o programa para ele parar de rodar
	for {
		fmt.Println("1- Iniciar monitoramento")
		fmt.Println("2- Exibir logs")
		fmt.Println("0- Sair do programa")

		comando := leComando()

		//no go se o if é um pouco diferente (além da sintaxe)
		//ele só aceita expressões que retornam true ou false, nada de tentar verificar se o cara sozinho existe
		// ex if comando {} -> isso em outra linguagem estaria vendo se a variavel comando existe
		// if comando == 1 {
		// 	fmt.Println("Monitorando...")
		// } else if comando == 2 {
		// 	fmt.Println("Exibindo logs...")
		// } else if comando == 0 {
		// 	fmt.Println("Saindo do programa...")
		// } else {
		// 	fmt.Println("Não reconheço este comando!")
		// }

		//usando switch
		//o break no go é implicito, por isso não é necessario adicionar

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs...")
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Não reconheço este comando!")
			os.Exit(-1)
		}
	}
}

// quando vou devolver mais de uma coisa, preciso envolver em parenteses os tipos
func devolveNomeEIdade() (string, int) {
	nome := "Elquiane"
	idade := 23
	return nome, idade
}

func saudacao() {
	//em go, você pode usar esse := como operador de atribuição para
	//inferir o tipo e dizer implicitamente que isso é uma variável
	//muito usado (declaração curta de variaveis)
	nome := "Elquiane"
	var versao float32 = 1.1

	fmt.Println("Hello world,", nome)
	fmt.Println("Esse programa está na versão: ", versao)
	// fmt.Println("O tipo da variavel nome é:", reflect.TypeOf(nome))
}

func leComando() int {
	var comando int
	//captura o input do usuario
	//também diz que variavel vai armazenar o valor que recebe e seu tipo
	// %d -> representa um inteiro -> é o modificador
	//& -> endereço da variavel comando -> ponteiro
	// fmt.Scanf("%d", &comando)
	//função da qual ele já identifica o tipo da variavel
	fmt.Scan(&comando)

	// fmt.Println("O ponteiro da minha variavel comando é:", &comando)
	fmt.Println("O comando escolhido foi:", comando)

	return comando
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	site := "https://httpbin.org/status/404"
	// site := "https://www.alura.com.br"

	//a função get do pacote http retorna mais de um valor
	//
	response, _ := http.Get(site)

	if response.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", response.StatusCode)
	}

	// fmt.Println(response)
	// fmt.Println(error)

}
