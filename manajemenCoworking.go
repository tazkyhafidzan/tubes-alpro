package main

import "fmt"

const (
	NMAX         int = 1000
	maxFasilitas int = 10
)

type CoworkingSpace struct {
	No              int
	Nama            string
	Lokasi          string
	HargaSewa       int
	Fasilitas       [maxFasilitas]string
	JumlahFasilitas int
	Rating          float64
}

type Ulasan struct {
	No        int
	Nama      string
	IsiUlasan string
	Rating    float64
}

type tabCoworking [NMAX]CoworkingSpace
type tabUlasan [NMAX]Ulasan

var nomorCoworking = 1

func main() {
	var dataCoworking tabCoworking
	var nCoworking int
	var pilihan int

	for {
		fmt.Println("==========================================================")
		fmt.Println("=     APLIKASI MANAJEMEN DAN REVIEW CO-WORKING SPACE     =")
		fmt.Println("=                Created By Tazky & Ratna                =")
		fmt.Println("=               Algoritma Pemrograman 2026               =")
		fmt.Println("==========================================================")
		fmt.Println("=          1. Lihat Daftar Co-Working Space              =")
		fmt.Println("=          2. Tambah Daftar Co-Working Space             =")
		fmt.Println("=          3. Edit Daftar Co-Working Space               =")
		fmt.Println("=          4. Hapus Daftar Co-Working Space              =")
		fmt.Println("=          5. Keluar                                     =")
		fmt.Println("==========================================================")

		fmt.Print("Masukkan pilihan (1-5): ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			lihatDaftarCoworking(dataCoworking, nCoworking)

			if nCoworking != 0 {
				lihatDetailCoworking(dataCoworking, nCoworking)
			}
		} else if pilihan == 2 {
			tambahCoworking(&dataCoworking, &nCoworking)
		} else if pilihan == 3 {
			editCoworking(&dataCoworking, nCoworking)
		} else if pilihan == 4 {

		} else if pilihan == 5 {
			break
		} else {
			fmt.Println("\nPilihan tidak valid!\n")
		}
	}
}

func cariIndeks(A tabCoworking, n, x int) int {
	var idx = -1
	var i = 0

	for idx == -1 && i < n {
		if A[i].No == x {
			idx = i
		}

		i++
	}

	return idx
}

func lihatDaftarCoworking(A tabCoworking, n int) {
	if n == 0 {
		fmt.Println("\n=================================================================")
		fmt.Println("\nBelum ada data co-working space.\n")
		fmt.Println("=================================================================\n")
		return
	}

	fmt.Println("\n=================================================================")
	fmt.Printf("%-5s %-20s %-15s %-10s %-6s\n", "NO", "Nama", "Lokasi", "Harga", "Rating")
	fmt.Println("------------------------------------------------------------------")

	for i := 0; i < n; i++ {
		fmt.Printf("%-5d %-20s %-15s Rp%-10d %-6.1f\n",
			A[i].No, A[i].Nama, A[i].Lokasi, A[i].HargaSewa, A[i].Rating)

	}
	fmt.Println("=================================================================\n")
}

func lihatDetailCoworking(A tabCoworking, n int) {
	for {
		fmt.Print("\nMasukkan nomor Co-Working untuk melihat detail, ketik 0 untuk kembali: ")
		var nomor int
		fmt.Scan(&nomor)

		if nomor == 0 {
			break
		}

		var idx = cariIndeks(A, n, nomor)
		fmt.Println("\n=== DETAIL CO-WORKING SPACE ===")
		fmt.Printf("No        : %d\n", A[idx].No)
		fmt.Printf("Nama      : %s\n", A[idx].Nama)
		fmt.Printf("Lokasi    : %s\n", A[idx].Lokasi)
		fmt.Printf("Harga     : Rp%d\n", A[idx].HargaSewa)
		fmt.Printf("Rating    : %.1f\n", A[idx].Rating)

		fmt.Print("Fasilitas : ")
		if A[idx].JumlahFasilitas == 0 {
			fmt.Print("-")
		}

		for i := 0; i < A[idx].JumlahFasilitas; i++ {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Print(A[idx].Fasilitas[i])
		}
		fmt.Println("\n===============================\n")
	}
	fmt.Println()
}

func tambahCoworking(A *tabCoworking, n *int) {
	if *n >= NMAX {
		*n = NMAX
		fmt.Println("\n=================================================================")
		fmt.Println("\nData sudah penuh!\nTidak bisa menambah data!\n")
		fmt.Println("=================================================================")
		return
	}

	var temp CoworkingSpace

	temp.No = nomorCoworking
	nomorCoworking++

	fmt.Println("\n=================================================================")
	fmt.Print("Nama co-working space : ")
	fmt.Scan(&temp.Nama)
	fmt.Print("Lokasi                : ")
	fmt.Scan(&temp.Lokasi)
	fmt.Print("Harga sewa per hari   : ")
	fmt.Scan(&temp.HargaSewa)
	temp.Rating = 0

	fmt.Print("Masukkan jumlah fasilitas (maks 10): ")
	var nFasilitas int
	fmt.Scan(&nFasilitas)

	if nFasilitas > maxFasilitas {
		nFasilitas = maxFasilitas
	}

	for i := 0; i < nFasilitas; i++ {
		fmt.Scan(&temp.Fasilitas[i])
	}

	temp.JumlahFasilitas = nFasilitas

	A[*n] = temp
	*n++

	fmt.Printf("\nData Co-Working \"%s\" berhasil ditambahkan!\n\n", temp.Nama)
}

func editCoworking(A *tabCoworking, n int) {
	if n == 0 {
		fmt.Println("\n=================================================================")
		fmt.Println("\nBelum ada data co-working space.\n")
		fmt.Println("=================================================================\n")
		return
	}

	lihatDaftarCoworking(*A, n)

	fmt.Print("\nMasukkan nomor Co-Working yang ingin diedit, ketik 0 untuk kembali: ")
	var nomor int
	fmt.Scan(&nomor)

	if nomor == 0 {
		return
	}

	var idx = cariIndeks(*A, n, nomor)
	var temp = A[idx]

	fmt.Println("Ketik data baru untuk mengubah, ketik '-' jika tidak ingin mengubah")
	fmt.Printf("Nama [%s]: ", temp.Nama)
	var input string
	fmt.Scan(&input)
	if input != "-" {
		temp.Nama = input
	}

	fmt.Printf("Lokasi [%s]: ", temp.Lokasi)
	fmt.Scan(&input)
	if input != "-" {
		temp.Lokasi = input
	}

	fmt.Printf("Harga sewa [%d]: ", temp.HargaSewa)
	var harga int
	fmt.Scan(&harga)
	if harga > 0 {
		temp.HargaSewa = harga
	}

	A[idx] = temp

	fmt.Println("\nData berhasil diperbarui!")
	fmt.Println("=================================================================\n")
}

func hapusCoworking() {

}
