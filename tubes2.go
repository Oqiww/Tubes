package main

import "fmt"

//tipe bentukan untuk produk
type produk struct {
	nama  string
	harga int
}

//tipe bentukan untuk tim pada perusahaan
type tim struct {
	nama, peran string
}

//tipe bentukan perusahaan
type perusahaan struct {
	barang                                                             [100]produk
	orang                                                              [100]tim
	tanggal, pengeluaran, pemasukan, income, jumlahOrang, banyakBarang int
	namaPerusahaan, lokasi, bidang                                     string
	ownerID                                                            int
}

//tipe bentukan untuk owner
type owner struct {
	namaOwner string
	pw        string
}

//tipe bentukan untuk investor
type investor struct {
	namaInvestor string
	pw           string
}

//array dengan maksimal 100 dari perusahaan
type tabPerusahaan [100]perusahaan

//array dengan maksimal 100 dari investor
type tabInvestor [100]investor

//array dengan maksimal 100 dari owner
type tabOwner [100]owner

//prosedur untuk menampilkan tampilan awal pada program
//Fitur: menampilkan tampilan awal program pada User untuk memilih ataupun keluar aplikasi
func menu() {
	fmt.Println("╔═══════════════════════════════════╗")
	fmt.Println("║        Halo Selamat Datang        ║")
	fmt.Println("╠═══════════════════════════════════╣")
	fmt.Println("║           Daftar Pengguna         ║")
	fmt.Println("╚═══════════════════════════════════╝")
	fmt.Println("1. Owner")
	fmt.Println("2. Investor")
	fmt.Println("3. Keluar aplikasi")

}

//prosedur untuk menampilkan opsi menu pada owner maupun investor
//Fitur: memberikan User opsi untuk mendaftarkan akun baru atau masuk dengan akun yang sudah terdaftar
func menuLogin() {
	fmt.Println("╔═══════════════════════════════════╗")
	fmt.Println("║        Halo Selamat Datang        ║")
	fmt.Println("╠═══════════════════════════════════╣")
	fmt.Println("║             Pilih Opsi            ║")
	fmt.Println("╚═══════════════════════════════════╝")
	fmt.Println("1. Sign Up")
	fmt.Println("2. Sign In")
	fmt.Println("3. Kembali ke menu ")

}

//prosedur untuk menampilkan menu investor
//Fitur: menyediakan navigasi bagi investor untuk melihat, mengurutkan, dan mencari data perusahaan
func menuInvestor(investor tabInvestor, data tabPerusahaan, nPerusahaan int) {
	var pilih int
	aman := true
	for aman {
		fmt.Println("╔═══════════════════════════════════╗")
		fmt.Println("║        Halo Selamat Datang        ║")
		fmt.Println("╠═══════════════════════════════════╣")
		fmt.Println("║             Pilih Opsi            ║")
		fmt.Println("╚═══════════════════════════════════╝")
		fmt.Println("1. Tampilkan perusahaan ")
		fmt.Println("2. Urutkan perusahaan ")
		fmt.Println("3. Cari perusahaan")
		fmt.Println("4. Kembali ke menu ")

		fmt.Print("Silahkan pilih menu: ")
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			cekDataPerusahaan(data, nPerusahaan)
			detailPerushaan(data, nPerusahaan)
		case 2:
			urutkanPerusahaan(&data, nPerusahaan)
		case 3:
			cariPerusahaan(data, nPerusahaan)
		case 4:
			return
		default:
			fmt.Println("Masukkan tidak valid")
		}
	}

}

