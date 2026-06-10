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
		if film[i].ID == n {
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
			fmt.Println("Film berhasil diubah!")
			return
		}
	}
	fmt.Println("Film tidak ditemukan.")
}

func hapusfilm() {
	var n int
	fmt.Print("Masukkan ID Film Yang Ingin Di Hapus: ")
	fmt.Scanln(&n)
	for i:=0;i<len(film);i++{
		if film[i].ID == n {
			film = append(film[:i], film[i+1:]...)
			totalfilm--
			fmt.Println("Film berhasil dihapus!")
			return
		}
	}
	fmt.Println("Film tidak ditemukan.")
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
	fmt.Println("1. Cari Berdasarkan Judul (Sequential Search)")
	fmt.Println("2. Cari Berdasarkan Judul (Binary Search)")
	var pilih int
	fmt.Print("Masukkan Pilihan: ")
	fmt.Scanln(&pilih)

	var kata string

	if pilih == 1 {
		fmt.Print("Masukkan Judul Film: ")
		fmt.Scanln(&kata)
		ketemu := false
		for i:=0;i<len(film);i++{
			if film[i].Judul == kata {
				fmt.Println("---")
				fmt.Println("ID:", film[i].ID)
				fmt.Println("Judul:", film[i].Judul)
				fmt.Println("Tahun:", film[i].Tahun)
				fmt.Println("Deskripsi:", film[i].Desk)
				fmt.Println("Rating:", film[i].rating)
				fmt.Println("Genre:", film[i].Genre)
				fmt.Println("---")
				ketemu = true
			}
		}
		if !ketemu {
			fmt.Println("Film tidak ditemukan.")
		}

	} else if pilih == 2 {
		fmt.Print("Masukkan Genre Film: ")
		fmt.Scanln(&kata)
		ketemu := false
		for i:=0;i<len(film);i++{
			if film[i].Genre == kata {
				fmt.Println("---")
				fmt.Println("ID:", film[i].ID)
				fmt.Println("Judul:", film[i].Judul)
				fmt.Println("Tahun:", film[i].Tahun)
				fmt.Println("Deskripsi:", film[i].Desk)
				fmt.Println("Rating:", film[i].rating)
				fmt.Println("Genre:", film[i].Genre)
				fmt.Println("---")
				ketemu = true
			}
		}
		if !ketemu {
			fmt.Println("Film tidak ditemukan.")
		}

	} else if pilih == 3 {
		fmt.Print("Masukkan Judul Film: ")
		fmt.Scanln(&kata)

		sorted := make([]Film, len(film))
		copy(sorted, film)
		for i:=1;i<len(sorted);i++{
			key := sorted[i]
			j := i - 1
			for j>=0 && sorted[j].Judul > key.Judul {
				sorted[j+1] = sorted[j]
				j--
			}
			sorted[j+1] = key
		}

		lo := 0
		hi := len(sorted) - 1
		ketemu := false
		for lo <= hi {
			mid := (lo + hi) / 2
			if sorted[mid].Judul == kata {
				fmt.Println("---")
				fmt.Println("ID:", sorted[mid].ID)
				fmt.Println("Judul:", sorted[mid].Judul)
				fmt.Println("Tahun:", sorted[mid].Tahun)
				fmt.Println("Deskripsi:", sorted[mid].Desk)
				fmt.Println("Rating:", sorted[mid].rating)
				fmt.Println("Genre:", sorted[mid].Genre)
				fmt.Println("---")
				ketemu = true
				break
			} else if sorted[mid].Judul < kata {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
		if !ketemu {
			fmt.Println("Film tidak ditemukan.")
		}
	}
}

func urutkandatafilm() {
	fmt.Println("1. Urutkan Berdasarkan Rating (Selection Sort)")
	fmt.Println("2. Urutkan Berdasarkan Tahun (Insertion Sort)")
	var pilih int
	fmt.Print("Masukkan Pilihan: ")
	fmt.Scanln(&pilih)

	if pilih == 1 {
		n := len(film)
		for i:=0;i<n-1;i++{
			maxIdx := i
			for j:=i+1;j<n;j++{
				if film[j].rating > film[maxIdx].rating {
					maxIdx = j
				}
			}
			film[i], film[maxIdx] = film[maxIdx], film[i]
		}
		fmt.Println("Film diurutkan berdasarkan rating tertinggi ke terendah.")
		lihatlistfilm()

	} else if pilih == 2 {
		n := len(film)
		for i:=1;i<n;i++{
			key := film[i]
			j := i - 1
			for j>=0 && film[j].Tahun < key.Tahun {
				film[j+1] = film[j]
				j--
			}
			film[j+1] = key
		}
		fmt.Println("Film diurutkan berdasarkan tahun terbaru ke terlama.")
		lihatlistfilm()
	}
}

func tampilkanstatistik() {
	if len(film) == 0 {
		fmt.Println("Belum ada data film.")
		return
	}

	var total float32
	for i:=0;i<len(film);i++{
		total += film[i].rating
	}
	rataRata := total / float32(len(film))

	fmt.Println("=== STATISTIK FILM ===")
	fmt.Println("Total Film      :", len(film))
	fmt.Printf("Rata-rata Rating: %.2f\n", rataRata)
	fmt.Println("Jumlah Film per Genre:")
	for i:=0;i<len(film);i++{
		genre := film[i].Genre
		sudahDihitung := false
		for j:=0;j<i;j++{
			if film[j].Genre == genre {
				sudahDihitung = true
				break
			}
		}
		if !sudahDihitung {
			jumlah := 0
			for k:=0;k<len(film);k++{
				if film[k].Genre == genre {
					jumlah++
				}
			}
			fmt.Println(" ", genre, ":", jumlah, "film")
		}
	}
	fmt.Println("======================")
}