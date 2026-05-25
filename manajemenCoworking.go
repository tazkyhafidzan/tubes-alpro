package main

import (
	"fmt"
	"strings"
)

const (
	CoworkingMax = 1000
	UlasanMax    = 1000
	FasilitasMax = 10
)

type CoworkingSpace struct {
	ID           string
	Nama         string
	Lokasi       string
	Fasilitas    [FasilitasMax]string
	JmlFasilitas int
	HargaSewa    int
	Rating       float64
	Ulasan       [UlasanMax]Ulasan
	JmlUlasan    int
}

type Ulasan struct {
	ID       string
	Penulis  string
	Komentar string
	Rating   float64
}

type tabCoworking [CoworkingMax]CoworkingSpace
type tabUlasan [UlasanMax]Ulasan

func main() {
	mainMenu()
}

func mainMenu() {
	var dataCoworking tabCoworking
	var nData int

	for {
		fmt.Println("----------------------------------------------------------")
		fmt.Println("|     APLIKASI MANAJEMEN DAN REVIEW CO-WORKING SPACE     |")
		fmt.Println("|                Created By Tazky & Ratna                |")
		fmt.Println("|               Algoritma Pemrograman 2026               |")
		fmt.Println("----------------------------------------------------------")
		fmt.Println("|          1. Lihat Daftar Co-Working Space              |")
		fmt.Println("|          2. Tambah Daftar Co-Working Space             |")
		fmt.Println("|          3. Edit Daftar Co-Working Space               |")
		fmt.Println("|          4. Hapus Daftar Co-Working Space              |")
		fmt.Println("|          5. Cari Co-Working Space                      |")
		fmt.Println("|          6. Keluar                                     |")
		fmt.Println("----------------------------------------------------------")

		fmt.Print("Masukkan pilihan (1-5): ")
		var pilihan int
		fmt.Scan(&pilihan)
		fmt.Println()

		if pilihan == 1 {
			cetakCoworking(dataCoworking, nData)

			for {
				var x string
				fmt.Print("Masukkan ID Co-Working Space untuk melihat detail, ketik 0 untuk kembali: ")
				fmt.Scan(&x)
				fmt.Println()

				indeks := searchByID(dataCoworking, nData, x)

				if x == "0" || indeks == -1 {
					break
				}

				detailCoworking(dataCoworking, indeks)
			}
		} else if pilihan == 2 {
			tambahCoworking(&dataCoworking, &nData)
		} else if pilihan == 3 {
			editCoworking(&dataCoworking, nData)
		} else if pilihan == 4 {
			hapusCoworking(&dataCoworking, &nData)
		} else if pilihan == 5 {
			cariMenu(dataCoworking, nData)
		} else if pilihan == 6 {
			break
		} else {
			fmt.Println("Pilihan tidak valid!")
			continue
		}
	}
}

func ulasanMenu() {

}

func cariMenu(A tabCoworking, n int) {
	for {
		fmt.Println("----------------------------------------------------------")
		fmt.Println("|     APLIKASI MANAJEMEN DAN REVIEW CO-WORKING SPACE     |")
		fmt.Println("|                Created By Tazky & Ratna                |")
		fmt.Println("|               Algoritma Pemrograman 2026               |")
		fmt.Println("----------------------------------------------------------")
		fmt.Println("|                   Cari berdasarkan:                    |")
		fmt.Println("|                       1. Nama                          |")
		fmt.Println("|                       2. Lokasi                        |")
		fmt.Println("|                       0. Kembali                        |")
		fmt.Println("----------------------------------------------------------")

		fmt.Print("Masukkan pilihan (1-2) atau ketik 0 untuk kembali: ")
		var pilihan int
		fmt.Scan(&pilihan)
		fmt.Println()

		cetakCoworking(A, n)

		if pilihan == 1 {
			var x string
			fmt.Print("Masukkan Nama Co-Working Space: ")
			fmt.Scan(&x)

			indeks := searchByNama(A, n, strings.ToLower(x))

			if indeks == -1 {
				break
			}

			detailCoworking(A, indeks)
		} else if pilihan == 0 {
			break
		} else {
			fmt.Println("Pilihan tidak valid!")
			break
		}
	}
}

