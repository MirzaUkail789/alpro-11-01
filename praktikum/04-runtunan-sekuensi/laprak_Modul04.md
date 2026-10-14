# <h1 align="center">Laporan Praktikum Modul [Nomor Modul] - Runtunan atau Sekuensi</h1>
<p align="center">Mirza Ukail Falah Triyarso - 109092600006</p>

## Dasar Teori

### A. Pengertian Runtunan atau Sekuensi
Runtunan, atau sering disebut sekuensi, memiliki arti yang sama dengan “berurutan”. Dalam dunia pemrograman, setiap instruksi yang ditulis dengan bahasa pemrograman apa pun akan dijalankan komputer secara bertahap, mulai dari instruksi paling awal sampai yang paling akhir, yaitu dari atas ke bawah dan dari kiri ke kanan. Karena itu, memahami cara menyusun urutan instruksi dengan benar adalah hal yang sangat penting. Jika urutan baris kode yang kita tulis keliru, hasil akhir program pun bisa ikut berubah dan tidak sesuai harapan.

## Guided

### 1. grade.go

```go
package main
import "bufio"
import "fmt"
import "os"

func main() {
	var (
		nilai float64
		nama,grade string
	)
	// bufio scanner untuk input
	scanner := bufio.NewScanner(os.Stdin)
	
	fmt.Println("Masukkan nama: ")
	scanner.Scan()
	nama = scanner.Text()

	fmt.Println("Masukkan nilai: ")
	fmt.Scan(&nilai)

	if nilai >= 90 && nilai <= 100 {
		grade = "A"
	} else if nilai >= 80 && nilai < 90 {
		grade = "B"
	} else if nilai >= 70 && nilai < 80 {
		grade = "C"
	} else if nilai >= 60 && nilai < 70 {
		grade = "D"
	} else {
		grade = "E"
	}

	fmt.Println(nama)
	fmt.Println(grade)
}
}
```
### Deskripsi
Program di atas adalah aplikasi berbasis teks sederhana yang digunakan untuk mengklasifikasikan atau menentukan grade (huruf mutu) nilai seorang siswa berdasarkan rentang nilai angka yang dimasukkan.

#### Deklarasi Variabel:

- Variabel nilai dengan tipe data float64 untuk menyimpan angka nilai siswa.

- Variabel nama dan grade dengan tipe data string untuk menyimpan nama siswa dan hasil huruf mutu.
 
#### Pengambilan Input Pengguna (bufio dan fmt):

-  Program menggunakan bufio.NewScanner(os.Stdin) untuk membaca input teks nama siswa. Penggunaan bufio ini memungkinkan program untuk menerima input nama yang terdiri dari lebih dari satu kata (misalnya nama lengkap dengan spasi).

- Program kemudian meminta input angka nilai siswa menggunakan perintah fmt.Scan(&nilai).

#### Logika Percabangan (if-else):
Program menyeleksi nilai yang dimasukkan ke dalam beberapa rentang untuk menentukan grade yang sesuai:

- Jika nilai di antara 90 hingga 100, maka grade = "A"

- Jika nilai di antara 80 hingga kurang dari 90, maka grade = "B"

- Jika nilai di antara 70 hingga kurang dari 80, maka grade = "C"

- Jika nilai di antara 60 hingga kurang dari 70, maka grade = "D"

- Jika nilai kurang dari 60, maka grade = "E"

#### Menampilkan Output:
Terakhir, program mencetak nama siswa serta grade yang didapat ke layar menggunakan perintah fmt.Println().

### 2. klasifikasi.go

```go
package main
import "fmt"

func main() {
	var (
		usia,gajiTahunan int
	)
	fmt.Print("Masukkan usia: ")
	fmt.Scan(&usia)

	fmt.Print("Masukkan gaji tahunan (juta): ")
	fmt.Scan(&gajiTahunan)

	switch {
		case usia < 18 && gajiTahunan == 0 :
			fmt.Println("Masih Sekolah")
		case (usia >= 18 && usia <= 25) && gajiTahunan >= 50 :
			fmt.Println("Muda Sukses")
		case (usia >= 18 && usia <= 25) && gajiTahunan < 50 :
			fmt.Println("Masih belajar hidup")
		case (usia >= 26 && usia <= 40) && gajiTahunan >= 100 :
			fmt.Println("Pekerja Mapan")
		case (usia >= 26 && usia <= 40) && gajiTahunan < 100 :
			fmt.Println("Perlu perbaikan karir")
		case usia > 40 && gajiTahunan >= 150 :
			fmt.Println("Profesional Berpengalaman")
		case usia > 40 && gajiTahunan < 150 :
			fmt.Println("Perlu Evaluasi Finansial")
	}
}
```
#### Deskripsi
Program di atas adalah aplikasi konsol sederhana yang digunakan untuk mengklasifikasikan status atau kategori kehidupan seseorang berdasarkan dua parameter utama, yaitu usia (dalam tahun) dan gaji tahunan (dalam satuan juta rupiah).

Berikut adalah tahapan atau alur kerja dari program tersebut:

Deklarasi Variabel:

Variabel usia dan gajiTahunan dideklarasikan dengan tipe data int (bilangan bulat) untuk menampung masukan angka dari pengguna.

Pengambilan Input Pengguna:

Program meminta pengguna memasukkan nilai usia menggunakan perintah fmt.Scan(&usia).

Program kemudian meminta pengguna memasukkan nilai gaji tahunan dalam juta rupiah menggunakan perintah fmt.Scan(&gajiTahunan).

Logika Seleksi Kondisi (switch-case tanpa ekspresi):
Program menggunakan struktur switch dengan kondisi bercabang (boolean expressions) untuk mencocokkan kombinasi usia dan gaji tahunan ke dalam beberapa kategori status berikut:

