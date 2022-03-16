package main

import "fmt"

func main() {
	//Uma função pode retornar mais de um valor
	// A construção := cria a variavel caso ela não exista, já definindo seu tipo

	ok, err := say("Hello word")
	if err != nil {
		panic(err.Error())
	}
	switch ok {
	case true:
		fmt.Println("Deu certo")
	default:
		fmt.Println("Deu errado")
	}
}

//As funções devem ceclarar o tipo de cada variável que recebe ou que retorna

func say(what string) (bool, error) {
	if what == "" {
		return false, fmt.Errorf("Empty string")
	}
	fmt.Println(what)
	return true, nil
}