func detailCoworking(A tabCoworking, indeks int) {
	fmt.Println("---------------------------------------")
	fmt.Println("        DETAIL CO-WORKING SPACE        ")
	fmt.Println("---------------------------------------")
	fmt.Printf("ID        : %s\n", A[indeks].ID)
	fmt.Printf("Nama      : %s\n", A[indeks].Nama)
	fmt.Printf("Lokasi    : %s\n", A[indeks].Lokasi)
	fmt.Printf("Harga     : Rp%d\n", A[indeks].HargaSewa)
	fmt.Printf("Rating    : %.1f\n", A[indeks].Rating)

	fmt.Print("Fasilitas : ")
	if A[indeks].JmlFasilitas == 0 {
		fmt.Print("-")
	}

	for i := 0; i < A[indeks].JmlFasilitas; i++ {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(A[indeks].Fasilitas[i])
	}
	fmt.Println("\n---------------------------------------\n")
}

func cetakCoworking(A tabCoworking, n int) {
	fmt.Printf("%-5s | %-5s | %-20s | %-15s | %-10s | %-6s\n",
		"No", "ID", "Nama", "Lokasi", "Harga Sewa", "Rating")
	fmt.Println("----------------------------------------------------------------------------")

	for i := 0; i < n; i++ {
		fmt.Printf("%-5d | %-5s | %-20s | %-15s | Rp%-8d | %-6.1f\n",
			i+1, A[i].ID, A[i].Nama, A[i].Lokasi, A[i].HargaSewa, A[i].Rating)
	}
	fmt.Println()
}

func tambahCoworking(A *tabCoworking, n *int) {
	if *n > CoworkingMax {
		*n = CoworkingMax
	}

	for {
		var temp CoworkingSpace

		fmt.Print("ID Co-Working space : ")
		fmt.Scan(&temp.ID)
		fmt.Print("Nama Co-Working space : ")
		fmt.Scan(&temp.Nama)
		fmt.Print("Lokasi                : ")
		fmt.Scan(&temp.Lokasi)
		fmt.Print("Harga sewa per hari   : ")
		fmt.Scan(&temp.HargaSewa)
		temp.Rating = 0

		/* ------------------ Fasilitas ------------------ */
		fmt.Print("Masukkan jumlah fasilitas (maks 10): ")
		var nFasilitas int
		fmt.Scan(&nFasilitas)

		if nFasilitas > FasilitasMax {
			nFasilitas = FasilitasMax
		}

		for i := 0; i < nFasilitas; i++ {
			fmt.Printf("  Fasilitas %d: ", i+1)
			fmt.Scan(&temp.Fasilitas[i])
		}

		temp.JmlFasilitas = nFasilitas
		/* ------------------ Fasilitas ------------------ */

		A[*n] = temp
		*n++

		fmt.Printf("\nData Co-Working \"%s\" berhasil ditambahkan!\n\n", temp.Nama)

		/* ============================================================================================================ */

		fmt.Println("Apakah masih ada data yang ingin anda tambah?")
		fmt.Println("1. Ya")
		fmt.Println("2. Tidak")
		fmt.Print("Masukkan pilihan : ")
		var pilihan int
		fmt.Scan(&pilihan)
		fmt.Println()

		if pilihan == 1 {
			continue
		} else if pilihan == 2 {
			break
		} else {
			fmt.Println("Pilihan tidak valid!")
			break
		}
	}
}

