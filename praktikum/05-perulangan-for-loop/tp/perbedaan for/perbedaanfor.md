### Perbedaan Utama
#### 1 `for` dengan Tiga Komponen (Inisialisasi, Kondisi, Post Statement)

- Penggunaan: Digunakan ketika Anda sudah mengetahui berapa kali perulangan harus berjalan, atau saat Anda membutuhkan variabel counter (penghitung) yang naik/turun secara teratur (misalnya dari 1 sampai 10).

- Struktur: Semua variabel pengendali (`i := 0`), batas (`i < 10`), dan perubahan (`i++`) didefinisikan dalam satu baris di dalam header perulangan.

#### 2 `for` dengan Kondisi Saja (Gaya `while`)

- Penggunaan: Digunakan ketika jumlah iterasi tidak pasti atau bergantung pada kondisi eksternal yang berubah di dalam blok kode (misalnya membaca file sampai akhir, menunggu masukan pengguna, atau kondisi logika tertentu).

- Struktur: Inisialisasi variabel dilakukan di luar perulangan, dan post statement (seperti `i++`) diletakkan di dalam blok kode. Di Go, bentuk ini menggantikan fungsi `while` tradisional.

### Contoh Kode Masing-Masing
#### 1. Contoh for dengan Tiga Komponen
Biasanya digunakan untuk perulangan berbasis counter:
```go
package main

import "fmt"

func main() {
    // Inisialisasi (i := 1), Kondisi (i <= 3), Post Statement (i++)
    for i := 1; i <= 3; i++ {
        fmt.Println("Perulangan ke-", i)
    }
}
```

#### 2. Contoh for dengan Kondisi Saja (Gaya while)
Biasanya digunakan ketika variabel pengendali diubah di dalam blok kode:
```go
package main

import "fmt"

func main() {
    // Inisialisasi variabel dilakukan DI LUAR perulangan
    n := 1

    // Hanya menyertakan kondisi (mirip while di bahasa lain)
    for n <= 3 {
        fmt.Println("Nilai n:", n)
        
        // Post statement (pengubahan nilai) dilakukan DI DALAM blok
        n++ 
    }
}
```
Pada contoh kedua, `for n <= 3` bertindak persis seperti `while (n <= 3)` pada bahasa pemrograman seperti Java atau C++.