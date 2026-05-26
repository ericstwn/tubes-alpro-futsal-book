package main

import "fmt"

const NMAX int = 100

type Lapangan struct {
	idLapangan   string
	namaLapangan string
	jenis        string
	hargaPerJam  int
	jamBuka      int
	jamTutup     int
}

var dataLapangan [NMAX]Lapangan

func cariIndexLapangan(n int, id string) int {
	var idx int = -1
	var i int = 0

	for i < n && idx == -1 {
		if dataLapangan[i].idLapangan == id {
			idx = i
		}
		i = i + 1
	}

	return idx
}

func tambahLapangan(n *int) {
	var id, nama, jenis string
	var harga, jamBuka, jamTutup int
	var idx int

	fmt.Println("========================================")
	fmt.Println("          TAMBAH DATA LAPANGAN")
	fmt.Println("========================================")

	if *n < NMAX {
		fmt.Print("Masukkan ID Lapangan       : ")
		fmt.Scan(&id)

		idx = cariIndexLapangan(*n, id)

		if idx == -1 {
			fmt.Print("Masukkan Nama Lapangan     : ")
			fmt.Scan(&nama)

			fmt.Print("Masukkan Jenis Lapangan    : ")
			fmt.Scan(&jenis)

			fmt.Print("Masukkan Harga per Jam     : ")
			fmt.Scan(&harga)

			for harga <= 0 {
				fmt.Println("Harga harus lebih dari 0.")
				fmt.Print("Masukkan Harga per Jam     : ")
				fmt.Scan(&harga)
			}

			fmt.Print("Masukkan Jam Buka          : ")
			fmt.Scan(&jamBuka)

			fmt.Print("Masukkan Jam Tutup         : ")
			fmt.Scan(&jamTutup)

			for jamTutup <= jamBuka {
				fmt.Println("Jam tutup harus lebih besar dari jam buka.")
				fmt.Print("Masukkan Jam Tutup         : ")
				fmt.Scan(&jamTutup)
			}

			dataLapangan[*n].idLapangan = id
			dataLapangan[*n].namaLapangan = nama
			dataLapangan[*n].jenis = jenis
			dataLapangan[*n].hargaPerJam = harga
			dataLapangan[*n].jamBuka = jamBuka
			dataLapangan[*n].jamTutup = jamTutup

			*n = *n + 1

			fmt.Println()
			fmt.Println("Data lapangan berhasil ditambahkan.")
		} else {
			fmt.Println()
			fmt.Println("ID lapangan sudah digunakan.")
		}
	} else {
		fmt.Println("Data lapangan sudah penuh.")
	}
}

func tampilLapangan(n int) {
	var i int

	fmt.Println("======================================================================")
	fmt.Println("                           DATA LAPANGAN")
	fmt.Println("======================================================================")

	if n > 0 {
		fmt.Printf("%-6s %-18s %-12s %-12s %-8s %-8s\n",
			"ID", "Nama", "Jenis", "Harga/Jam", "Buka", "Tutup")
		fmt.Println("----------------------------------------------------------------------")

		for i = 0; i < n; i++ {
			fmt.Printf("%-6s %-18s %-12s %-12d %-8d %-8d\n",
				dataLapangan[i].idLapangan,
				dataLapangan[i].namaLapangan,
				dataLapangan[i].jenis,
				dataLapangan[i].hargaPerJam,
				dataLapangan[i].jamBuka,
				dataLapangan[i].jamTutup)
		}
	} else {
		fmt.Println("Belum ada data lapangan.")
	}

	fmt.Println("======================================================================")
}

