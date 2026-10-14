# <h1 align="center">Tugas Pendahuluan Modul [Nomor Modul] - [Judul Modul/Topik]</h1>
<p align="center">[Nama Praktikan] - [NIM]</p>

### 1. Sisa Kue

```go
package main

import "fmt"

func main() {
	var y, x int

	fmt.Scan(&y, &x)

	sisaKue := y % x

	fmt.Println(sisaKue)
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/alpro-11-01/praktikum/03-tipe-data-dan-instruksi-dasar/tp/sisa/Output%20sisa.png)


#### Deskripsi
`package main`: Baris ini mendeklarasikan bahwa file kode ini adalah bagian dari paket main, yang menandakan bahwa program ini merupakan aplikasi utama yang dapat dieksekusi secara mandiri.

`import "fmt"`: Mengimpor pustaka standar Format (fmt) milik Go yang berfungsi untuk menangani proses masukan (input) dan keluaran (output) data, seperti membaca ketikan pengguna atau mencetak teks ke layar.

`func main() { ... }`: Fungsi utama (entry point) yang menjadi titik awal bagi sistem runtime Go untuk mulai menjalankan baris-baris perintah di dalam program.

`var y, x int`: Mendeklarasikan dua buah variabel bernama y dan x dengan tipe data bilangan bulat (integer).

`fmt.Scan(&y, &x)`: Berfungsi untuk membaca nilai input yang dimasukkan oleh pengguna melalui terminal/keyboard. Tanda ampersand (&) digunakan untuk mengakses alamat memori dari variabel y dan x agar nilai yang diinputkan dapat disimpan langsung ke dalam variabel tersebut.

`sisaKue := y % x`: Melakukan operasi aritmatika modulus (%) untuk menghitung sisa hasil bagi dari variabel y dibagi dengan x. Hasil perhitungan tersebut kemudian dimasukkan dan disimpan ke dalam variabel baru bernama sisaKue.

`fmt.Println(sisaKue)`: Berfungsi untuk mencetak dan menampilkan nilai akhir dari variabel sisaKue ke layar terminal.

### 2. bool.go

```go
package main

import "fmt"

func main() {
	var status bool
	fmt.Scan(&status)
	fmt.Println(status)
}

```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/alpro-11-01/praktikum/03-tipe-data-dan-instruksi-dasar/tp/bool/Output%20bool.png)


#### Deskripsi
`package main`: Baris ini mendeklarasikan bahwa file kode ini adalah bagian dari paket main, yang menandakan bahwa program ini merupakan aplikasi utama yang dapat dieksekusi secara mandiri.

`import "fmt"`: Mengimpor pustaka standar Format (fmt) milik Go yang berfungsi untuk menangani proses masukan (input) dan keluaran (output) data, seperti membaca ketikan pengguna atau mencetak teks ke layar.

`func main() { ... }`: Fungsi utama (entry point) yang menjadi titik awal bagi sistem runtime Go untuk mulai menjalankan baris-baris perintah di dalam program.

`var y, x int`: Mendeklarasikan dua buah variabel bernama y dan x dengan tipe data bilangan bulat (integer).

`fmt.Scan(&y, &x)`: Berfungsi untuk membaca nilai input yang dimasukkan oleh pengguna melalui terminal/keyboard. Tanda ampersand (&) digunakan untuk mengakses alamat memori dari variabel y dan x agar nilai yang diinputkan dapat disimpan langsung ke dalam variabel tersebut.

`sisaKue` := y % x: Melakukan operasi aritmatika modulus (%) untuk menghitung sisa hasil bagi dari variabel y dibagi dengan x. Hasil perhitungan tersebut kemudian dimasukkan dan disimpan ke dalam variabel baru bernama sisaKue.

`fmt.Println(sisaKue)`: Berfungsi untuk mencetak dan menampilkan nilai akhir dari variabel sisaKue ke layar terminal.


### 3. konversi.go

```go
package main

import "fmt"

func main() {
	var mil float64

	fmt.Scan(&mil)

	kilometer := mil * 1.6

	fmt.Printf("%.1f\n", kilometer)
}

```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/alpro-11-01/praktikum/03-tipe-data-dan-instruksi-dasar/tp/konversi/output%20konversi.png)


#### Deskripsi
`package main`: Mendeklarasikan bahwa file kode ini adalah bagian dari paket main, yang menandakan bahwa program ini merupakan aplikasi utama yang dapat dieksekusi secara mandiri.

`import "fmt"`: Mengimpor pustaka standar Format (fmt) yang digunakan untuk menangani proses masukan (input) dan keluaran (output), seperti membaca data dari terminal atau mencetak teks ke layar.

`func main()`: Fungsi utama (entry point) yang menjadi titik awal bagi sistem runtime Go untuk mulai menjalankan baris-baris perintah di dalam program.

`var mil float64`: Mendeklarasikan sebuah variabel bernama mil dengan tipe data bilangan desimal presisi ganda (floating-point 64-bit).

`fmt.Scan(&mil)`: Berfungsi untuk membaca nilai input desimal yang dimasukkan oleh pengguna melalui terminal/keyboard, lalu menyimpannya ke dalam alamat memori variabel mil.

`kilometer` := mil * 1.6: Melakukan perhitungan konversi dengan mengalikan nilai variabel mil dengan angka 1.6, kemudian mendeklarasikan dan menyimpan hasilnya ke dalam variabel baru bernama kilometer.

`fmt.Printf("%.1f\n", kilometer)`: Berfungsi untuk mencetak hasil akhir variabel kilometer ke layar terminal dengan format tampilan angka desimal yang dibulatkan hingga 1 angka di belakang koma (%.1f) disertai dengan baris baru (\n).

## Kesimpulan
Berdasarkan seluruh rangkaian kegiatan praktikum yang telah dilakukan mengenai penggunaan bahasa pemrograman Go, dapat ditarik kesimpulan sebagai berikut:
    Mahasiswa mampu memahami dan menerapkan struktur dasar program Go yang mencakup penggunaan package main dan fungsi utama func main() sebagai titik awal eksekusi program.   

    Mahasiswa berhasil mendeklarasikan berbagai tipe data (seperti integer, boolean, dan float64) serta menggunakan variabel untuk memproses data masukan secara dinamis menggunakan fungsi fmt.Scan.

    Program-program sederhana yang dibuat—mulai dari perhitungan aritmatika modulus, pengelolaan variabel bertipe boolean, hingga konversi satuan nilai desimal—dapat berjalan dengan baik dan menghasilkan keluaran (output) yang akurat sesuai dengan logika yang telah diprogramkan.