package main

import (
	"fmt"
	"math"
)

func main() {
	// ===== KONSTANTA =====
	// Konstanta adalah nilai yang tidak dapat diubah setelah dideklarasikan

	// Deklarasi konstanta dengan tipe data eksplisit
	const PI float64 = 3.14159

	// Deklarasi konstanta tanpa tipe data eksplisit (tipe akan disimpulkan)
	const GRAVITY = 9.8

	// Deklarasi multiple konstanta
	const (
		APP_NAME    = "Aplikasi Belajar Go"
		APP_VERSION = "1.0.0"
		IS_BETA     = true
	)

	// Konstanta iota (auto-increment)
	const (
		SENIN  = iota // 0
		SELASA        // 1
		RABU          // 2
		KAMIS         // 3
		JUMAT         // 4
	)

	fmt.Println("===== KONSTANTA =====")
	fmt.Println("PI:", PI)
	fmt.Println("GRAVITY:", GRAVITY)
	fmt.Println("APP_NAME:", APP_NAME)
	fmt.Println("HARI RABU:", RABU)

	// ===== VARIABEL =====
	// Variabel adalah nilai yang dapat diubah

	// Deklarasi variabel dengan tipe data eksplisit
	var umur int = 25

	// Deklarasi variabel tanpa nilai awal
	var nama string
	nama = "Budi"

	// Deklarasi variabel dengan inferensi tipe (tanpa var)
	tinggi := 175.5 // float64

	// Deklarasi multiple variabel
	var (
		alamat  string = "Jalan Merdeka No. 10"
		menikah bool   = false
		// gaji    int    = 5000000
	)

	// Mengubah nilai variabel
	umur = 26

	fmt.Println("\n===== VARIABEL =====")
	fmt.Println("Nama:", nama)
	fmt.Println("Umur:", umur)
	fmt.Println("Tinggi:", tinggi)
	fmt.Println("Alamat:", alamat)
	fmt.Println("Menikah:", menikah)

	// ===== TIPE DATA DASAR =====
	fmt.Println("\n===== TIPE DATA DASAR =====")

	// 1. Numerik Integer
	fmt.Println("\n--- Integer ---")
	var a int8 = 127         // -128 hingga 127
	var b int16 = 32767      // -32768 hingga 32767
	var c int32 = 2147483647 // -2147483648 hingga 2147483647
	var d int64 = 9223372036854775807
	var e uint8 = 255 // 0 hingga 255

	fmt.Println("int8:", a)
	fmt.Println("int16:", b)
	fmt.Println("int32:", c)
	fmt.Println("int64:", d)
	fmt.Println("uint8:", e)

	// 2. Numerik Float
	fmt.Println("\n--- Float ---")
	var f float32 = 3.14
	var g float64 = math.Pi

	fmt.Println("float32:", f)
	fmt.Println("float64:", g)

	// 3. Boolean
	fmt.Println("\n--- Boolean ---")
	var benar bool = true
	var salah bool = false

	fmt.Println("benar:", benar)
	fmt.Println("salah:", salah)

	// 4. String
	fmt.Println("\n--- String ---")
	var pesan string = "Halo, Dunia!"
	var karakter byte = pesan[0] // 'H' dalam ASCII

	fmt.Println("pesan:", pesan)
	fmt.Println("karakter pertama (byte):", karakter)
	fmt.Println("karakter pertama (string):", string(karakter))
	fmt.Println("panjang string:", len(pesan))

	// 5. Rune (karakter UTF-8)
	fmt.Println("\n--- Rune ---")
	var emoji rune = '😀' // Karakter Unicode

	fmt.Printf("emoji: %c (nilai desimal: %d)\n", emoji, emoji)

	// 6. Complex
	fmt.Println("\n--- Complex ---")
	var bilangan complex64 = 3 + 4i

	fmt.Println("bilangan kompleks:", bilangan)
	fmt.Println("bagian real:", real(bilangan))
	fmt.Println("bagian imajiner:", imag(bilangan))

	// 7. Array
	fmt.Println("\n--- Array ---")
	var buah [3]string = [3]string{"Apel", "Jeruk", "Mangga"}

	fmt.Println("buah:", buah)
	fmt.Println("buah[1]:", buah[1])
	fmt.Println("jumlah elemen:", len(buah))

	// 8. Slice (array dinamis)
	fmt.Println("\n--- Slice ---")
	var warna = []string{"Merah", "Hijau", "Biru"}
	warna = append(warna, "Kuning") // Menambah elemen

	fmt.Println("warna:", warna)
	fmt.Println("warna[2]:", warna[2])
	fmt.Println("jumlah elemen:", len(warna))

	// 9. Map
	fmt.Println("\n--- Map ---")
	var kamus = map[string]string{
		"hello": "halo",
		"world": "dunia",
		"good":  "baik",
	}

	fmt.Println("kamus:", kamus)
	fmt.Println("kamus['hello']:", kamus["hello"])
	fmt.Println("jumlah elemen:", len(kamus))

	// 10. Struct
	fmt.Println("\n--- Struct ---")
	type Mahasiswa struct {
		Nama  string
		NIM   string
		IPK   float64
		Aktif bool
	}

	var mhs = Mahasiswa{
		Nama:  "Andi",
		NIM:   "12345",
		IPK:   3.75,
		Aktif: true,
	}

	fmt.Println("mahasiswa:", mhs)
	fmt.Println("nama:", mhs.Nama)
	fmt.Println("IPK:", mhs.IPK)
}
