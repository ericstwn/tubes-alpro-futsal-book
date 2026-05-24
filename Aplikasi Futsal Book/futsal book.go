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

type Penyewa struct {
	idPenyewa   string
	namaPenyewa string
	noTelepon   string
}

var dataLapangan [NMAX]Lapangan
var dataPenyewa [NMAX]Penyewa

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

func cariIndexPenyewa(n int, id string) int {
	var idx int = -1
	var i int = 0

	for i < n && idx == -1 {
		if dataPenyewa[i].idPenyewa == id {
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

			fmt.Println("Data lapangan berhasil ditambahkan.")
		} else {
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

			fmt.Println("Data lapangan berhasil diubah.")
		} else {
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

			fmt.Println("Data lapangan berhasil dihapus.")
		} else {
			fmt.Println("Data lapangan tidak ditemukan.")
		}
	} else {
		fmt.Println("Belum ada data lapangan.")
	}
}

func tambahPenyewa(n *int) {
	var id, nama, noTelp string
	var idx int

	fmt.Println("========================================")
	fmt.Println("          TAMBAH DATA PENYEWA")
	fmt.Println("========================================")

	if *n < NMAX {
		fmt.Print("Masukkan ID Penyewa     : ")
		fmt.Scan(&id)

		idx = cariIndexPenyewa(*n, id)

		if idx == -1 {
			fmt.Print("Masukkan Nama Penyewa   : ")
			fmt.Scan(&nama)

			fmt.Print("Masukkan Nomor Telepon  : ")
			fmt.Scan(&noTelp)

			dataPenyewa[*n].idPenyewa = id
			dataPenyewa[*n].namaPenyewa = nama
			dataPenyewa[*n].noTelepon = noTelp

			*n = *n + 1

			fmt.Println("Data penyewa berhasil ditambahkan.")
		} else {
			fmt.Println("ID penyewa sudah digunakan.")
		}
	} else {
		fmt.Println("Data penyewa sudah penuh.")
	}
}

func tampilPenyewa(n int) {
	var i int

	fmt.Println("====================================================")
	fmt.Println("                   DATA PENYEWA")
	fmt.Println("====================================================")

	if n > 0 {
		fmt.Printf("%-6s %-20s %-15s\n", "ID", "Nama", "No Telepon")
		fmt.Println("----------------------------------------------------")

		for i = 0; i < n; i++ {
			fmt.Printf("%-6s %-20s %-15s\n",
				dataPenyewa[i].idPenyewa,
				dataPenyewa[i].namaPenyewa,
				dataPenyewa[i].noTelepon)
		}
	} else {
		fmt.Println("Belum ada data penyewa.")
	}

	fmt.Println("====================================================")
}

func ubahPenyewa(n int) {
	var id string
	var idx int
	var namaBaru, noTelpBaru string

	fmt.Println("========================================")
	fmt.Println("           UBAH DATA PENYEWA")
	fmt.Println("========================================")

	if n > 0 {
		fmt.Print("Masukkan ID Penyewa yang akan diubah: ")
		fmt.Scan(&id)

		idx = cariIndexPenyewa(n, id)

		if idx != -1 {
			fmt.Println()
			fmt.Println("Data lama:")
			fmt.Println("ID Penyewa    :", dataPenyewa[idx].idPenyewa)
			fmt.Println("Nama Penyewa  :", dataPenyewa[idx].namaPenyewa)
			fmt.Println("No Telepon    :", dataPenyewa[idx].noTelepon)
			fmt.Println()

			fmt.Print("Masukkan Nama Penyewa Baru  : ")
			fmt.Scan(&namaBaru)

			fmt.Print("Masukkan Nomor Telepon Baru : ")
			fmt.Scan(&noTelpBaru)

			dataPenyewa[idx].namaPenyewa = namaBaru
			dataPenyewa[idx].noTelepon = noTelpBaru

			fmt.Println("Data penyewa berhasil diubah.")
		} else {
			fmt.Println("Data penyewa tidak ditemukan.")
		}
	} else {
		fmt.Println("Belum ada data penyewa.")
	}
}

func hapusPenyewa(n *int) {
	var id string
	var idx int
	var i int

	fmt.Println("========================================")
	fmt.Println("          HAPUS DATA PENYEWA")
	fmt.Println("========================================")

	if *n > 0 {
		fmt.Print("Masukkan ID Penyewa yang akan dihapus: ")
		fmt.Scan(&id)

		idx = cariIndexPenyewa(*n, id)

		if idx != -1 {
			for i = idx; i < *n-1; i++ {
				dataPenyewa[i] = dataPenyewa[i+1]
			}

			dataPenyewa[*n-1] = Penyewa{}
			*n = *n - 1

			fmt.Println("Data penyewa berhasil dihapus.")
		} else {
			fmt.Println("Data penyewa tidak ditemukan.")
		}
	} else {
		fmt.Println("Belum ada data penyewa.")
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

// ======================================================
// Menu Data Penyewa
// ======================================================
func menuPenyewa(nPenyewa *int) {
	var pilihan int = -1

	for pilihan != 0 {
		fmt.Println()
		fmt.Println("========================================")
		fmt.Println("          KELOLA DATA PENYEWA")
		fmt.Println("========================================")
		fmt.Println("1. Tambah Data Penyewa")
		fmt.Println("2. Tampilkan Data Penyewa")
		fmt.Println("3. Ubah Data Penyewa")
		fmt.Println("4. Hapus Data Penyewa")
		fmt.Println("0. Kembali")
		fmt.Println("========================================")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)
		fmt.Println()

		if pilihan == 1 {
			tambahPenyewa(nPenyewa)
		} else if pilihan == 2 {
			tampilPenyewa(*nPenyewa)
		} else if pilihan == 3 {
			ubahPenyewa(*nPenyewa)
		} else if pilihan == 4 {
			hapusPenyewa(nPenyewa)
		} else if pilihan == 0 {
			fmt.Println("Kembali ke menu utama.")
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

// ======================================================
// Menu Utama
// ======================================================
func menuUtama(nLapangan *int, nPenyewa *int) {
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
			menuPenyewa(nPenyewa)
		} else if pilihan == 3 {
			fmt.Println("Menu transaksi penyewaan belum dikerjakan.")
		} else if pilihan == 4 {
			fmt.Println("Menu cari data penyewa belum dikerjakan.")
		} else if pilihan == 5 {
			fmt.Println("Menu jadwal kosong belum dikerjakan.")
		} else if pilihan == 6 {
			fmt.Println("Menu statistik belum dikerjakan.")
		} else if pilihan == 0 {
			fmt.Println("Terima kasih telah menggunakan Futsal-Book.")
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

// ======================================================
// Program Utama
// ======================================================
func main() {
	var nLapangan int = 0
	var nPenyewa int = 0

	menuUtama(&nLapangan, &nPenyewa)
}
