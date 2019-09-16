package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	fmt.Println(1, 2, 100)
	fmt.Println("Literal Inteiro é", reflect.TypeOf(32000))

	// sem sinal
	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	//com sinal
	i1 := math.MaxInt64
	fmt.Println("O valor maximo do Int é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a'
	fmt.Println("O rune é", reflect.TypeOf(i2))
	fmt.Println(i2)

	// numeros reais
	var x = float32(49.99)
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipo do literal 49.99 é", reflect.TypeOf(49.99))

	//boolean
	bo := true
	fmt.Println("O tipo de bo é:", reflect.TypeOf(bo))
	fmt.Println(!bo)

	//string
	s1 := "Olá meu nome é Marcus"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string é", len(s1))

	//string com multiplas linhas
	s2 := `Olá
	meu 
	nome 
	é 
	Marcus`
	fmt.Println("O tamanho da string é", len(s2))

	// Char
	char := 'a'
	fmt.Println("O tipo de char é", reflect.TypeOf(char))
	fmt.Println(char)

}
