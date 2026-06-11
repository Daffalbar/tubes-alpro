package main

import "fmt"

type Film struct {
	ID     int
	Judul  string
	Tahun  int
	Desk   string
	Rating float32
	Genre  string
}

var film []Film


var totalfilm int = 0

func main() {
	var n int

	for {
		fmt.Println("\n--- Selamat Datang di CineReview ---")
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
		default:
			fmt.Println("Pilihan tidak valid")
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
	fmt.Print("Rating Film (0.0 - 5.0): ")
	fmt.Scanln(&sementara.Rating)
	fmt.Print("Genre Film: ")
	fmt.Scanln(&sementara.Genre)

	film = append(film, sementara)
	totalfilm++
	fmt.Println("Film berhasil ditambahkan!")
}

func editfilm() {
	var n int
	fmt.Print("Masukkan ID Film Yang Ingin Diubah: ")
	fmt.Scanln(&n)
	for i := 0; i < len(film); i++ {
		if film[i].ID == n {
			fmt.Print("Judul Film: ")
			fmt.Scanln(&film[i].Judul)
			fmt.Print("Tahun Film: ")
			fmt.Scanln(&film[i].Tahun)
			fmt.Print("Deskripsi Film: ")
			fmt.Scanln(&film[i].Desk)
			fmt.Print("Rating Film (0.0 - 5.0): ")
			fmt.Scanln(&film[i].Rating)
			fmt.Print("Genre Film: ")
			fmt.Scanln(&film[i].Genre)
			fmt.Println("Film berhasil diubah!")
			return
		}
	}
	fmt.Println("Film tidak ditemukan")
}

func hapusfilm() {
	var n int
	fmt.Print("Masukkan ID Film Yang Ingin Dihapus: ")
	fmt.Scanln(&n)
	for i := 0; i < len(film); i++ {
		if film[i].ID == n {

			for j := i; j < len(film)-1; j++ {
				film[j] = film[j+1]
			}
			film = film[:len(film)-1]
			totalfilm--

			for j := i; j < len(film); j++ {
				film[j].ID = j + 1
			}
			fmt.Println("Film berhasil dihapus!")
			return
		}
	}
	fmt.Println("Film tidak ditemukan")
}

func cetakfilm(f Film) {
	fmt.Println("---")
	fmt.Println("ID     :", f.ID)
	fmt.Println("Judul  :", f.Judul)
	fmt.Println("Tahun  :", f.Tahun)
	fmt.Println("Desk   :", f.Desk)
	fmt.Printf("Rating : %.2f\n", f.Rating)
	fmt.Println("Genre  :", f.Genre)
	fmt.Println("---")
}

func lihatlistfilm() {
	if len(film) == 0 {
		fmt.Println("Belum ada data film")
		return
	}
	for i := 0; i < len(film); i++ {
		cetakfilm(film[i])
	}
}

func caridatafilm() {
	fmt.Println("1. Cari Berdasarkan Judul (Sequential Search)")
	fmt.Println("2. Cari Berdasarkan Genre (Binary Search)")

	var pilih int
	fmt.Print("Masukkan Pilihan: ")
	fmt.Scanln(&pilih)

	if pilih == 1 {
		var kata string

		fmt.Print("Masukkan Judul Film: ")
		fmt.Scanln(&kata)
		ketemu := false
		for i := 0; i < len(film); i++ {
			if film[i].Judul == kata {
				cetakfilm(film[i])
				ketemu = true
			}
		}
		if !ketemu {
			fmt.Println("Film tidak ditemukan")
		}

	} else if pilih == 2 {
		var kata string

		fmt.Print("Masukkan Genre Film: ")
		fmt.Scanln(&kata)

		for i := 1; i < len(film); i++ {

			key := film[i]
			j := i - 1

			for j >=0 && film[j].Genre > key.Genre {
				film[j+1] = film[j]
				j--
			}
			film[j+1] = key
		}

		lo := 0
		hi := len(film) - 1
		midfound:= -1
		for lo <= hi {
			mid := (lo + hi) /2
			if film[mid].Genre == kata {
				midfound = mid
				break

			} else if film[mid].Genre < kata {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}

		if midfound == -1 {
			fmt.Println("Film tidak ditemukan")
		} else { 
			left := midfound
			for left > 0 && film[left-1].Genre == kata {
				left--
			}
			for i := left; i < len(film) && film[i].Genre == kata; i++ {
				cetakfilm(film[i])
			}
		}

	} else {
		fmt.Println("Pilihan tidak valid.")
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
		for i := 0; i < n-1; i++ {
			maxIdx := i
			for j := i + 1; j < n; j++ {
				if film[j].Rating > film[maxIdx].Rating {
					maxIdx = j
				}
			}
			film[i], film[maxIdx] = film[maxIdx], film[i]
		}
		fmt.Println("Film diurutkan berdasarkan rating tertinggi ke terendah")
		lihatlistfilm()

	} else if pilih == 2 {
		n := len(film)

		for i := 1; i < n; i++ {

			key := film[i]
			j := i - 1
			for j >= 0 && film[j].Tahun < key.Tahun {
				film[j+1] = film[j]
				j--
			}
			film[j+1] = key
		}
		fmt.Println("Film diurutkan berdasarkan tahun terbaru ke terlama")
		lihatlistfilm()

	} else {
		fmt.Println("Pilihan tidak valid.")
		return
	}

	n := len(film)

	for i := 1; i < n; i++ {
		key := film[i]
		j := i - 1

		for j >= 0 && film[j].ID > key.ID {
			film[j+1] = film[j]
			j--
		}
		film[j+1] = key
	}
}

func tampilkanstatistik() {
	if len(film) == 0 {
		fmt.Println("Belum ada data film.")
		return
	}

	var total float32
	for i := 0; i < len(film); i++ {
		total += film[i].Rating
	}
	rerata := total / float32(len(film))

	fmt.Println("\n=== STATISTIK FILM ===")
	fmt.Println("Total Film      :", len(film))
	fmt.Printf("Rata-rata Rating: %.2f\n", rerata)
	fmt.Println("Jumlah Film per Genre:")

	for i := 0; i < len(film); i++ {

		genre := film[i].Genre
		sudahdihitung := false

		for j := 0; j < i; j++ {
			if film[j].Genre == genre {
				sudahdihitung = true
				break
			}
		}
		if !sudahdihitung {
			jumlah := 0
			for k := 0; k < len(film); k++ {
				if film[k].Genre == genre {
					jumlah++
				}
			}
			fmt.Printf("%s : %d film\n", genre, jumlah)
		}	

	}
	fmt.Println("======================")
}