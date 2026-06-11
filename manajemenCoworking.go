package main

import (
	"fmt"
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
	ID       int
	Penulis  string
	Komentar string
	Rating   float64
}

type tabCoworking [CoworkingMax]CoworkingSpace

func main() {
	// Data Dummy
	var dataCoworking tabCoworking
	var nData int

	nData = 8

	dataCoworking[0] = CoworkingSpace{
		ID:           "CS001",
		Nama:         "Ruang Kreatif",
		Lokasi:       "Jakarta",
		Fasilitas:    [FasilitasMax]string{"WiFi", "Kopi", "AC"},
		JmlFasilitas: 3,
		HargaSewa:    150000,
		Rating:       4.5,
		JmlUlasan:    2,
	}
	dataCoworking[0].Ulasan[0] = Ulasan{ID: 1, Penulis: "Budi", Komentar: "Tempat nyaman", Rating: 4.5}
	dataCoworking[0].Ulasan[1] = Ulasan{ID: 2, Penulis: "Siti", Komentar: "Internet cepat", Rating: 4.5}

	dataCoworking[1] = CoworkingSpace{
		ID:           "CS002",
		Nama:         "Focus Hub",
		Lokasi:       "Bandung",
		Fasilitas:    [FasilitasMax]string{"WiFi", "Printer", "AC", "Whiteboard"},
		JmlFasilitas: 4,
		HargaSewa:    120000,
		Rating:       4.8,
		JmlUlasan:    1,
	}
	dataCoworking[1].Ulasan[0] = Ulasan{ID: 1, Penulis: "Andi", Komentar: "Tenang sekali", Rating: 4.8}

	dataCoworking[2] = CoworkingSpace{
		ID:           "CS003",
		Nama:         "Sinergi Space",
		Lokasi:       "Yogyakarta",
		Fasilitas:    [FasilitasMax]string{"WiFi", "Kopi", "Meeting Room"},
		JmlFasilitas: 3,
		HargaSewa:    90000,
		Rating:       4.2,
		JmlUlasan:    1,
	}
	dataCoworking[2].Ulasan[0] = Ulasan{ID: 1, Penulis: "Rian", Komentar: "Murah dan strategis", Rating: 4.2}

	dataCoworking[3] = CoworkingSpace{
		ID:           "CS004",
		Nama:         "Workify",
		Lokasi:       "Surabaya",
		Fasilitas:    [FasilitasMax]string{"WiFi", "AC", "Loker"},
		JmlFasilitas: 3,
		HargaSewa:    130000,
		Rating:       4.0,
		JmlUlasan:    1,
	}
	dataCoworking[3].Ulasan[0] = Ulasan{ID: 1, Penulis: "Dewi", Komentar: "Standar tapi oke", Rating: 4.0}

	dataCoworking[4] = CoworkingSpace{
		ID:           "CS005",
		Nama:         "Nexa Lab",
		Lokasi:       "Semarang",
		Fasilitas:    [FasilitasMax]string{"WiFi", "Projector", "AC"},
		JmlFasilitas: 3,
		HargaSewa:    110000,
		Rating:       4.6,
		JmlUlasan:    1,
	}
	dataCoworking[4].Ulasan[0] = Ulasan{ID: 1, Penulis: "Eko", Komentar: "Fasilitas lengkap", Rating: 4.6}

	dataCoworking[5] = CoworkingSpace{
		ID:           "CS006",
		Nama:         "Co-Zone",
		Lokasi:       "Bandung",
		Fasilitas:    [FasilitasMax]string{"WiFi", "Kopi", "Free Parking"},
		JmlFasilitas: 3,
		HargaSewa:    100000,
		Rating:       4.3,
		JmlUlasan:    1,
	}
	dataCoworking[5].Ulasan[0] = Ulasan{ID: 1, Penulis: "Mega", Komentar: "Kopinya enak", Rating: 4.3}

	dataCoworking[6] = CoworkingSpace{
		ID:           "CS007",
		Nama:         "The Hive",
		Lokasi:       "Jakarta",
		Fasilitas:    [FasilitasMax]string{"WiFi", "Pool Access", "AC", "Cafe"},
		JmlFasilitas: 4,
		HargaSewa:    250000,
		Rating:       4.9,
		JmlUlasan:    1,
	}
	dataCoworking[6].Ulasan[0] = Ulasan{ID: 1, Penulis: "John", Komentar: "Great view and internet", Rating: 4.9}

	dataCoworking[7] = CoworkingSpace{
		ID:           "CS008",
		Nama:         "Nomad Corner",
		Lokasi:       "Yogyakarta",
		Fasilitas:    [FasilitasMax]string{"WiFi", "AC"},
		JmlFasilitas: 2,
		HargaSewa:    95000,
		Rating:       4.1,
		JmlUlasan:    1,
	}
	dataCoworking[7].Ulasan[0] = Ulasan{ID: 1, Penulis: "Ali", Komentar: "Cukup memuaskan", Rating: 4.1}

	mainMenu(dataCoworking, nData)
}

