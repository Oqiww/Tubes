package main

import "fmt"

type produk struct {
	nama  string
	harga int
}
type tim struct {
	nama, peran string
}
type perusahaan struct {
	barang                                                             [100]produk
	orang                                                              [100]tim
	tanggal, pengeluaran, pemasukan, income, jumlahOrang, banyakBarang int
	namaPerusahaan, lokasi                                             string
	ownerID                                                            int
}

type owner struct {
	namaOwner, noTelp, email, tanggal, lokasi string
	pw                                        string
}

type investor struct {
	namaInvestor, email, noTelp string
	pw                          string
}

type tabPerusahaan [100]perusahaan
type tabInvestor [100]investor
type tabOwner [100]owner

func menu() {
	fmt.Println("===== Halo Selamat Datang =====")
	fmt.Println("===== Daftar pengguna =====")
	fmt.Println("1. Owner")
	fmt.Println("2. Investor")
	fmt.Println("3. Keluar aplikasi")

}

func menuLogin() {
	fmt.Println("===== Halo Selamat Datang =====")
	fmt.Println("===== Daftar pengguna =====")
	fmt.Println("1. Sign Up")
	fmt.Println("2. Sign In")
	fmt.Println("3. Kembali ke menu ")

}

func menuInvestor(investor tabInvestor, data tabPerusahaan, nPerusahaan int) {
	var pilih int
	for {
		fmt.Println("===== Halo Selamat Datang =====")
		fmt.Println("===== Daftar pengguna =====")
		fmt.Println("1. Tampilkan perusahaan ")
		fmt.Println("2. Urutkan perusahaan ")
		fmt.Println("3. Kembali ke menu ")

		fmt.Print("Silahkan pilih menu: ")
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			cekDataPerusahaan(data, nPerusahaan)
			detailPerushaan(data, nPerusahaan)
		case 2:
			urutkanPerusahaan(&data, nPerusahaan)
		case 3:
			return
		default:
			fmt.Println("Masukkan tidak valid")
			continue
		}
	}

}

func detailPerushaan(data tabPerusahaan, nPerushaan int) {
	var pilih string
	var nomor int
	for {
		fmt.Print("Apakah anda ingin mengecek data perusahaan? ")
		fmt.Scan(&pilih)
		switch pilih {
		case "Ya":
			fmt.Print("Silahkan pilih perusahaan: ")
			fmt.Scan(&nomor)
			if nomor < 1 || nomor > nPerushaan {
				fmt.Println("Nomor perusahaan tidak valid")
				continue
			}
			fmt.Println("1. Nama perusahaan: ", data[nomor-1].namaPerusahaan)
			fmt.Println("1. Tanggal berdiri perusahaan: ", data[nomor-1].tanggal)
			fmt.Println("3. Lokasi perusahaan: ", data[nomor-1].lokasi)
			fmt.Println("4. Pemasukkan perusahaan: ", data[nomor-1].pemasukan)
			fmt.Println("5. Pengeluaran perusahaan: ", data[nomor-1].pengeluaran)
			fmt.Print("6. Laba perusahaan: ", data[nomor-1].income)
			fmt.Println()
			if data[nomor-1].jumlahOrang == 0 {
				fmt.Print("7. Tim: -")
				fmt.Println()
			} else {
				fmt.Println("7. Berikut adalah data tim perusahaan ", data[nomor-1].namaPerusahaan)
				for i := 0; i < data[nomor-1].jumlahOrang; i++ {
					fmt.Println("- Nama: ", data[nomor-1].orang[i].nama)
					fmt.Println("Peran: ", data[nomor-1].orang[i].peran)
				}
			}
			if data[nomor-1].banyakBarang == 0 {
				fmt.Print("8. Produk: -")
				fmt.Println()
			} else {
				fmt.Println("8. Berikut adalah daftar produk perusahaan ", data[nomor-1].namaPerusahaan)
				for i := 0; i < data[nomor-1].banyakBarang; i++ {
					fmt.Printf("%d. nama: %s harga: %d\n", i+1, data[nomor-1].barang[i].nama, data[nomor-1].barang[i].harga)
					fmt.Println()
				}
			}
		case "Tidak":
			return
		default:
			fmt.Println("Maaf pilihan tidak valid")
		}
	}

}

