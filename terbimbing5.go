package main

import "fmt"

func main() {
	type Buku struct {
		judul string
		genre string
		rilis string
		penilaian string
}
	var buku1,buku2 Buku
	fmt.Scan(&buku1.judul, &buku2.judul)
	fmt.Scan(&buku1.genre, &buku2.genre)
	fmt.Scan(&buku1.rilis, &buku2.rilis)
	fmt.Scan(&buku1.penilaian, &buku2.penilaian)

	fmt.Println("--- BUKUPERTAMA ---")
	fmt.Println(buku1.judul)
	fmt.Println(buku1.genre)
	fmt.Println(buku1.rilis)
	fmt.Println(buku1.penilaian)
	fmt.Println("--- BUKUKEDUA---")
	fmt.Println(buku2.judul)
	fmt.Println(buku2.genre)
	fmt.Println(buku2.rilis)
	fmt.Println(buku2.penilaian)
	fmt.Println("--------------")
}


