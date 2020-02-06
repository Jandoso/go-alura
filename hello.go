package main 

import "fmt"
// import "reflect"
import "os"

func main() {

	exibeIntroducao()
	exibeMenu()
	comando := leComando()

	// Com if:

/*	if comando == 1 {
		fmt.Println("Monitorando...")
	} else if comando == 2 {
		fmt.Println("Exibindo Logs...")
	} else if comando == 0 {
		fmt.Println("Saindo do programa...")
	} else {
		fmt.Println("Não conheço este comando")
	} */

	// Com switch case

	switch comando {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo Logs...")
	case 0: 
		fmt.Println("Saindo do programa")
		os.Exit(0) // 0 -> saiu com sucesso
	default: 
		fmt.Println("Não conheço este comando")
		os.Exit(-1) // -1 -> saiu com algum problema
	}

/*	fmt.Println("O tipo da variável nome é", reflect.TypeOf(nome))
	
	fmt.Println("O tipo da variável versao é", reflect.TypeOf(versao)) */
}

func exibeIntroducao() {
// Atribuidor de variáveis curtas := (ele também infere o tipo da variável)
nome := "Caroline"
versao := 1.1
fmt.Println("Olá, srta.", nome)
fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func leComando() int{
	var comandoLido int
	
	//fmt.Scanf("%d", &comando)
	// Não precisamos mais especificar o tipo de modificador da variável
	fmt.Scan(&comandoLido)

//	fmt.Println("O endereço da minha variável comando é", &comando)

	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
}

