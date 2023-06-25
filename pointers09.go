package main

import "fmt"

type Livro struct {
	Titulo string
	Autor  string
	Preco  float64
}

func desconto(l *Livro) {

	l.Preco -= l.Preco * 0.1

}

func main() {

	l := Livro{
		Titulo: "123456",
		Autor:  "1234545677686",
		Preco:  500,
	}

	desconto(&l)

	fmt.Println(l.Preco)

}