//prosedur untuk menampilkan data pada perusahaan yang dipilih (ada opsi dipilih maupun tidak)
func detailPerushaan(data tabPerusahaan, nPerushaan int) {
	var pilih string
	var nomor int
	if nPerushaan == 0 {
		return
	}
	for {
		fmt.Print("Apakah anda ingin mengecek data perusahaan? ")
		fmt.Scan(&pilih)
		switch pilih {
		case "Ya":
			fmt.Print("Silahkan pilih perusahaan: ")
			fmt.Scan(&nomor)
			if nomor < 1 || nomor > nPerushaan {
				fmt.Println("Nomor perusahaan tidak valid")
			} else {
				fmt.Println("1. Nama perusahaan: ", data[nomor-1].namaPerusahaan)
				fmt.Println("1. Tanggal berdiri perusahaan: ", data[nomor-1].tanggal)
				fmt.Println("3. Lokasi perusahaan: ", data[nomor-1].lokasi)
				fmt.Println("4. Pemasukkan perusahaan: ", data[nomor-1].pemasukan)
				fmt.Println("5. Pengeluaran perusahaan: ", data[nomor-1].pengeluaran)
				fmt.Println("6. Laba perusahaan: ", data[nomor-1].income)
				fmt.Print("7. Bidang perusahaan: ", data[nomor-1].bidang)
				fmt.Println()
				if data[nomor-1].jumlahOrang == 0 {
					fmt.Print("8. Tim: -")
					fmt.Println()
				} else {
					fmt.Println("8. Berikut adalah data tim perusahaan ", data[nomor-1].namaPerusahaan)
					for i := 0; i < data[nomor-1].jumlahOrang; i++ {
						fmt.Printf("%-10s %s\n", "- Nama:", data[nomor-1].orang[i].nama)
						fmt.Printf("%-10s %s\n", "  Peran:", data[nomor-1].orang[i].peran)
					}
				}
				if data[nomor-1].banyakBarang == 0 {
					fmt.Print("9. Produk: -")
					fmt.Println()
				} else {
					fmt.Println("9. Berikut adalah daftar produk perusahaan ", data[nomor-1].namaPerusahaan)
					for i := 0; i < data[nomor-1].banyakBarang; i++ {
						fmt.Printf("%3d. nama: %-25s harga: %12d\n", i+1, data[nomor-1].barang[i].nama, data[nomor-1].barang[i].harga)
						fmt.Println()
					}
				}
			}
		case "Tidak":
			return
		default:
			fmt.Println("Maaf pilihan tidak valid")
		}
	}

}

//prosedur untuk menampilkan menu owner
func menuOwner(owner tabOwner, data *tabPerusahaan, nPerusahaan *int, ownerID int) {
	var pilih int
	aman := true
	for aman {
		fmt.Println("╔═══════════════════════════════════╗")
		fmt.Println("║        Halo Selamat Datang        ║")
		fmt.Println("╠═══════════════════════════════════╣")
		fmt.Println("║             Pilih Opsi            ║")
		fmt.Println("╚═══════════════════════════════════╝")
		fmt.Println("1. Isi data perusahaan ")
		fmt.Println("2. Masukkan produk ")
		fmt.Println("3. Hapus perusahaan ")
		fmt.Println("4. Ubah data perusahaan")
		fmt.Println("5. Kembali ke menu ")

		fmt.Print("Silahkan pilih menu: ")
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			isiDataPerusahaan(data, nPerusahaan, ownerID)
		case 2:
			if *nPerusahaan == 0 {
				fmt.Println("Belum ada perusahaan yang terdaftar, silahkan daftarkan perusahaan terlebih dahulu")
			} else {
				cekDataPerusahaan(*data, *nPerusahaan)
				var pilihPerusahaan int
				fmt.Print("Silahkan pilih perusahaan untuk dimasukkan produknya: ")
				fmt.Scan(&pilihPerusahaan)
				if pilihPerusahaan < 1 || pilihPerusahaan > *nPerusahaan {
					fmt.Println("Nomor perusahaan tidak valid")
				} else {
					produkPerusahaan(&data[pilihPerusahaan-1], ownerID)
				}
			}
		case 3:
			cekDataPerusahaan(*data, *nPerusahaan)
			hapusDataPerusahaan(*&data, *&nPerusahaan, ownerID)
		case 4:
			cekDataPerusahaan(*data, *nPerusahaan)
			if *nPerusahaan != 0 {
				var mau int
				fmt.Print("Silahkan pilih perusahaan yang ingin diubah: ")
				fmt.Scan(&mau)
				if mau > *nPerusahaan {
					fmt.Println("Maaf pilihan tidak valid")
				} else {
					ubahDataPerusahaan(&data[mau-1], *nPerusahaan, ownerID)
				}
			}
		case 5:
			return
		default:
			fmt.Println("Maaf pilihan anda tidak valid")
		}
	}

}