func menuOwner(owner tabOwner, data *tabPerusahaan, nPerusahaan *int, ownerID int) {
	var pilih int
	for {
		fmt.Println()
		fmt.Println("===== Halo Selamat Datang =====")
		fmt.Println("===== Daftar pengguna: =====")
		fmt.Println("1. Isi data perusahaan ")
		fmt.Println("2. Masukkan produk ")
		fmt.Println("3. Kembali ke menu ")

		fmt.Print("Silahkan pilih menu: ")
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			isiDataPerusahaan(data, nPerusahaan, ownerID)
		case 2:
			if *nPerusahaan == 0 {
				fmt.Println("Belum ada perusahaan yang terdaftar, silahkan daftarkan perusahaan terlebih dahulu")
				continue
			}
			cekDataPerusahaan(*data, *nPerusahaan)
			var pilihPerusahaan int
			fmt.Print("Silahkan pilih perusahaan untuk dimasukkan produknya: ")
			fmt.Scan(&pilihPerusahaan)
			if pilihPerusahaan < 1 || pilihPerusahaan > *nPerusahaan {
				fmt.Println("Nomor perusahaan tidak valid")
				continue
			}
			produkPerusahaan(&data[pilihPerusahaan-1], ownerID)
		case 3:
			return
		}
	}

}

func main() {
	var pilihan, login, nOwner, nInvestor, nPerusahaan int
	var dataOwner tabOwner
	var dataInvestor tabInvestor
	var dataPerusahaan tabPerusahaan
	for {
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
					fmt.Print("Pilihan tidak valid")
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
					fmt.Print("Pilihan tidak valid")

				}
			}

		case 3:
			fmt.Println("Terimakasih telah memakai aplikasi kami")
			return
		}

	}

}

func signUpOwner(I *tabOwner, nOwner *int) {
	var nama, pass string
	aman := true
	fmt.Print("Masukkan username anda: ")
	fmt.Scan(&nama)

	fmt.Print("Masukkan password anda: ")
	fmt.Scan(&pass)

	for i := 0; i < *nOwner; i++ {
		if nama == I[i].namaOwner {
			aman = false
			break
		}
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

func signUpInvestor(I *tabInvestor, nInvestor *int) {
	if *nInvestor < 100 {
		fmt.Print("Masukkan username anda: ")
		fmt.Scan(&I[*nInvestor].namaInvestor)

		fmt.Print("Masukkan password anda: ")
		fmt.Scan(&I[*nInvestor].pw)
		*nInvestor = *nInvestor + 1
		fmt.Println("Sign up berhasil silahkan login kembali ")
	} else {
		fmt.Print("Jumlah investor sudah mencapai batas")
	}
}

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

func isiDataPerusahaan(P *tabPerusahaan, nPerusahaan *int, ownerID int) {
	var masuk string
	var jumlah int = 0
	fmt.Print("Silahkan isi nama perusahaan anda: ")
	fmt.Scan(&P[*nPerusahaan].namaPerusahaan)

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

func cekDataPerusahaan(P tabPerusahaan, nPerusahaan int) {
	if nPerusahaan == 0 {
		fmt.Println("Data perusahaan belum tersedia")
		return
	}

	fmt.Println("==== Daftar perusahaan ====")
	for i := 0; i < nPerusahaan; i++ {
		fmt.Printf("%d. %s\n", i+1, P[i].namaPerusahaan)
		fmt.Println()
	}
}

func urutkanPerusahaan(P *tabPerusahaan, nPerusahaan int) {
	var pilih int
	for {
		fmt.Println()
		fmt.Println("==== SILAHKAN PILIH ====")
		fmt.Println("1. Pemasukkan dari yang terbesar")
		fmt.Println("2. Pemasukkan dari yang terkecil")
		fmt.Println("3. Keuntungan dari yang terbesar")
		fmt.Println("4. Keuntungan dari yang terkecil")
		fmt.Println("5. Tahun berdiri dari yang paling tua")
		fmt.Println("6. Tahun berdiri dari yang paling baru")
		fmt.Println("7. Pemasukkan terbesar")
		fmt.Println("8. Pemasukkan terkecil")
		fmt.Println("9. Kembali ke menu")

		fmt.Print("Silahkan masukkan pilihan anda: ")
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			urutkanPerusahaan(P, nPerusahaan)
		case 7:
			pemasukkanTerbesar(*P, nPerusahaan)
		case 8:
			pemasukkanTerkecil(*P, nPerusahaan)
		case 9:
			return
		default:
			fmt.Println("Pilihan tidak valid")
			continue
		}
	}
}

func urutkanTerbesar(P *tabPerusahaan, nPerusahaan int) {
	for i := 0; i < nPerusahaan-1; i++ {
		for j := i + 1; j < nPerusahaan; j++ {
			if P[i].pemasukan < P[j].pemasukan {
				P[i], P[j] = P[j], P[i]
			}
		}
	}
	for i := 0; i < nPerusahaan; i++ {
		fmt.Printf("%d. %s, %d\n", i+1, P[i].namaPerusahaan, P[i].pemasukan)
		fmt.Println()
	}
}

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
