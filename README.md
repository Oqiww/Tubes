Aplikasi Manajemen Startup Sederhana
Aplikasi ini adalah sistem manajemen data perusahaan berbasis Command Line Interface (CLI) yang dirancang untuk mempermudah Owner (pemilik perusahaan) dan Investor (calon penanam modal) dalam mengakses dan mengelola informasi perusahaan.

Fitur Utama
1. Fitur Login & Sign Up
Setiap pengguna, baik Owner maupun Investor, wajib melakukan sign up dan login terlebih dahulu sebelum dapat mengakses menu utama aplikasi. Ini memastikan keamanan dan personalisasi akses.

2. Fitur Menu Utama
Setelah berhasil login, aplikasi akan menampilkan menu yang berisi daftar fitur yang dapat diakses sesuai dengan peran pengguna (Owner atau Investor). Terdapat juga opsi untuk keluar dari aplikasi.

Fitur untuk Owner
1. Isi Data Perusahaan
Sebagai Owner, Anda dapat menambahkan data perusahaan baru dengan mengisi informasi berikut:

Nama Perusahaan
Bidang Perusahaan
Tanggal Berdiri Perusahaan (format tahun, contoh: 2020)
Lokasi Perusahaan
Pemasukan Perusahaan
Pengeluaran Perusahaan
Data Tim Perusahaan (format: Nama - Peran, bisa kosong/null)
2. Tambahkan Produk
Fitur ini memungkinkan Owner untuk menambahkan produk ke perusahaan yang Anda miliki. Fitur ini hanya berfungsi jika perusahaan tersebut terdaftar di bawah akun Owner yang sedang login dan tidak dapat dijalankan jika belum ada perusahaan yang terdaftar.

3. Hapus Perusahaan
Owner dapat menghapus perusahaan yang Anda miliki dari sistem. Fitur ini tidak dapat dijalankan jika belum ada perusahaan yang terdaftar.

4. Ubah Data Perusahaan
Owner memiliki kemampuan untuk mengubah seluruh data perusahaan yang terdaftar di bawah kepemilikan Anda. Fitur ini tidak dapat dijalankan jika belum ada perusahaan yang terdaftar.

Fitur untuk Investor
1. Tampilkan Perusahaan
Investor dapat melihat seluruh nama perusahaan yang terdaftar dalam sistem. Setelah itu, Anda dapat memilih salah satu nama perusahaan untuk melihat detail data perusahaan tersebut. Fitur ini tidak dapat dijalankan jika belum ada perusahaan yang terdaftar.

2. Urutkan Perusahaan
Fitur ini memungkinkan Investor untuk mengurutkan daftar perusahaan berdasarkan kriteria berikut:

Pemasukan (Terbesar ke Terkecil): Menampilkan nama perusahaan dan pemasukan secara DESC (Descending).
Pemasukan (Terkecil ke Terbesar): Menampilkan nama perusahaan dan pemasukan secara ASC (Ascending).
Tahun Berdiri (Paling Tua ke Paling Baru): Menampilkan nama perusahaan dan tahun berdiri secara DESC (Descending), artinya tahun yang lebih kecil (lebih tua) akan muncul lebih dulu.
Tahun Berdiri (Paling Baru ke Paling Tua): Menampilkan nama perusahaan dan tahun berdiri secara ASC (Ascending), artinya tahun yang lebih besar (lebih baru) akan muncul lebih dulu.
Pemasukan Terbesar: Menampilkan nama perusahaan dan nilai pemasukan dari perusahaan dengan pemasukan tertinggi.
Pemasukan Terkecil: Menampilkan nama perusahaan dan nilai pemasukan dari perusahaan dengan pemasukan terendah.
3. Cari Perusahaan
Fitur ini membantu Investor dalam mencari perusahaan berdasarkan kriteria tertentu:

Nama Perusahaan: Sistem akan menampilkan seluruh data perusahaan jika nama perusahaan yang dicari terdaftar.
Bidang Perusahaan: Sistem akan menampilkan daftar nama perusahaan yang bergerak di bidang yang dipilih. Selain itu, sistem juga akan menjumlahkan berapa banyak perusahaan yang terdaftar di bidang tersebut.
