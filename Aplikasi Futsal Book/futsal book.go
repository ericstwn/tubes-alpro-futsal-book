package main

import (
	"fmt"
	"os"
)

const NMAX int = 100

type Lapangan struct {
	ID       string
	Nama     string
	Jenis    string
	Harga    int
	JamBuka  int
	JamTutup int
}

type Penyewa struct {
	ID     string
	Nama   string
	NoTelp string
}

type Sewa struct {
	IDSewa     string
	IDPenyewa  string
	IDLapangan string
	Tanggal    int
	Bulan      int
	Tahun      int
	JamMulai   int
	JamSelesai int
	TotalHarga int
}

type JadwalKosong struct {
	IDLapangan string
	JamMulai   int
	JamSelesai int
	Harga      int
}

var dataLapangan [NMAX]Lapangan
var dataPenyewa [NMAX]Penyewa
var dataSewa [NMAX]Sewa
var dataJadwal [NMAX]JadwalKosong


func simpanDataLapangan(n int) {
	var file *os.File
	var err error
	var i int

	file, err = os.Create("lapangan.txt")

	if err == nil {
		fmt.Fprintln(file, n)

		for i = 0; i < n; i++ {
			fmt.Fprintln(file,
				dataLapangan[i].ID,
				dataLapangan[i].Nama,
				dataLapangan[i].Jenis,
				dataLapangan[i].Harga,
				dataLapangan[i].JamBuka,
				dataLapangan[i].JamTutup)
		}

		file.Close()
	}
}

func bacaDataLapangan(n *int) {
	var file *os.File
	var err error
	var jumlah int
	var i int

	file, err = os.Open("lapangan.txt")

	if err == nil {
		fmt.Fscan(file, &jumlah)

		if jumlah > NMAX {
			jumlah = NMAX
		}

		*n = jumlah

		for i = 0; i < *n; i++ {
			fmt.Fscan(file,
				&dataLapangan[i].ID,
				&dataLapangan[i].Nama,
				&dataLapangan[i].Jenis,
				&dataLapangan[i].Harga,
				&dataLapangan[i].JamBuka,
				&dataLapangan[i].JamTutup)
		}

		file.Close()
	}
}

func simpanDataPenyewa(n int) {
	var file *os.File
	var err error
	var i int

	file, err = os.Create("penyewa.txt")

	if err == nil {
		fmt.Fprintln(file, n)

		for i = 0; i < n; i++ {
			fmt.Fprintln(file,
				dataPenyewa[i].ID,
				dataPenyewa[i].Nama,
				dataPenyewa[i].NoTelp)
		}

		file.Close()
	}
}

func bacaDataPenyewa(n *int) {
	var file *os.File
	var err error
	var jumlah int
	var i int

	file, err = os.Open("penyewa.txt")

	if err == nil {
		fmt.Fscan(file, &jumlah)

		if jumlah > NMAX {
			jumlah = NMAX
		}

		*n = jumlah

		for i = 0; i < *n; i++ {
			fmt.Fscan(file,
				&dataPenyewa[i].ID,
				&dataPenyewa[i].Nama,
				&dataPenyewa[i].NoTelp)
		}

		file.Close()
	}
}

func simpanDataSewa(n int) {
	var file *os.File
	var err error
	var i int

	file, err = os.Create("sewa.txt")

	if err == nil {
		fmt.Fprintln(file, n)

		for i = 0; i < n; i++ {
			fmt.Fprintln(file,
				dataSewa[i].IDSewa,
				dataSewa[i].IDPenyewa,
				dataSewa[i].IDLapangan,
				dataSewa[i].Tanggal,
				dataSewa[i].Bulan,
				dataSewa[i].Tahun,
				dataSewa[i].JamMulai,
				dataSewa[i].JamSelesai,
				dataSewa[i].TotalHarga)
		}

		file.Close()
	}
}