//fungsi main utama pada program
func main() {
	var pilihan, login, nOwner, nInvestor, nPerusahaan int
	var dataOwner tabOwner
	var dataInvestor tabInvestor
	var dataPerusahaan tabPerusahaan
	aman := true
	for aman {
		menu()
		fmt.Print("Pilih: ")
		fmt.Scan(&pilihan)
		switch pilihan {
		case 1:
			ownerLogin := false
			for !ownerLogin {
				menuLogin()
				fmt.Print("Silahkan pilih: ")
				fmt.Scan(&login)
				switch login {
				case 1:
					signUpOwner(&dataOwner, &nOwner)
				case 2:
					ownerID := loginOwner(dataOwner, nOwner)
					if ownerID != -1 {
						menuOwner(dataOwner, &dataPerusahaan, &nPerusahaan, ownerID)
						ownerLogin = true
					}
				case 3:
					ownerLogin = true
				default:
					fmt.Println("Pilihan tidak valid")
				}
			}

		case 2:
			investorLogin := false
			for !investorLogin {
				menuLogin()
				fmt.Print("Silahkan pilih: ")
				fmt.Scan(&login)
				switch login {
				case 1:
					signUpInvestor(&dataInvestor, &nInvestor)
				case 2:
					index := loginInvestor(dataInvestor, nInvestor)
					if index != -1 {
						menuInvestor(dataInvestor, dataPerusahaan, nPerusahaan)
						investorLogin = true
					}
				case 3:
					investorLogin = true
				default:
					fmt.Println("Pilihan tidak valid")

				}
			}

		case 3:
			fmt.Println("Terimakasih telah memakai aplikasi kami")
			return

		default:
			fmt.Println("Maaf pilihan tidak valid")
		}

	}

}

//prosedur untuk owner melakukan signup akun
func signUpOwner(I *tabOwner, nOwner *int) {
	var nama, pass string
	aman := true
	fmt.Print("Masukkan username anda: ")
	fmt.Scan(&nama)

	fmt.Print("Masukkan password anda: ")
	fmt.Scan(&pass)

	i := 0
	for i < *nOwner && aman == true {
		if nama == I[i].namaOwner {
			aman = false
		}
		i++
	}

	if *nOwner < 100 && aman {
		I[*nOwner].namaOwner = nama
		I[*nOwner].pw = pass
		*nOwner = *nOwner + 1
		fmt.Println("Sign up berhasil silahkan login kembali ")
	} else {
		fmt.Println("Jumlah owner sudah penuh atau username sudah terpakai")
	}
}

//fungsi untuk owner melakukan login owner dengan sequential search dan mengembalikan index owner
func loginOwner(I tabOwner, nOwner int) int {
	var nama, pass string
	if nOwner == 0 {
		fmt.Println("Maaf username anda tidak ditemukan")
		return -1
	}

	for {
		fmt.Print("Masukkan username anda: ")
		fmt.Scan(&nama)

		fmt.Print("Masukkan password anda: ")
		fmt.Scan(&pass)

		for i := 0; i < nOwner; i++ {
			if nama == I[i].namaOwner && pass == I[i].pw {
				return i
			}
		}
		fmt.Println("Username/password salah silahkan coba lagi")

	}
}

//prosedur untuk inverstor melakukan sign up
func signUpInvestor(I *tabInvestor, nInvestor *int) {
	var nama, pass string
	aman := true
	fmt.Print("Masukkan username anda: ")
	fmt.Scan(&nama)

	fmt.Print("Masukkan password anda: ")
	fmt.Scan(&pass)

	i := 0
	for i < *nInvestor && aman == true {
		if nama == I[i].namaInvestor {
			aman = false
		}
		i++
	}

	if *nInvestor < 100 && aman {
		I[*nInvestor].namaInvestor = nama
		I[*nInvestor].pw = pass
		*nInvestor = *nInvestor + 1
		fmt.Println("Sign up berhasil silahkan login kembali ")
	} else {
		fmt.Println("Jumlah owner sudah penuh atau username sudah terpakai")
	}
}

