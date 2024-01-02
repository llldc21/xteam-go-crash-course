package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")

	// -- Variables --

	var (
		index               int
		noWord              string
		didWeLearnSomething bool = true
	)

	noWord = "Nope"
	someWord := "Yup"

	fmt.Println(index)
	fmt.Println(noWord)
	fmt.Println(didWeLearnSomething)
	fmt.Println(someWord)

	// -- Pointers --

	var pointer *int
	fmt.Println(pointer) // <nil>
	i := 4
	pointer = &i          // & é utilizado para pegar o endereço de memória da variável, nesse caso i que equivale a 4
	fmt.Println(pointer)  // 0x1400000e110 - endereço de memória
	fmt.Println(*pointer) // 4 - valor da variável que está no endereço de memória

	i = 42

	fmt.Println(*pointer) // 42 - alterando o valor da variável que está no endereço de memória
	fmt.Println(pointer)  // 0x1400000e120

}
