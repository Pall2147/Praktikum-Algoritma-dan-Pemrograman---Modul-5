package main

import "fmt"

func main() {
	var a,b,c int
	var saklar bool

	fmt.Scan(&a, &b, &c)
	saklar = (a == 1 && b == 1) || (a == 1 && c==1) || (b == 1 && c == 1)

	fmt.Print(saklar)
}