//fungsi investor untuk melakukan login invsestor dengan sequential search dan mengembalikan index investor
func loginInvestor(I tabInvestor, nInvestor int) int {
	var nama, pass string

	if nInvestor == 0 {
		fmt.Println("Maaf username anda tidak ditemukan")
		return -1
	}

	for {
		fmt.Print("Masukkan username anda: ")
		fmt.Scan(&nama)

		fmt.Print("Masukkan password anda: ")
		fmt.Scan(&pass)

		for i := 0; i < nInvestor; i++ {
			if nama == I[i].namaInvestor && pass == I[i].pw {
				return i
			}
		}
		fmt.Println("Username/password salah silahkan coba lagi")
	}
}

//prosedur untuk menghapus data perusahaan
func hapusDataPerusahaan(P *tabPerusahaan, nPerusahaan *int, ownerID int) {
	if *nPerusahaan == 0 {
		fmt.Println("Maaf belum ada perusahaan anda yang terdaftar")
		return
	}
	var pilih int
	fmt.Print("Silahkan pilih perusahaan yang ingin anda hapus: ")
	fmt.Scan(&pilih)

	if pilih > *nPerusahaan {
		fmt.Println("Pilihan tidak valid")
		return
	}

	i := pilih - 1

	if P[pilih-1].ownerID != ownerID {
		fmt.Println("Maaf anda bukan pemilik perusahaan ini")
		return
	}

	for j := i; j < *nPerusahaan; j++ {
		P[j] = P[j+1]
	}

	*nPerusahaan = *nPerusahaan - 1
	fmt.Println("Data telah berhasil dihapus")

}

//prosedur untuk mengisi data perusahaan
func isiDataPerusahaan(P *tabPerusahaan, nPerusahaan *int, ownerID int) {
	var masuk string
	var jumlah int = 0
	fmt.Print("Silahkan isi nama perusahaan anda: ")
	fmt.Scan(&P[*nPerusahaan].namaPerusahaan)

	fmt.Print("Silahkan isi bidang perusahaan anda: ")
	fmt.Scan(&P[*nPerusahaan].bidang)

	fmt.Print("Silahkan masukkan tahun berdiri perusahaan: ")
	fmt.Scan(&P[*nPerusahaan].tanggal)

	fmt.Print("Silahkan masukkan lokasi perusahaan: ")
	fmt.Scan(&P[*nPerusahaan].lokasi)

	fmt.Print("Masukkan pemasukkan perusahaan: ")
	fmt.Scan(&P[*nPerusahaan].pemasukan)

	fmt.Print("Masukkan pengeluaran perusahaan: ")
	fmt.Scan(&P[*nPerusahaan].pengeluaran)
	P[*nPerusahaan].income = P[*nPerusahaan].pemasukan - P[*nPerusahaan].pengeluaran

	P[*nPerusahaan].ownerID = ownerID

	fmt.Print("Apakah anda ingin memasukkan data orang dan peran: ")
	fmt.Scan(&masuk)
	for masuk != "Tidak" {
		fmt.Printf("Masukkan nama orang ke-%d: ", jumlah+1)
		fmt.Scan(&P[*nPerusahaan].orang[jumlah].nama)

		fmt.Printf("Masukkan perang orang ke-%d: ", jumlah+1)
		fmt.Scan(&P[*nPerusahaan].orang[jumlah].peran)

		jumlah++

		fmt.Print("Apakah anda ingin memasukkan data orang dan peran lagi: ")
		fmt.Scan(&masuk)

	}
	P[*nPerusahaan].jumlahOrang = jumlah
	*nPerusahaan = *nPerusahaan + 1
	fmt.Println("Data perusahaan telah berhasil disimpan")

}

//prosedur untuk mengisi produk perusahaan
func produkPerusahaan(P *perusahaan, ownerID int) {
	if P.ownerID != ownerID {
		fmt.Println("Maaf anda bukan pemilik perusahaan ini")
		return
	}

	var banyak int
	fmt.Print("Berapa banyak jumlah produk yang ingin kamu masukkan: ")
	fmt.Scan(&banyak)

	for i := P.banyakBarang; i < banyak+P.banyakBarang; i++ {
		fmt.Printf("Masukkan nama dan harga produk ke-%d: ", i+1)
		fmt.Println()
		fmt.Print("Masukkan nama: ")
		fmt.Scan(&P.barang[i].nama)

		fmt.Print("Masukkan harga: ")
		fmt.Scan(&P.barang[i].harga)
	}
	P.banyakBarang += banyak
	fmt.Println("Data produk telah berhasil disimpan")
}