func editCoworking(A *tabCoworking, n int) {
	for {
		cetakCoworking(*A, n)

		var x string
		fmt.Print("Masukkan ID Co-Working Space yang ingin diedit : ")
		fmt.Scan(&x)

		indeks := searchByID(*A, n, x)

		idLama := A[indeks].ID
		ulasanLama := A[indeks].Ulasan
		jmlUlasanLama := A[indeks].JmlUlasan
		ratingLama := A[indeks].Rating

		fmt.Println("Masukkan data baru (semua field harus diisi ulang):")

		fmt.Print("Nama baru        : ")
		fmt.Scan(&A[indeks].Nama)

		fmt.Print("Lokasi baru      : ")
		fmt.Scan(&A[indeks].Lokasi)

		fmt.Print("Harga sewa baru  : ")
		fmt.Scan(&A[indeks].HargaSewa)

		/* ------------------ Fasilitas ------------------ */
		fmt.Print("Masukkan jumlah fasilitas baru (maks 10): ")
		var nFasilitas int
		fmt.Scan(&nFasilitas)

		if nFasilitas > FasilitasMax {
			nFasilitas = FasilitasMax
		}

		for i := 0; i < FasilitasMax; i++ {
			A[indeks].Fasilitas[i] = ""
		}

		for i := 0; i < nFasilitas; i++ {
			fmt.Printf("  Fasilitas %d: ", i+1)
			fmt.Scan(&A[indeks].Fasilitas[i])
		}

		A[indeks].JmlFasilitas = nFasilitas
		/* ------------------ Fasilitas ------------------ */

		A[indeks].ID = idLama
		A[indeks].Ulasan = ulasanLama
		A[indeks].JmlUlasan = jmlUlasanLama
		A[indeks].Rating = ratingLama

		fmt.Println("\n✅ Data berhasil diperbarui!\n")

		/* ============================================================================================================ */

		fmt.Println("Apakah masih ada data yang ingin anda edit?")
		fmt.Println("1. Ya")
		fmt.Println("2. Tidak")
		fmt.Print("Masukkan pilihan : ")
		var pilihan int
		fmt.Scan(&pilihan)
		fmt.Println()

		if pilihan == 1 {
			continue
		} else if pilihan == 2 {
			break
		} else {
			fmt.Println("Pilihan tidak valid!")
			break
		}
	}
}

func hapusCoworking(A *tabCoworking, n *int) {
	for {
		cetakCoworking(*A, *n)

		var x string
		fmt.Print("Masukkan ID Co-Working Space yang ingin dihapus : ")
		fmt.Scan(&x)

		indeks := searchByID(*A, *n, x)

		if indeks != -1 {
			for i := indeks; i < *n-1; i++ {
				A[i] = A[i+1]
			}
			*n--
		} else {
			fmt.Println("Data tidak ditemukan")
		}
		fmt.Printf("\n✅ Co-working space \"%s\" berhasil dihapus!\n\n", A[indeks].Nama)

		/* ============================================================================================================ */

		fmt.Println("Apakah masih ada data yang ingin anda hapus?")
		fmt.Println("1. Ya")
		fmt.Println("2. Tidak")
		fmt.Print("Masukkan pilihan : ")
		var pilihan int
		fmt.Scan(&pilihan)
		fmt.Println()

		if pilihan == 1 {
			continue
		} else if pilihan == 2 {
			break
		} else {
			fmt.Println("Pilihan tidak valid!")
			break
		}
	}
}

func searchByID(A tabCoworking, n int, x string) int {
	var idx, i int

	idx = -1
	i = 0

	for idx == -1 && i < n {
		if A[i].ID == x {
			idx = i
		}

		i++
	}

	return idx
}

func searchByNama(A tabCoworking, n int, x string) int {
	var idx, i int

	idx = -1
	i = 0

	for idx == -1 && i < n {
		if strings.ToLower(A[i].Nama) == x {
			idx = i
		}

		i++
	}

	return idx
}

//
//func seqSearchByLokasi() int {
//
//}
//
//func binSearchByNama() int {
//
//}
//
//func binSearchByLokasi() int {
//
//}
