package main

import "fmt"

func main() {
	const (
		kuis float64 = 0.15
		tugas float64 = 0.15			
		praktikum float64 = 0.20
		UTS float64 = 0.25
		UAS float64 = 0.25
	)

	type mahasiswa struct {
		NIM  string
		Nama  string
		NilaiTugas float64
		NilaiKuis float64
		NilaiPraktikum float64 
		NilaiUTS float64 
		NilaiUAS float64 
		NilaiAkhir float64
	}
	var Datamahasiswa mahasiswa 
	
	fmt.Scan(&Datamahasiswa.Nama)
	fmt.Scan(&Datamahasiswa.NIM)	
	fmt.Scan(&Datamahasiswa.NilaiTugas)
	fmt.Scan(&Datamahasiswa.NilaiKuis)
	fmt.Scan(&Datamahasiswa.NilaiPraktikum)
	fmt.Scan(&Datamahasiswa.NilaiUTS)
	fmt.Scan(&Datamahasiswa.NilaiUAS)

	Datamahasiswa.NilaiAkhir = Datamahasiswa.NilaiTugas * tugas +
	 Datamahasiswa.NilaiKuis * kuis +
	   Datamahasiswa.NilaiPraktikum * praktikum +
	    Datamahasiswa.NilaiUTS * UTS +
		 Datamahasiswa.NilaiUAS * UAS 

	fmt.Println("Data Mahasiswa")
	fmt.Println("Nama", Datamahasiswa.Nama)
	fmt.Println("NIM", Datamahasiswa.NIM)
	fmt.Println("Nilai Tugas", Datamahasiswa.NilaiTugas)
	fmt.Println("Nilai Kuis", Datamahasiswa.NilaiKuis)
	fmt.Println("Nilai Praktikum", Datamahasiswa.NilaiPraktikum)
	fmt.Println("Nilai UTS", Datamahasiswa.NilaiUTS)
	fmt.Println("Nilai UAS", Datamahasiswa.NilaiUAS)
	fmt.Println("Nilai Akhir", Datamahasiswa.NilaiAkhir)
}

	