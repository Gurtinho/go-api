package main

import "fmt"

func main() {
	variavel := 10
	fmt.Println(variavel)
	fmt.Println(&variavel)

	ponteiroFalso := variavel
	fmt.Println(ponteiroFalso)

	ponteiro := &variavel
	*ponteiro = 20
	fmt.Println(*ponteiro)

	callback(&variavel)
	fmt.Println(variavel)
}

func callback(a *int) {
	*a = 200
}