func bacaDataSewa(n *int) {
	var file *os.File
	var err error
	var jumlah int
	var i int

	file, err = os.Open("sewa.txt")

	if err == nil {
		fmt.Fscan(file, &jumlah)

		if jumlah > NMAX {
			jumlah = NMAX
		}

		*n = jumlah

		for i = 0; i < *n; i++ {
			fmt.Fscan(file,
				&dataSewa[i].IDSewa,
				&dataSewa[i].IDPenyewa,
				&dataSewa[i].IDLapangan,
				&dataSewa[i].Tanggal,
				&dataSewa[i].Bulan,
				&dataSewa[i].Tahun,
				&dataSewa[i].JamMulai,
				&dataSewa[i].JamSelesai,
				&dataSewa[i].TotalHarga)
		}

		file.Close()
	}
}

func cariIndexLapangan(n int, id string) int {
	var idx int = -1
	var i int = 0

	for i < n && idx == -1 {
		if dataLapangan[i].ID == id {
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
		if dataPenyewa[i].ID == id {
			idx = i
		}
		i = i + 1
	}

	return idx
}

func cariIndexSewa(n int, id string) int {
	var idx int = -1
	var i int = 0

	for i < n && idx == -1 {
		if dataSewa[i].IDSewa == id {
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

			dataLapangan[*n].ID = id
			dataLapangan[*n].Nama = nama
			dataLapangan[*n].Jenis = jenis
			dataLapangan[*n].Harga = harga
			dataLapangan[*n].JamBuka = jamBuka
			dataLapangan[*n].JamTutup = jamTutup

			*n = *n + 1
			simpanDataLapangan(*n)

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
		fmt.Printf("%-6s %-18s %-12s %-12s %-8s %-8s\n", "ID", "Nama", "Jenis", "Harga/Jam", "Buka", "Tutup")
		fmt.Println("----------------------------------------------------------------------")

		for i = 0; i < n; i++ {
			fmt.Printf("%-6s %-18s %-12s %-12d %-8d %-8d\n",
				dataLapangan[i].ID,
				dataLapangan[i].Nama,
				dataLapangan[i].Jenis,
				dataLapangan[i].Harga,
				dataLapangan[i].JamBuka,
				dataLapangan[i].JamTutup)
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
			fmt.Println("Data lama:")
			fmt.Println("ID Lapangan   :", dataLapangan[idx].ID)
			fmt.Println("Nama Lapangan :", dataLapangan[idx].Nama)
			fmt.Println("Jenis         :", dataLapangan[idx].Jenis)
			fmt.Println("Harga/Jam     :", dataLapangan[idx].Harga)
			fmt.Println("Jam Buka      :", dataLapangan[idx].JamBuka)
			fmt.Println("Jam Tutup     :", dataLapangan[idx].JamTutup)
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

			dataLapangan[idx].Nama = namaBaru
			dataLapangan[idx].Jenis = jenisBaru
			dataLapangan[idx].Harga = hargaBaru
			dataLapangan[idx].JamBuka = jamBukaBaru
			dataLapangan[idx].JamTutup = jamTutupBaru
			simpanDataLapangan(n)

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
			simpanDataLapangan(*n)

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

			dataPenyewa[*n].ID = id
			dataPenyewa[*n].Nama = nama
			dataPenyewa[*n].NoTelp = noTelp

			*n = *n + 1
			simpanDataPenyewa(*n)

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
				dataPenyewa[i].ID,
				dataPenyewa[i].Nama,
				dataPenyewa[i].NoTelp)
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
			fmt.Println("Data lama:")
			fmt.Println("ID Penyewa   :", dataPenyewa[idx].ID)
			fmt.Println("Nama Penyewa :", dataPenyewa[idx].Nama)
			fmt.Println("No Telepon   :", dataPenyewa[idx].NoTelp)
			fmt.Println()

			fmt.Print("Masukkan Nama Penyewa Baru  : ")
			fmt.Scan(&namaBaru)

			fmt.Print("Masukkan Nomor Telepon Baru : ")
			fmt.Scan(&noTelpBaru)

			dataPenyewa[idx].Nama = namaBaru
			dataPenyewa[idx].NoTelp = noTelpBaru
			simpanDataPenyewa(n)

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
			simpanDataPenyewa(*n)

			fmt.Println("Data penyewa berhasil dihapus.")
		} else {
			fmt.Println("Data penyewa tidak ditemukan.")
		}
	} else {
		fmt.Println("Belum ada data penyewa.")
	}
}

func sequentialSearchPenyewaByNama(n int, nama string) int {
	var idx int = -1
	var i int = 0

	for i < n && idx == -1 {
		if dataPenyewa[i].Nama == nama {
			idx = i
		}
		i = i + 1
	}

	return idx
}

func insertionSortPenyewaByTelp(n int, asc bool) {
	var pass, i int
	var temp Penyewa

	for pass = 1; pass < n; pass++ {
		temp = dataPenyewa[pass]
		i = pass

		if asc {
			for i > 0 && temp.NoTelp < dataPenyewa[i-1].NoTelp {
				dataPenyewa[i] = dataPenyewa[i-1]
				i = i - 1
			}
		} else {
			for i > 0 && temp.NoTelp > dataPenyewa[i-1].NoTelp {
				dataPenyewa[i] = dataPenyewa[i-1]
				i = i - 1
			}
		}

		dataPenyewa[i] = temp
	}
}

func binarySearchPenyewaByTelp(n int, noTelp string) int {
	var kiri, kanan, tengah int
	var idx int = -1

	kiri = 0
	kanan = n - 1

	for kiri <= kanan && idx == -1 {
		tengah = (kiri + kanan) / 2

		if dataPenyewa[tengah].NoTelp == noTelp {
			idx = tengah
		} else if dataPenyewa[tengah].NoTelp < noTelp {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}

	return idx
}

func tampilSatuPenyewa(idx int) {
	fmt.Println("ID Penyewa   :", dataPenyewa[idx].ID)
	fmt.Println("Nama Penyewa :", dataPenyewa[idx].Nama)
	fmt.Println("No Telepon   :", dataPenyewa[idx].NoTelp)
}

func cariIndexLapanganTransaksi(n int, id string) int {
	return cariIndexLapangan(n, id)
}

func cariIndexPenyewaTransaksi(n int, id string) int {
	return cariIndexPenyewa(n, id)
}

func cekKetersediaan(idLapangan string, tanggal int, bulan int, tahun int, jamMulai int, jamSelesai int, nSewa int) bool {
	var i int
	var tersedia bool = true

	for i = 0; i < nSewa && tersedia; i++ {
		if dataSewa[i].IDLapangan == idLapangan && dataSewa[i].Tanggal == tanggal && dataSewa[i].Bulan == bulan && dataSewa[i].Tahun == tahun {
			if jamMulai < dataSewa[i].JamSelesai && jamSelesai > dataSewa[i].JamMulai {
				tersedia = false
			}
		}
	}

	return tersedia
}

func tambahSewa(nSewa *int, nLapangan int, nPenyewa int) {
	var sewa Sewa
	var idxLapangan int
	var idxPenyewa int
	var idxSewa int
	var valid bool = true

	fmt.Println()
	fmt.Println("========================================")
	fmt.Println("             TAMBAH SEWA")
	fmt.Println("========================================")

	if *nSewa < NMAX {
		fmt.Print("ID Sewa      : ")
		fmt.Scan(&sewa.IDSewa)

		idxSewa = cariIndexSewa(*nSewa, sewa.IDSewa)

		if idxSewa != -1 {
			fmt.Println("ID sewa sudah digunakan.")
			valid = false
		}

		if valid {
			fmt.Print("ID Penyewa   : ")
			fmt.Scan(&sewa.IDPenyewa)

			idxPenyewa = cariIndexPenyewaTransaksi(nPenyewa, sewa.IDPenyewa)

			if idxPenyewa == -1 {
				fmt.Println("Penyewa tidak ditemukan.")
				valid = false
			}
		}

		if valid {
			fmt.Print("ID Lapangan  : ")
			fmt.Scan(&sewa.IDLapangan)

			idxLapangan = cariIndexLapanganTransaksi(nLapangan, sewa.IDLapangan)

			if idxLapangan == -1 {
				fmt.Println("Lapangan tidak ditemukan.")
				valid = false
			}
		}

		if valid {
			fmt.Print("Tanggal      : ")
			fmt.Scan(&sewa.Tanggal)

			fmt.Print("Bulan        : ")
			fmt.Scan(&sewa.Bulan)

			fmt.Print("Tahun        : ")
			fmt.Scan(&sewa.Tahun)

			fmt.Print("Jam Mulai    : ")
			fmt.Scan(&sewa.JamMulai)

			fmt.Print("Jam Selesai  : ")
			fmt.Scan(&sewa.JamSelesai)

			if sewa.JamSelesai <= sewa.JamMulai {
				fmt.Println("Jam selesai harus lebih besar dari jam mulai.")
				valid = false
			}
		}

		if valid {
			if sewa.JamMulai < dataLapangan[idxLapangan].JamBuka || sewa.JamSelesai > dataLapangan[idxLapangan].JamTutup {
				fmt.Println("Jam sewa berada di luar jam operasional lapangan.")
				valid = false
			}
		}

		if valid {
			if cekKetersediaan(sewa.IDLapangan, sewa.Tanggal, sewa.Bulan, sewa.Tahun, sewa.JamMulai, sewa.JamSelesai, *nSewa) {
				sewa.TotalHarga = (sewa.JamSelesai - sewa.JamMulai) * dataLapangan[idxLapangan].Harga
				dataSewa[*nSewa] = sewa
				*nSewa = *nSewa + 1
				simpanDataSewa(*nSewa)

				fmt.Println("Data sewa berhasil ditambahkan.")
				fmt.Println("Total harga:", sewa.TotalHarga)
			} else {
				fmt.Println("Jadwal bentrok, lapangan sudah dipesan pada jam tersebut.")
			}
		}
	} else {
		fmt.Println("Data sewa sudah penuh.")
	}
}

func tampilSewa(nSewa int) {
	var i int

	fmt.Println()
	fmt.Println("======================================================================")
	fmt.Println("                              DATA SEWA")
	fmt.Println("======================================================================")

	if nSewa > 0 {
		for i = 0; i < nSewa; i++ {
			fmt.Println("------------------------------------------------------------------")
			fmt.Println("ID Sewa     :", dataSewa[i].IDSewa)
			fmt.Println("ID Penyewa  :", dataSewa[i].IDPenyewa)
			fmt.Println("ID Lapangan :", dataSewa[i].IDLapangan)
			fmt.Println("Tanggal     :", dataSewa[i].Tanggal, "/", dataSewa[i].Bulan, "/", dataSewa[i].Tahun)
			fmt.Println("Jam         :", dataSewa[i].JamMulai, "-", dataSewa[i].JamSelesai)
			fmt.Println("Total Harga :", dataSewa[i].TotalHarga)
		}
	} else {
		fmt.Println("Belum ada data sewa.")
	}

	fmt.Println("======================================================================")
}

func buatDataJadwalKosong(nLapangan int, nSewa int, tanggal int, bulan int, tahun int) int {
	var i, j, jam int
	var nJadwal int = 0
	var kosong bool

	for i = 0; i < nLapangan; i++ {
		for jam = dataLapangan[i].JamBuka; jam < dataLapangan[i].JamTutup; jam++ {
			kosong = true

			for j = 0; j < nSewa && kosong; j++ {
				if dataSewa[j].IDLapangan == dataLapangan[i].ID && dataSewa[j].Tanggal == tanggal && dataSewa[j].Bulan == bulan && dataSewa[j].Tahun == tahun {
					if jam >= dataSewa[j].JamMulai && jam < dataSewa[j].JamSelesai {
						kosong = false
					}
				}
			}

			if kosong && nJadwal < NMAX {
				dataJadwal[nJadwal].IDLapangan = dataLapangan[i].ID
				dataJadwal[nJadwal].JamMulai = jam
				dataJadwal[nJadwal].JamSelesai = jam + 1
				dataJadwal[nJadwal].Harga = dataLapangan[i].Harga
				nJadwal = nJadwal + 1
			}
		}
	}

	return nJadwal
}

func tampilDataJadwalKosong(nJadwal int) {
	var i int

	fmt.Println()
	fmt.Println("===============================================================")
	fmt.Println("                       JADWAL KOSONG")
	fmt.Println("===============================================================")

	if nJadwal > 0 {
		fmt.Printf("%-12s %-10s %-10s %-12s\n", "ID Lapangan", "Mulai", "Selesai", "Harga")
		fmt.Println("---------------------------------------------------------------")

		for i = 0; i < nJadwal; i++ {
			fmt.Printf("%-12s %-10d %-10d %-12d\n",
				dataJadwal[i].IDLapangan,
				dataJadwal[i].JamMulai,
				dataJadwal[i].JamSelesai,
				dataJadwal[i].Harga)
		}
	} else {
		fmt.Println("Tidak ada jadwal kosong atau data lapangan belum tersedia.")
	}

	fmt.Println("===============================================================")
}

func selectionSortJadwalByJam(n int, asc bool) {
	var pass, idx, i int
	var temp JadwalKosong

	for pass = 0; pass < n-1; pass++ {
		idx = pass

		for i = pass + 1; i < n; i++ {
			if asc {
				if dataJadwal[i].JamMulai < dataJadwal[idx].JamMulai {
					idx = i
				}
			} else {
				if dataJadwal[i].JamMulai > dataJadwal[idx].JamMulai {
					idx = i
				}
			}
		}

		temp = dataJadwal[pass]
		dataJadwal[pass] = dataJadwal[idx]
		dataJadwal[idx] = temp
	}
}

func insertionSortJadwalByHarga(n int, asc bool) {
	var pass, i int
	var temp JadwalKosong

	for pass = 1; pass < n; pass++ {
		temp = dataJadwal[pass]
		i = pass

		if asc {
			for i > 0 && temp.Harga < dataJadwal[i-1].Harga {
				dataJadwal[i] = dataJadwal[i-1]
				i = i - 1
			}
		} else {
			for i > 0 && temp.Harga > dataJadwal[i-1].Harga {
				dataJadwal[i] = dataJadwal[i-1]
				i = i - 1
			}
		}

		dataJadwal[i] = temp
	}
}

func totalPendapatanBulanan(nSewa int, bulan int) int {
	var total int = 0
	var i int

	for i = 0; i < nSewa; i++ {
		if dataSewa[i].Bulan == bulan {
			total = total + dataSewa[i].TotalHarga
		}
	}

	return total
}

func jamPalingSeringDipesan(nSewa int) int {
	var frek [24]int
	var i, jam int
	var maxJam int = 0
	var maxFrek int = 0

	for i = 0; i < nSewa; i++ {
		for jam = dataSewa[i].JamMulai; jam < dataSewa[i].JamSelesai; jam++ {
			if jam >= 0 && jam < 24 {
				frek[jam] = frek[jam] + 1
			}
		}
	}

	for jam = 0; jam < 24; jam++ {
		if frek[jam] > maxFrek {
			maxFrek = frek[jam]
			maxJam = jam
		}
	}

	return maxJam
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

func menuCariPenyewa(nPenyewa int) {
	var pilihan int = -1
	var nama, noTelp string
	var idx int

	for pilihan != 0 {
		fmt.Println()
		fmt.Println("========================================")
		fmt.Println("            CARI DATA PENYEWA")
		fmt.Println("========================================")
		fmt.Println("1. Sequential Search berdasarkan Nama")
		fmt.Println("2. Binary Search berdasarkan No Telepon")
		fmt.Println("0. Kembali")
		fmt.Println("========================================")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)
		fmt.Println()

		if pilihan == 1 {
			fmt.Print("Masukkan nama penyewa: ")
			fmt.Scan(&nama)
			idx = sequentialSearchPenyewaByNama(nPenyewa, nama)

			if idx != -1 {
				fmt.Println("Data penyewa ditemukan:")
				tampilSatuPenyewa(idx)
			} else {
				fmt.Println("Data penyewa tidak ditemukan.")
			}
		} else if pilihan == 2 {
			insertionSortPenyewaByTelp(nPenyewa, true)
			fmt.Print("Masukkan nomor telepon: ")
			fmt.Scan(&noTelp)
			idx = binarySearchPenyewaByTelp(nPenyewa, noTelp)

			if idx != -1 {
				fmt.Println("Data penyewa ditemukan:")
				tampilSatuPenyewa(idx)
			} else {
				fmt.Println("Data penyewa tidak ditemukan.")
			}
		} else if pilihan == 0 {
			fmt.Println("Kembali ke menu utama.")
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func menuTransaksi(nSewa *int, nLapangan int, nPenyewa int) {
	var pilihan int = -1

	for pilihan != 0 {
		fmt.Println()
		fmt.Println("========================================")
		fmt.Println("             MENU TRANSAKSI")
		fmt.Println("========================================")
		fmt.Println("1. Tambah Sewa")
		fmt.Println("2. Tampil Sewa")
		fmt.Println("0. Kembali")
		fmt.Println("========================================")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)
		fmt.Println()

		if pilihan == 1 {
			tambahSewa(nSewa, nLapangan, nPenyewa)
		} else if pilihan == 2 {
			tampilSewa(*nSewa)
		} else if pilihan == 0 {
			fmt.Println("Kembali ke menu utama.")
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func menuJadwalKosong(nLapangan int, nSewa int) {
	var tanggal, bulan, tahun int
	var pilihanSort, urutan int
	var asc bool = true
	var nJadwal int

	fmt.Println()
	fmt.Println("========================================")
	fmt.Println("             JADWAL KOSONG")
	fmt.Println("========================================")
	fmt.Print("Masukkan tanggal: ")
	fmt.Scan(&tanggal)
	fmt.Print("Masukkan bulan  : ")
	fmt.Scan(&bulan)
	fmt.Print("Masukkan tahun  : ")
	fmt.Scan(&tahun)

	nJadwal = buatDataJadwalKosong(nLapangan, nSewa, tanggal, bulan, tahun)

	fmt.Println()
	fmt.Println("Urutkan jadwal:")
	fmt.Println("1. Tanpa sorting")
	fmt.Println("2. Selection Sort berdasarkan jam mulai")
	fmt.Println("3. Insertion Sort berdasarkan harga")
	fmt.Print("Pilih: ")
	fmt.Scan(&pilihanSort)

	if pilihanSort == 2 || pilihanSort == 3 {
		fmt.Println("1. Ascending")
		fmt.Println("2. Descending")
		fmt.Print("Pilih urutan: ")
		fmt.Scan(&urutan)

		if urutan == 2 {
			asc = false
		} else {
			asc = true
		}
	}

	if pilihanSort == 2 {
		selectionSortJadwalByJam(nJadwal, asc)
	} else if pilihanSort == 3 {
		insertionSortJadwalByHarga(nJadwal, asc)
	}

	tampilDataJadwalKosong(nJadwal)
}

func menuStatistik(nSewa int) {
	var bulan int

	fmt.Println()
	fmt.Println("========================================")
	fmt.Println("               STATISTIK")
	fmt.Println("========================================")

	if nSewa > 0 {
		fmt.Print("Masukkan Bulan : ")
		fmt.Scan(&bulan)

		fmt.Println("Total Pendapatan :", totalPendapatanBulanan(nSewa, bulan))
		fmt.Println("Jam Paling Ramai :", jamPalingSeringDipesan(nSewa))
	} else {
		fmt.Println("Belum ada data sewa untuk dihitung statistik.")
	}
}

func menuUtama(nLapangan *int, nPenyewa *int, nSewa *int) {
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
			menuTransaksi(nSewa, *nLapangan, *nPenyewa)
		} else if pilihan == 4 {
			menuCariPenyewa(*nPenyewa)
		} else if pilihan == 5 {
			menuJadwalKosong(*nLapangan, *nSewa)
		} else if pilihan == 6 {
			menuStatistik(*nSewa)
		} else if pilihan == 0 {
			simpanDataLapangan(*nLapangan)
			simpanDataPenyewa(*nPenyewa)
			simpanDataSewa(*nSewa)
			fmt.Println("Terima kasih telah menggunakan Futsal-Book.")
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func main() {
	var nLapangan int = 0
	var nPenyewa int = 0
	var nSewa int = 0

	bacaDataLapangan(&nLapangan)
	bacaDataPenyewa(&nPenyewa)
	bacaDataSewa(&nSewa)

	menuUtama(&nLapangan, &nPenyewa, &nSewa)
}