func ubahLapangan(n int) {
	var id string
	var idx int
	var namaBaru, jenisBaru string
	var hargaBaru, jamBukaBaru, jamTutupBaru int

	fmt.Println("========================================")
	fmt.Println("           UBAH DATA LAPANGAN")
	fmt.Println("========================================")

	if n > 0 {
		fmt.Print("Masukkan ID Lapangan yang akan diubah: ")
		fmt.Scan(&id)

		idx = cariIndexLapangan(n, id)

		if idx != -1 {
			fmt.Println()
			fmt.Println("Data lama:")
			fmt.Println("ID Lapangan   :", dataLapangan[idx].idLapangan)
			fmt.Println("Nama Lapangan :", dataLapangan[idx].namaLapangan)
			fmt.Println("Jenis         :", dataLapangan[idx].jenis)
			fmt.Println("Harga/Jam     :", dataLapangan[idx].hargaPerJam)
			fmt.Println("Jam Buka      :", dataLapangan[idx].jamBuka)
			fmt.Println("Jam Tutup     :", dataLapangan[idx].jamTutup)
			fmt.Println()

			fmt.Print("Masukkan Nama Lapangan Baru : ")
			fmt.Scan(&namaBaru)

			fmt.Print("Masukkan Jenis Baru         : ")
			fmt.Scan(&jenisBaru)

			fmt.Print("Masukkan Harga Baru         : ")
			fmt.Scan(&hargaBaru)

			for hargaBaru <= 0 {
				fmt.Println("Harga harus lebih dari 0.")
				fmt.Print("Masukkan Harga Baru         : ")
				fmt.Scan(&hargaBaru)
			}

			fmt.Print("Masukkan Jam Buka Baru      : ")
			fmt.Scan(&jamBukaBaru)

			fmt.Print("Masukkan Jam Tutup Baru     : ")
			fmt.Scan(&jamTutupBaru)

			for jamTutupBaru <= jamBukaBaru {
				fmt.Println("Jam tutup harus lebih besar dari jam buka.")
				fmt.Print("Masukkan Jam Tutup Baru     : ")
				fmt.Scan(&jamTutupBaru)
			}

			dataLapangan[idx].namaLapangan = namaBaru
			dataLapangan[idx].jenis = jenisBaru
			dataLapangan[idx].hargaPerJam = hargaBaru
			dataLapangan[idx].jamBuka = jamBukaBaru
			dataLapangan[idx].jamTutup = jamTutupBaru

			fmt.Println()
			fmt.Println("Data lapangan berhasil diubah.")
		} else {
			fmt.Println()
			fmt.Println("Data lapangan tidak ditemukan.")
		}
	} else {
		fmt.Println("Belum ada data lapangan.")
	}
}

func hapusLapangan(n *int) {
	var id string
	var idx int
	var i int

	fmt.Println("========================================")
	fmt.Println("          HAPUS DATA LAPANGAN")
	fmt.Println("========================================")

	if *n > 0 {
		fmt.Print("Masukkan ID Lapangan yang akan dihapus: ")
		fmt.Scan(&id)

		idx = cariIndexLapangan(*n, id)

		if idx != -1 {
			for i = idx; i < *n-1; i++ {
				dataLapangan[i] = dataLapangan[i+1]
			}

			dataLapangan[*n-1] = Lapangan{}
			*n = *n - 1

			fmt.Println()
			fmt.Println("Data lapangan berhasil dihapus.")
		} else {
			fmt.Println()
			fmt.Println("Data lapangan tidak ditemukan.")
		}
	} else {
		fmt.Println("Belum ada data lapangan.")
	}
}

func menuLapangan(nLapangan *int) {
	var pilihan int = -1

	for pilihan != 0 {
		fmt.Println()
		fmt.Println("========================================")
		fmt.Println("          KELOLA DATA LAPANGAN")
		fmt.Println("========================================")
		fmt.Println("1. Tambah Data Lapangan")
		fmt.Println("2. Tampilkan Data Lapangan")
		fmt.Println("3. Ubah Data Lapangan")
		fmt.Println("4. Hapus Data Lapangan")
		fmt.Println("0. Kembali")
		fmt.Println("========================================")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)
		fmt.Println()

		if pilihan == 1 {
			tambahLapangan(nLapangan)
		} else if pilihan == 2 {
			tampilLapangan(*nLapangan)
		} else if pilihan == 3 {
			ubahLapangan(*nLapangan)
		} else if pilihan == 4 {
			hapusLapangan(nLapangan)
		} else if pilihan == 0 {
			fmt.Println("Kembali ke menu utama.")
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func menuUtama(nLapangan *int) {
	var pilihan int = -1

	for pilihan != 0 {
		fmt.Println()
		fmt.Println("========================================")
		fmt.Println("              FUTSAL-BOOK")
		fmt.Println("   Aplikasi Pemesanan Lapangan Futsal")
		fmt.Println("========================================")
		fmt.Println("1. Kelola Data Lapangan")
		fmt.Println("2. Kelola Data Penyewa")
		fmt.Println("3. Transaksi Penyewaan")
		fmt.Println("4. Cari Data Penyewa")
		fmt.Println("5. Lihat Jadwal Kosong")
		fmt.Println("6. Statistik")
		fmt.Println("0. Keluar")
		fmt.Println("========================================")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)
		fmt.Println()

		if pilihan == 1 {
			menuLapangan(nLapangan)
		} else if pilihan == 2 {
			fmt.Println("Menu Data Penyewa belum digabung.")
		} else if pilihan == 3 {
			fmt.Println("Menu Transaksi Penyewaan belum digabung.")
		} else if pilihan == 4 {
			fmt.Println("Menu Cari Data Penyewa belum digabung.")
		} else if pilihan == 5 {
			fmt.Println("Menu Jadwal Kosong belum digabung.")
		} else if pilihan == 6 {
			fmt.Println("Menu Statistik belum digabung.")
		} else if pilihan == 0 {
			fmt.Println("Terima kasih telah menggunakan Futsal-Book.")
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func main() {
	var nLapangan int = 0

	menuUtama(&nLapangan)
}