//prosedur untuk mengecek nama perusahaan yang terdaftar
func cekDataPerusahaan(P tabPerusahaan, nPerusahaan int) {
	if nPerusahaan == 0 {
		fmt.Println("Data perusahaan belum tersedia")
		return
	}

	fmt.Println("╔═══════════════════════════════════╗")
	fmt.Println("║          Daftar Perusahaan        ║")
	fmt.Println("╚═══════════════════════════════════╝")
	for i := 0; i < nPerusahaan; i++ {
		fmt.Printf("%d. %s\n", i+1, P[i].namaPerusahaan)
		fmt.Println()
	}
}

//prosedur untuk mengubah data perusahaan
func ubahDataPerusahaan(P *perusahaan, nPerusahaan int, ownerID int) {
	if P.ownerID != ownerID {
		fmt.Println("Maaf anda bukan pemilik perusahaan ini")
		return
	}

	if nPerusahaan == 0 {
		return
	}

	var jumlah int
	var masuk string
	fmt.Print("Silahkan isi nama perusahaan anda: ")
	fmt.Scan(&P.namaPerusahaan)

	fmt.Print("Silahkan isi bidang perusahaan anda: ")
	fmt.Scan(&P.bidang)

	fmt.Print("Silahkan masukkan tahun berdiri perusahaan: ")
	fmt.Scan(&P.tanggal)

	fmt.Print("Silahkan masukkan lokasi perusahaan: ")
	fmt.Scan(&P.lokasi)

	fmt.Print("Masukkan pemasukkan perusahaan: ")
	fmt.Scan(&P.pemasukan)

	fmt.Print("Masukkan pengeluaran perusahaan: ")
	fmt.Scan(&P.pengeluaran)
	P.income = P.pemasukan - P.pengeluaran

	fmt.Print("Apakah anda ingin memasukkan data orang dan peran: ")
	fmt.Scan(&masuk)
	for masuk != "Tidak" {
		fmt.Printf("Masukkan nama orang ke-%d: ", jumlah+1)
		fmt.Scan(&P.orang[jumlah].nama)

		fmt.Printf("Masukkan perang orang ke-%d: ", jumlah+1)
		fmt.Scan(&P.orang[jumlah].peran)

		jumlah++

		fmt.Print("Apakah anda ingin memasukkan data orang dan peran lagi: ")
		fmt.Scan(&masuk)

	}
	P.jumlahOrang = jumlah
	fmt.Println("Data telah berhasil diubah")

}