func mainMenu(dataCoworking tabCoworking, nData int) {
	var lanjut bool = true

	for lanjut {
		fmt.Println("----------------------------------------------------------")
		fmt.Println("|     APLIKASI MANAJEMEN DAN REVIEW CO-WORKING SPACE     |")
		fmt.Println("|                Created By Tazky & Ratna                |")
		fmt.Println("|               Algoritma Pemrograman 2026               |")
		fmt.Println("----------------------------------------------------------")
		fmt.Println("|          1. Lihat Daftar Co-Working Space              |")
		fmt.Println("|          2. Tambah Daftar Co-Working Space             |")
		fmt.Println("|          3. Edit Daftar Co-Working Space               |")
		fmt.Println("|          4. Hapus Daftar Co-Working Space              |")
		fmt.Println("|          5. Kelola Ulasan                              |")
		fmt.Println("|          6. Pencarian                                  |")
		fmt.Println("|          7. Pengurutan                                 |")
		fmt.Println("|          0. Keluar                                     |")
		fmt.Println("----------------------------------------------------------")

		fmt.Print("Masukkan pilihan (1-7), ketik 0 untuk keluar: ")
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
			ulasanMenu(&dataCoworking, nData)
		} else if pilihan == 6 {
			cariMenu(dataCoworking, nData)
		} else if pilihan == 7 {
			sortMenu(dataCoworking, nData)
		} else if pilihan == 0 {
			fmt.Println("Terima kasih sudah menggunakan aplikasi ini")
			lanjut = false
		} else {
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func cariMenu(A tabCoworking, n int) {
	var lanjut bool = true

	for lanjut {
		fmt.Println("----------------------------------------------------------")
		fmt.Println("|     APLIKASI MANAJEMEN DAN REVIEW CO-WORKING SPACE     |")
		fmt.Println("|                Created By Tazky & Ratna                |")
		fmt.Println("|               Algoritma Pemrograman 2026               |")
		fmt.Println("----------------------------------------------------------")
		fmt.Println("|                   Cari berdasarkan:                    |")
		fmt.Println("|                       1. Nama                          |")
		fmt.Println("|                       2. Lokasi                        |")
		fmt.Println("|                       0. Kembali                       |")
		fmt.Println("----------------------------------------------------------")

		fmt.Print("Masukkan pilihan (1-2) atau ketik 0 untuk kembali: ")
		var pilihan int
		fmt.Scan(&pilihan)
		fmt.Println()

		if pilihan == 1 {
			cetakCoworking(A, n)

			var x string
			fmt.Print("Masukkan Nama Co-Working Space: ")
			fmt.Scan(&x)

			indeks := searchByNama(A, n, x)

			if indeks == -1 {
				break
			}

			// Cetak Coworking yang sudah dicari
			fmt.Printf("%-5s | %-5s | %-20s | %-15s | %-10s | %-6s\n",
				"No", "ID", "Nama", "Lokasi", "Harga Sewa", "Rating")
			fmt.Println("----------------------------------------------------------------------------")

			fmt.Printf("%-5d | %-5s | %-20s | %-15s | Rp%-8d | %-6.1f\n",
				1, A[indeks].ID, A[indeks].Nama, A[indeks].Lokasi, A[indeks].HargaSewa, A[indeks].Rating)

			fmt.Println()

		} else if pilihan == 2 {
			cetakCoworking(A, n)

			var x string
			fmt.Print("Masukkan Lokasi Co-Working Space: ")
			fmt.Scan(&x)

			sortingLokasi(&A, n)
			indeks := searchByLokasi(A, n, x)

			for indeks > 0 && A[indeks-1].Lokasi == x {
				indeks--
			}

			// Cetak Coworking yang berdasarkan lokasi yang dicari dan sudah diurutkan
			fmt.Printf("%-5s | %-5s | %-20s | %-15s | %-10s | %-6s\n",
				"No", "ID", "Nama", "Lokasi", "Harga Sewa", "Rating")
			fmt.Println("----------------------------------------------------------------------------")

			no := 1
			for i := indeks; i < n && A[i].Lokasi == x; i++ {
				fmt.Printf("%-5d | %-5s | %-20s | %-15s | Rp%-8d | %-6.1f\n",
					no, A[i].ID, A[i].Nama, A[i].Lokasi, A[i].HargaSewa, A[i].Rating)
				no++
			}
			fmt.Println()

		} else if pilihan == 0 {
			lanjut = false
		} else {
			fmt.Println("Pilihan tidak valid!")
			lanjut = false
		}
	}
}

func sortMenu(A tabCoworking, n int) {
	var lanjut bool = true

	for lanjut {
		fmt.Println("----------------------------------------------------------")
		fmt.Println("|     APLIKASI MANAJEMEN DAN REVIEW CO-WORKING SPACE     |")
		fmt.Println("|                Created By Tazky & Ratna                |")
		fmt.Println("|               Algoritma Pemrograman 2026               |")
		fmt.Println("----------------------------------------------------------")
		fmt.Println("|                   Urutkan berdasarkan:                 |")
		fmt.Println("|                       1. Harga Sewa                    |")
		fmt.Println("|                       2. Rating                        |")
		fmt.Println("|                       0. Kembali                       |")
		fmt.Println("----------------------------------------------------------")

		fmt.Print("Masukkan pilihan (1-2) atau ketik 0 untuk kembali: ")
		var pilihan int
		fmt.Scan(&pilihan)
		fmt.Println()

		if pilihan == 1 {
			sortingHargaSewa(&A, n)
			cetakCoworking(A, n)

		} else if pilihan == 2 {
			sortingRating(&A, n)
			cetakCoworking(A, n)

		} else if pilihan == 0 {
			lanjut = false
		} else {
			fmt.Println("Pilihan tidak valid!")
			lanjut = false
		}
	}
}

func ulasanMenu(A *tabCoworking, n int) {
	var lanjut bool = true

	for lanjut {
		fmt.Println("----------------------------------------------------------")
		fmt.Println("|     APLIKASI MANAJEMEN DAN REVIEW CO-WORKING SPACE     |")
		fmt.Println("|                Created By Tazky & Ratna                |")
		fmt.Println("|               Algoritma Pemrograman 2026               |")
		fmt.Println("----------------------------------------------------------")
		fmt.Println("|              MENU ULASAN CO-WORKING SPACE              |")
		fmt.Println("|                   1. Tambah Ulasan                     |")
		fmt.Println("|                   2. Edit Ulasan                       |")
		fmt.Println("|                   3. Hapus Ulasan                      |")
		fmt.Println("|                   0. Kembali                           |")
		fmt.Println("----------------------------------------------------------")

		fmt.Print("Masukkan pilihan (1-2) atau ketik 0 untuk kembali: ")
		var pilihan int
		fmt.Scan(&pilihan)
		fmt.Println()

		if pilihan == 1 {
			tambahUlasan(A, &n)
		} else if pilihan == 2 {
			editUlasan(A, n)
		} else if pilihan == 3 {
			hapusUlasan(A, n)
		} else if pilihan == 0 {
			lanjut = false
		} else {
			fmt.Println("Pilihan tidak valid!")
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
	fmt.Printf("Harga     : Rp%d / hari\n", A[indeks].HargaSewa)
	fmt.Printf("Rating    : %.1f / 5.0\n", A[indeks].Rating)

	// Cetak Fasilitas
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
	fmt.Println()

	// Cetak Ulasan
	fmt.Printf("Ulasan    : %d ulasan\n", A[indeks].JmlUlasan)
	if A[indeks].JmlUlasan > 0 {
		fmt.Println("---------------------------------------")
		for i := 0; i < A[indeks].JmlUlasan; i++ {
			u := A[indeks].Ulasan[i]
			fmt.Printf("  [%d] %s (%.1f★)\n", u.ID, u.Penulis, u.Rating)
			fmt.Printf("       \"%s\"\n", u.Komentar)
		}
	}
	fmt.Println("---------------------------------------\n")
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

func cetakUlasan(dataCoworking *CoworkingSpace) {
	fmt.Println("\n--- Daftar Ulasan Saat Ini ---")
	for i := 0; i < dataCoworking.JmlUlasan; i++ {
		u := dataCoworking.Ulasan[i]
		fmt.Printf("%d. [%d] Oleh: %s | \"%s\" (Rating: %.1f)\n", i+1, u.ID, u.Penulis, u.Komentar, u.Rating)
	}
	fmt.Println()
}

func tambahUlasan(A *tabCoworking, n *int) {
	for {
		cetakCoworking(*A, *n)

		var id string
		fmt.Print("\nMasukkan ID Co-Working Space yang mau ditambahkan ulasannya: ")
		fmt.Scan(&id)

		indeks := searchByID(*A, *n, id)

		if indeks == -1 {
			fmt.Println("Gagal: ID tidak ditemukan!")
			return
		}

		var dataCoworking = &A[indeks]
		cetakUlasan(dataCoworking)

		fmt.Println("Tambahkan Ulasan Baru:")

		var temp Ulasan
		fmt.Print("ID Ulasan                             : ")
		fmt.Scan(&temp.ID)
		fmt.Print("Nama Penulis (Gunakan _ untuk spasi)  : ")
		fmt.Scan(&temp.Penulis)
		fmt.Print("Ketik ulasan (Gunakan _ untuk spasi)  : ")
		fmt.Scan(&temp.Komentar)
		fmt.Print("Rating (Contoh: 4.5)                  : ")
		fmt.Scan(&temp.Rating)

		dataCoworking.Ulasan[dataCoworking.JmlUlasan] = temp
		dataCoworking.JmlUlasan++

		fmt.Printf("\n✅ Ulasan dari \"%s\" berhasil ditambahkan!\n\n", temp.Penulis)

		/* ============================================================================================================ */
		cetakUlasan(dataCoworking)

		fmt.Println("Apakah masih ada ulasan yang ingin anda tambah?")
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

func editUlasan(A *tabCoworking, n int) {
	for {
		cetakCoworking(*A, n)

		var id string
		fmt.Print("\nMasukkan ID Co-Working Space yang ulasannya mau diedit: ")
		fmt.Scan(&id)

		indeks := searchByID(*A, n, id)

		if indeks == -1 {
			fmt.Println("Gagal: ID tidak ditemukan!")
			return
		}

		var dataCoworking = &A[indeks]

		if dataCoworking.JmlUlasan == 0 {
			fmt.Println("Gagal: Tempat ini belum memiliki ulasan untuk diedit.")
			return
		}

		cetakUlasan(dataCoworking)

		var idUlasan int
		fmt.Print("\nMasukkan ID ulasan yang mau diedit: ")
		fmt.Scan(&idUlasan)

		// Cari indeks ulasan yang ingin diedit
		idxUlasan := -1
		for i := 0; i < dataCoworking.JmlUlasan; i++ {
			if dataCoworking.Ulasan[i].ID == idUlasan {
				idxUlasan = i
				break
			}
		}

		idUlasanLama := dataCoworking.Ulasan[idxUlasan].ID

		fmt.Println("Masukkan data ulasan baru:")
		fmt.Print("Penulis baru  : ")
		fmt.Scan(&dataCoworking.Ulasan[idxUlasan].Penulis)
		fmt.Print("Komentar baru : ")
		fmt.Scan(&dataCoworking.Ulasan[idxUlasan].Komentar)
		fmt.Print("Rating baru   : ")
		fmt.Scan(&dataCoworking.Ulasan[idxUlasan].Rating)

		dataCoworking.Ulasan[idxUlasan].ID = idUlasanLama

		fmt.Println("\n✅ Ulasan berhasil diperbarui!\n")

		/* ============================================================================================================ */
		cetakUlasan(dataCoworking)

		fmt.Println("Apakah masih ada ulasan yang ingin anda edit?")
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

func hapusUlasan(A *tabCoworking, n int) {
	for {
		cetakCoworking(*A, n)

		var id string
		fmt.Print("Masukkan ID Co-Working Space yang ulasannya mau dihapus: ")
		fmt.Scan(&id)

		indeks := searchByID(*A, n, id)

		if indeks == -1 {
			fmt.Println("Gagal: ID tidak ditemukan!")
			return
		}

		var dataCoworking = &A[indeks]

		if dataCoworking.JmlUlasan == 0 {
			fmt.Println("Gagal: Tempat ini belum memiliki ulasan untuk diedit.")
			return
		}

		cetakUlasan(dataCoworking)

		var idUlasan int
		fmt.Print("\nMasukkan ID ulasan yang mau dihapus: ")
		fmt.Scan(&idUlasan)

		idxUlasan := -1
		for i := 0; i < dataCoworking.JmlUlasan; i++ {
			if dataCoworking.Ulasan[i].ID == idUlasan {
				idxUlasan = i
				break
			}
		}

		for i := idxUlasan; i < dataCoworking.JmlUlasan-1; i++ {
			dataCoworking.Ulasan[i] = dataCoworking.Ulasan[i+1]
		}

		dataCoworking.Ulasan[dataCoworking.JmlUlasan-1] = Ulasan{}
		dataCoworking.JmlUlasan--

		fmt.Printf("\n✅ Ulasan dari \"%s\" berhasil dihapus!\n\n", dataCoworking.Ulasan[idxUlasan].Penulis)

		/* ============================================================================================================ */
		cetakUlasan(dataCoworking)

		fmt.Println("Apakah masih ada ulasan yang ingin anda hapus?")
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
		if A[i].Nama == x {
			idx = i
		}

		i++
	}

	return idx
}

func searchByLokasi(A tabCoworking, n int, x string) int {
	var left, right, mid int
	var idx int

	left = 0
	right = n - 1
	idx = -1

	for left <= right && idx == -1 {
		mid = (left + right) / 2

		if x < A[mid].Lokasi {
			right = mid - 1
		} else if x > A[mid].Lokasi {
			left = mid + 1
		} else {
			idx = mid
		}
	}

	return idx
}

func sortingLokasi(A *tabCoworking, n int) {
	var temp tabCoworking

	for i := 0; i < n; i++ {
		temp[i] = A[i]
	}

	for i := 1; i < n; i++ {
		kunci := temp[i]
		j := i - 1
		for j >= 0 && temp[j].Lokasi > kunci.Lokasi {
			temp[j+1] = temp[j]
			j--
		}
		temp[j+1] = kunci
	}

	for i := 0; i < n; i++ {
		A[i] = temp[i]
	}

}

func sortingHargaSewa(A *tabCoworking, n int) {
	for pass := 1; pass <= n-1; pass++ {
		idx := pass - 1

		for i := pass; i < n; i++ {
			if A[i].HargaSewa > A[idx].HargaSewa {
				idx = i
			}
		}

		temp := A[pass-1]
		A[pass-1] = A[idx]
		A[idx] = temp
	}
}

func sortingRating(A *tabCoworking, n int) {
	for pass := 1; pass < n; pass++ {
		i := pass
		temp := A[pass]

		for i > 0 && temp.Rating < A[i-1].Rating {
			A[i] = A[i-1]
			i--
		}

		A[i] = temp
	}
}
