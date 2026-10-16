package main

import "fmt"

func main() {
	const (
		diskonpertama float64 = 0.10
		diskonkedua float64 = 0.20
	)
	var hargaAwal,hargaAkhir,setelahDiskon1 float64
	fmt.Scan(&hargaAwal)

	setelahDiskon1 = hargaAwal * (1- diskonpertama)
	hargaAkhir = setelahDiskon1 * (1- diskonkedua)

	fmt.Printf( "%.2f\n" , hargaAkhir)
}