"Masih Sekolah": Jika usia kurang dari 18 tahun dan gaji tahunan sama dengan 0.

"Muda Sukses": Jika usia berada di rentang 18 sampai 25 tahun dan gaji tahunan mencapai atau lebih dari 50 juta.

"Masih belajar hidup": Jika usia berada di rentang 18 sampai 25 tahun namun gaji tahunan kurang dari 50 juta.

"Pekerja Mapan": Jika usia berada di rentang 26 sampai 40 tahun dan gaji tahunan mencapai atau lebih dari 100 juta.

"Perlu perbaikan karir": Jika usia berada di rentang 26 sampai 40 tahun namun gaji tahunan kurang dari 100 juta.

"Profesional Berpengalaman": Jika usia di atas 40 tahun dan gaji tahunan mencapai atau lebih dari 150 juta.

"Perlu Evaluasi Finansial": Jika usia di atas 40 tahun namun gaji tahunan kurang dari 150 juta.

Menampilkan Output:
Setelah kondisi yang sesuai terpenuhi, program akan langsung mencetak teks status atau kategori yang cocok ke layar menggunakan perintah fmt.Println().

### 3. penilaian.go

```go
package main

import "bufio"
import "fmt"
import "os"

func main() {
	var (
		nama, grade    string
		nilai, pilihan int
	)

	fmt.Println("============= Menu =============")
	fmt.Println("1. Sistem penilaian")
	fmt.Println("0. Keluar")
	fmt.Print("Pilih opsi: ")
	fmt.Scanln(&pilihan)

	if pilihan == 1 {
		fmt.Print("Masukkan Nama:")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		nama = scanner.Text()

		fmt.Print("Masukkan Nilai:")
		fmt.Scan(&nilai)

		if nilai >= 90 && nilai <= 100 {
			grade = "A"
		} else if nilai >= 80 && nilai < 90 {
			grade = "B"
		} else if nilai >= 70 && nilai < 80 {
			grade = "C"
		} else if nilai >= 60 && nilai < 70 {
			grade = "D"
		} else {
			grade = "E"
		}

		fmt.Println(nama)
		fmt.Println(grade)
	} else if pilihan == 0 {
		return
	} else {
		fmt.Println("Pilihan tidak valid")
	}
}
```
#### Deskripsi
Program di atas adalah aplikasi konsol berbasis menu interaktif yang digunakan untuk mengelola sistem penilaian siswa secara kondisional berdasarkan pilihan menu yang dimasukkan oleh pengguna.

Berikut adalah tahapan atau alur kerja dari program tersebut:

Deklarasi Variabel:

Variabel nama dan grade dengan tipe data string untuk menyimpan nama siswa dan huruf mutu.

Variabel nilai dan pilihan dengan tipe data int (bilangan bulat) untuk menampung nilai angka siswa serta nomor opsi menu yang dipilih.

Menampilkan Menu Utama:

Program mencetak pilihan menu ke layar, yaitu opsi 1 untuk masuk ke "Sistem penilaian" dan opsi 0 untuk "Keluar".

Pengguna diminta memasukkan pilihan angka menggunakan perintah fmt.Scanln(&pilihan).

Logika Percabangan Menu Utama (if-else):

Jika pilihan == 1: Program akan meminta pengguna memasukkan nama siswa (menggunakan bufio agar bisa membaca spasi atau lebih dari satu kata) dan memasukkan angka nilai siswa. Nilai tersebut kemudian diseleksi menggunakan percabangan if-else bersarang untuk menentukan grade berdasarkan rentang berikut:

Nilai 90–100 mendapat grade "A"
Nilai 80–89 mendapat grade "B"
Nilai 70–79 mendapat grade "C"
Nilai 60–69 mendapat grade "D"
Di bawah 60 mendapat grade "E"
Terakhir, program mencetak nama dan grade siswa tersebut ke layar.

Jika pilihan == 0: Program akan langsung menghentikan eksekusi (return).

Jika memilih angka selain 1 dan 0: Program akan menampilkan pesan "Pilihan tidak valid".

## Unguided

### 1. [nama_soal, misal: cacahuang]

```go
[Tempelkan kode program di sini, contoh: cacahuang.go]
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](https://github.com/renwxyz/alpro-11-01/blob/main/praktikum/02-bahasa-pemrograman-go/unguided/cacahuang/output.png)


#### Deskripsi
[Tuliskan ringkasan proses praktikum: apa yang dikerjakan, bagian guided dan unguided yang diimplementasikan, serta hasil yang diperoleh.]

### 2. [nama_soal, misal: kalkulator]

```go
[Tempelkan kode program di sini, contoh: kalkulator.go]
```

##### Output
![Screenshot Output Unguided](unguided/[nama_soal]/output.png)

#### Deskripsi
[Tuliskan ringkasan proses praktikum: apa yang dikerjakan, bagian guided dan unguided yang diimplementasikan, serta hasil yang diperoleh.]

<!-- Duplikasi blok "### [nama_soal]" sesuai jumlah folder soal di dalam unguided -->


## Kesimpulan
[Tuliskan kesimpulan yang menjawab tujuan praktikum berdasarkan hasil yang diperoleh.]

## Referensi
1. [Nama Penulis]. ([Tahun]). *[Judul Buku/Sumber]*. [Kota]: [Penerbit]. Diakses pada [tanggal akses] melalui [tautan/DOI]
2. [Nama Penulis]. ([Tahun]). *[Judul Buku/Sumber]*. [Kota]: [Penerbit]. Diakses pada [tanggal akses] melalui [tautan/DOI]
<!-- Tambahkan nomor referensi berikutnya sesuai kebutuhan -->
