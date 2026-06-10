package main

import "fmt"

type Film struct{
	ID int
	Judul string
	Tahun int
	Desk string
	rating float32
	Genre string
}

var film []Film
var totalfilm int = 0
var panjang = len(film) + 1

func main() {
	var n int

	for {

	fmt.Println("--- Selamat Datang di CineReview ---")
	fmt.Println("1. Tambah Film")
	fmt.Println("2. Edit Film")
	fmt.Println("3. Hapus Film")
	fmt.Println("4. Lihat List Film")
	fmt.Println("5. Cari Data Film")
	fmt.Println("6. Urutkan Data Film")
	fmt.Println("7. Tampilkan Statistik Film")
	fmt.Println("0. Keluar")
	
	fmt.Print("Masukkan Pilihan: ")
	fmt.Scanln(&n)

	switch n {
	case 1:
		tambahfilm()
	case 2:
		editfilm()
	case 3:
		hapusfilm()
	case 4:
		lihatlistfilm()
	case 5:
		caridatafilm()
	case 6:
		urutkandatafilm()
	case 7:
		tampilkanstatistik()
	case 0:
		fmt.Println("--- Program Selesai ---")
		return
	}

	}
}

func tambahfilm() {
	var sementara Film
	sementara.ID = totalfilm + 1
	fmt.Print("Judul Film: ")
	fmt.Scanln(&sementara.Judul)
	fmt.Print("Tahun Film: ")
	fmt.Scanln(&sementara.Tahun)
	fmt.Print("Deskripsi Film: ")
	fmt.Scanln(&sementara.Desk)
	fmt.Print("Rating Film: ")
	fmt.Scanln(&sementara.rating)
	fmt.Print("Genre Film: ")
	fmt.Scanln(&sementara.Genre)

	film = append(film, sementara)
	totalfilm++
}

func editfilm() {
	var n int
	fmt.Print("Masukkan ID Film Yang Ingin DiUbah: ")
	fmt.Scanln(&n)
	for i:=0;i<len(film);i++{
		if i == n {
			fmt.Print("Judul Film: ")
			fmt.Scanln(&film[i].Judul)
			fmt.Print("Tahun Film: ")
			fmt.Scanln(&film[i].Tahun)
			fmt.Print("Deskripsi Film: ")
			fmt.Scanln(&film[i].Desk)
			fmt.Print("Rating Film: ")
			fmt.Scanln(&film[i].rating)
			fmt.Print("Genre Film: ")
			fmt.Scanln(&film[i].Genre)
		}
	}

}

func hapusfilm() {
	var n int
	fmt.Print("Masukkan ID Film Yang Ingin Di Hapus: ")
	fmt.Scanln(&n)

}

func lihatlistfilm() {
	for i:=0;i<len(film);i++{
		fmt.Println("---")
		fmt.Println("ID:", film[i].ID)
		fmt.Println("Judul:", film[i].Judul)
		fmt.Println("Tahun:", film[i].Tahun)
		fmt.Println("Deskripsi:", film[i].Desk)
		fmt.Println("Rating:", film[i].rating)
		fmt.Println("Genre:", film[i].Genre)
		fmt.Println("---\n")
	}

}

func caridatafilm() {

}

func urutkandatafilm() {

}

func tampilkanstatistik() {

}