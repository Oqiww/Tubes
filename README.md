# Aplikasi Manajemen Startup Sederhana 🚀

[![Made with Go](https://img.shields.io/badge/Made%20with-Go-00ADD8.svg?logo=go&logoColor=white)](https://go.dev/)
Aplikasi ini adalah sistem manajemen data perusahaan berbasis **Command Line Interface (CLI)** yang dirancang secara intuitif menggunakan **Go (Golang)** untuk mempermudah **Owner** (pemilik perusahaan) dan **Investor** (calon penanam modal) dalam mengakses serta mengelola informasi perusahaan.
---

## ✨ Fitur Utama

### 1. Akses Aman: Login & Sign Up
Keamanan data adalah prioritas utama. Setiap pengguna, baik **Owner** maupun **Investor**, **wajib** melakukan proses pendaftaran (`Sign Up`) dan masuk (`Login`) ke akun mereka terlebih dahulu. Ini memastikan penggunaan aplikasi yang lebih aman dan terstruktur.

### 2. Navigasi Intuitif: Menu Utama
Setelah berhasil login, Anda akan disambut dengan `Menu Utama` yang jelas. Menu ini berisi daftar lengkap fitur yang dapat Anda akses, disesuaikan secara otomatis berdasarkan peran Anda (Owner atau Investor). Tersedia juga opsi untuk keluar dari aplikasi kapan saja.

---

## 💼 Fitur untuk Owner

Sebagai Owner perusahaan, Anda memiliki kendali penuh atas informasi dan operasional perusahaan Anda:

### 1. Kelola Data Perusahaan 📝
Tambahkan dan perbarui informasi penting startup Anda dengan mudah:
* **Nama Perusahaan**: Identitas unik startup Anda.
* **Bidang Perusahaan**: Sektor industri tempat startup Anda beroperasi (misal: `Teknologi`, `F&B`, `Kesehatan`).
* **Tahun Berdiri**: Catat jejak sejarah perusahaan (format: `YYYY`, contoh: `2020`).
* **Lokasi Perusahaan**: Alamat fisik atau operasional utama startup Anda.
* **Pemasukan & Pengeluaran**: Pantau kesehatan finansial perusahaan secara rinci.
* **Data Tim Perusahaan**: Tambahkan nama dan peran anggota kunci tim Anda. Informasi ini bersifat **opsional** dan bisa dikosongkan.

### 2. Tambahkan & Kelola Produk 📦
Fitur ini memungkinkan Anda untuk menambahkan produk baru yang ditawarkan oleh perusahaan Anda.
* **Catatan**: Fitur ini hanya dapat digunakan untuk perusahaan yang Anda miliki dan terdaftar di bawah akun Anda. Tidak dapat dijalankan jika belum ada perusahaan yang terdaftar sama sekali.

### 3. Hapus Perusahaan 👋
Memberikan Anda fleksibilitas untuk menghapus data perusahaan yang ingin anda hapus.
* **Catatan**: Fitur ini hanya aktif jika Anda memiliki setidaknya satu perusahaan yang terdaftar.

### 4. Ubah Detail Perusahaan ✏️
Anda dapat dengan leluasa mengubah seluruh data perusahaan yang terdaftar di bawah kepemilikan Anda.
* **Catatan**: Fitur ini memerlukan setidaknya satu perusahaan terdaftar untuk dapat beroperasi.

---

## 📈 Fitur untuk Investor

Sebagai calon penanam modal, Anda membutuhkan akses cepat dan analisis data yang akurat untuk keputusan investasi yang cerdas:

### 1. Jelajahi Daftar Perusahaan 🔍
Lihatlah seluruh daftar nama perusahaan yang terdaftar dalam sistem. Setelah meninjau daftar, Anda dapat memilih nama perusahaan mana pun untuk melihat **detail data lengkapnya**, memberikan gambaran mendalam sebelum berinvestasi.
* **Catatan**: Fitur ini baru berfungsi jika sudah ada perusahaan yang terdaftar dalam sistem.

### 2. Urutkan Perusahaan Sesuai Kebutuhan Analisis Anda 📊
Mempermudah analisis data dengan kemampuan untuk mengurutkan daftar perusahaan berdasarkan berbagai kriteria penting:
* **Berdasarkan Pemasukan**:
    * **Terbesar ke Terkecil (DESC)**: Menampilkan nama perusahaan dan pemasukan, dari yang paling menguntungkan.
    * **Terkecil ke Terbesar (ASC)**: Menampilkan nama perusahaan dan pemasukan, untuk mengidentifikasi potensi pertumbuhan atau tantangan.
* **Berdasarkan Tahun Berdiri**:
    * **Paling Tua ke Paling Baru (ASC)**: Mengurutkan dari startup yang paling berpengalaman.
    * **Paling Baru ke Paling Tua (DESC)**: Menampilkan startup terkini di ekosistem.
* **Pemasukan Teratas & Terbawah**:
    * **Pemasukan Terbesar**: Langsung menunjukkan perusahaan dengan performa finansial tertinggi.
    * **Pemasukan Terkecil**: Mengidentifikasi perusahaan dengan pemasukan terendah untuk studi lebih lanjut.

### 3. Cari Perusahaan Spesifik 🎯
Temukan perusahaan yang Anda inginkan dengan fitur pencarian yang akurat dan efisien:
* **Berdasarkan Nama Perusahaan**: Cukup ketik nama perusahaan, dan sistem akan menampilkan seluruh detailnya jika ditemukan.
* **Berdasarkan Bidang Industri**: Cari perusahaan berdasarkan bidangnya. Sistem akan menampilkan daftar nama perusahaan di bidang tersebut, lengkap dengan **jumlah total perusahaan** yang terdaftar di bidang yang sama, memberikan wawasan pasar.

---