//prosedur untuk menampilkan menu urutkanperusahaan
func urutkanPerusahaan(P *tabPerusahaan, nPerusahaan int) {
	if nPerusahaan == 0 {
		fmt.Println("Data perusahaan belum tersedia")
		return
	}

	var pilih int
	aman := true
	for aman {
		fmt.Println("╔═══════════════════════════════════╗")
		fmt.Println("║             Pilih Opsi            ║")
		fmt.Println("╚═══════════════════════════════════╝")
		fmt.Println("1. Pemasukkan dari yang terbesar")
		fmt.Println("2. Pemasukkan dari yang terkecil")
		fmt.Println("3. Tahun berdiri dari yang paling tua")
		fmt.Println("4. Tahun berdiri dari yang paling baru")
		fmt.Println("5. Pemasukkan terbesar")
		fmt.Println("6. Pemasukkan terkecil")
		fmt.Println("7. Kembali ke menu")

		fmt.Print("Silahkan masukkan pilihan anda: ")
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			urutkanTerbesar(P, nPerusahaan)
		case 2:
			urutkanTerkecil(P, nPerusahaan)
		case 3:
			urutkanTahunTerlama(P, nPerusahaan)
		case 4:
			urutkanTahunTermuda(P, nPerusahaan)
		case 5:
			pemasukkanTerbesar(*P, nPerusahaan)
		case 6:
			pemasukkanTerkecil(*P, nPerusahaan)
		case 7:
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

//prosedur untuk mengurutkan perusahaan dan menampilkan nama dan tanggal berdiri perusahaan berdasarkan tahun secara ASC dengan insertion sort
func urutkanTahunTermuda(P *tabPerusahaan, nPerusahaan int) {
	for i := 1; i < nPerusahaan; i++ {
		j := i
		for j > 0 && P[j-1].tanggal < P[j].tanggal {
			P[j], P[j-1] = P[j-1], P[j]
			j--
		}
	}
	for i := 0; i < nPerusahaan; i++ {
		fmt.Printf("%d. %s: %d\n", i+1, P[i].namaPerusahaan, P[i].tanggal)
		fmt.Println()
	}

}

////prosedur untuk mengurutkan perusahaan dan menampilkan nama dan tanggal berdiri perusahaan berdasarkan tahun secara DESC dengan insertion sort
func urutkanTahunTerlama(P *tabPerusahaan, nPerusahaan int) {
	for i := 1; i < nPerusahaan; i++ {
		j := i
		for j > 0 && P[j-1].tanggal > P[j].tanggal {
			P[j], P[j-1] = P[j-1], P[j]
			j--
		}
	}
	for i := 0; i < nPerusahaan; i++ {
		fmt.Printf("%d. %s: %d\n", i+1, P[i].namaPerusahaan, P[i].tanggal)
		fmt.Println()
	}
}

////prosedur untuk mengurutkan perusahaan dan menampilkan nama dan pemasukan perusahaan berdasarkan pemasukan secara DESC dengan insertion sort
func urutkanTerbesar(P *tabPerusahaan, nPerusahaan int) {
	for i := 0; i < nPerusahaan; i++ {
		max := i
		for j := i + 1; j < nPerusahaan; j++ {
			if P[i].pemasukan < P[j].pemasukan {
				max = j
			}
		}
		P[i], P[max] = P[max], P[i]
	}

	for i := 0; i < nPerusahaan; i++ {
		fmt.Printf("%d. %s: %d\n", i+1, P[i].namaPerusahaan, P[i].pemasukan)
		fmt.Println()
	}
}

//////prosedur untuk mengurutkan perusahaan dan menampilkan nama dan pemasukan perusahaan berdasarkan pemasukan secara ASC dengan insertion sort
func urutkanTerkecil(P *tabPerusahaan, nPerusahaan int) {
	for i := 0; i < nPerusahaan; i++ {
		min := i
		for j := i + 1; j < nPerusahaan; j++ {
			if P[i].pemasukan > P[j].pemasukan {
				min = j
			}
		}
		P[i], P[min] = P[min], P[i]
	}

	for i := 0; i < nPerusahaan; i++ {
		fmt.Printf("%d. %s: %d\n", i+1, P[i].namaPerusahaan, P[i].pemasukan)
		fmt.Println()
	}
}

//prosedur untuk menampilkan pemasukan terbesar perusahaan dengan findmax
func pemasukkanTerbesar(P tabPerusahaan, nPerusahaan int) {
	max := 0
	for i := 1; i < nPerusahaan; i++ {
		if P[i].pemasukan > P[max].pemasukan {
			max = i
		}
	}
	fmt.Println("Perusahaan dengan pemasukan terbesar: ")
	fmt.Println(P[max].namaPerusahaan, P[max].pemasukan)
}

////prosedur untuk menampilkan pemasukan terbesar perusahaan dengan findmin
func pemasukkanTerkecil(P tabPerusahaan, nPerusahaan int) {
	min := 0
	for i := 1; i < nPerusahaan; i++ {
		if P[i].pemasukan < P[min].pemasukan {
			min = i
		}
	}
	fmt.Println("Perusahaan dengan pemasukan terkecil: ")
	fmt.Println(P[min].namaPerusahaan, P[min].pemasukan)
}

//prodesur untuk menampilkna menu untuk mencari perusahaan
func cariPerusahaan(P tabPerusahaan, nPerusahaan int) {
	if nPerusahaan == 0 {
		fmt.Println("Data perusahaan belum tersedia")
		return
	}
	aman := true
	for aman {
		fmt.Println("╔═══════════════════════════════════╗")
		fmt.Println("║             Pilih Opsi            ║")
		fmt.Println("╚═══════════════════════════════════╝")
		fmt.Println("1. Cari nama perusahaan")
		fmt.Println("2. Cari bidang perusahaan")
		fmt.Println("3. Kembali ke menu")
		fmt.Print("Silahkan pilih: ")
		var pilih int
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			var mana string
			fmt.Print("Silahkan cari nama perusahaan yang anda inginkan: ")
			fmt.Scan(&mana)
			index := cariNamaPerusahaan(&P, nPerusahaan, mana)
			if index != -1 {
				fmt.Println("1. Nama perusahaan: ", P[index].namaPerusahaan)
				fmt.Println("1. Tanggal berdiri perusahaan: ", P[index].tanggal)
				fmt.Println("3. Lokasi perusahaan: ", P[index].lokasi)
				fmt.Println("4. Pemasukkan perusahaan: ", P[index].pemasukan)
				fmt.Println("5. Pengeluaran perusahaan: ", P[index].pengeluaran)
				fmt.Println("6. Laba perusahaan: ", P[index].income)
				fmt.Print("7. Bidang perusahaan: ", P[index].bidang)
				fmt.Println()
				if P[index].jumlahOrang == 0 {
					fmt.Print("8. Tim: -")
					fmt.Println()
				} else {
					fmt.Println("8. Berikut adalah data tim perusahaan ", P[index].namaPerusahaan)
					for i := 0; i < P[index].jumlahOrang; i++ {
						fmt.Println("- Nama: ", P[index].orang[i].nama)
						fmt.Println("Peran: ", P[index].orang[i].peran)
					}
				}
				if P[index].banyakBarang == 0 {
					fmt.Print("9. Produk: -")
					fmt.Println()
				} else {
					fmt.Println("9. Berikut adalah daftar produk perusahaan ", P[index].namaPerusahaan)
					for i := 0; i < P[index].banyakBarang; i++ {
						fmt.Printf("%d. nama: %s harga: %d\n", i+1, P[index].barang[i].nama, P[index].barang[i].harga)
						fmt.Println()
					}
				}
			} else {
				fmt.Println("Maaf perusahaan tidak ditemukan")
			}
		case 2:
			cariBidangPerusahaan(P, nPerusahaan)
		case 3:
			return
		default:
			fmt.Println("Maaf pilihan anda tidak valid")
		}
	}

}

//prosedur untuk mencari nama perusahaan dengan algoritma binary search
func cariNamaPerusahaan(P *tabPerusahaan, nPerusahaan int, x string) int {
	urutkanNama(P, nPerusahaan)
	low := 0
	high := nPerusahaan - 1
	for low <= high {
		mid := (low + high) / 2
		if P[mid].namaPerusahaan == x {
			return mid
		} else if P[mid].namaPerusahaan < x {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

//prosedur untuk mengurutkan perusahaan berdasarkan nama dengan algoritma selection sort
func urutkanNama(P *tabPerusahaan, nPerusahaan int) {
	for i := 0; i < nPerusahaan; i++ {
		min := i
		for j := i + 1; j < nPerusahaan; j++ {
			if P[i].namaPerusahaan > P[j].namaPerusahaan {
				min = j
			}
		}
		P[i], P[min] = P[min], P[i]
	}

}

//prosedur untuk mencari bidang perusahaan dengan algoritma sequential search
func cariBidangPerusahaan(P tabPerusahaan, nPerusahaan int) {
	var nomorUrut int
	var pilih string
	ada := false
	fmt.Print("Silahkan pilih bidang perusahaan yang anda inginkan: ")
	fmt.Scan(&pilih)

	for i := 0; i < nPerusahaan; i++ {
		if P[i].bidang == pilih {
			ada = true
			nomorUrut++
			fmt.Printf("%d. %s\n", nomorUrut, P[i].namaPerusahaan)
			fmt.Println()
		}
	}
	if !ada {
		fmt.Println("Maaf bidang perusahaan tidak ditemukan")
	}
	fmt.Printf("Jumlah perusahaan berdasarkan bidang %s: %d ", pilih, nomorUrut)
	fmt.Println()
}
