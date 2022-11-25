package main

import "fmt"

func main() {
	var c int
	var k = 373
	c = k - 273

	mensagem := fmt.Sprintf("O ponto de ebulição da aguá é de %d celcius", c)

	fmt.Println(mensagem)
}
