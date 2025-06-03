// Di Go, pointer adalah variabel yang menyimpan alamat memori dari variabel lain.
// A pointer is a variable that stores the memory address of another variable
// — not the value itself, but the location where it's stored.
package main

import "fmt"

// var x int = 42
// var p *int = &x
// - x holds the value 42
// - p is a pointer to x — it holds the address in memory where x lives.
// - *p lets you access or change the value at that memory address (called dereferencing).

// 🧠 Why Use Pointers in Go?
// "Why not just pass around variables?"
// If you pass a struct or array into a function, you’re passing a copy.
// Changes inside the function won’t affect the original.
// But if you pass a pointer, you're working with the same data — no copy. So:
// ✔️ Less memory usage
// ✔️ Can mutate original value
// ✔️ Better performance for large data

// Using & to get the address of an existing variable:
// i := 10
// p := &i // pointer to i

// Using new() to create space:
// p := new(int) // creates a pointer to a newly allocated int (default value: 0)
// new() di Go itu fungsinya buat nyiapin ruang kosong di memori buat sebuah tipe data tertentu,
// lalu ngembaliin alamat memori-nya (alias pointer-nya).

// Go bikin sebuah integer kosong (nilainya 0 secara default).
// Trus new() ngasih kamu alamat di mana integer itu disimpan di memori.
// p jadi pointer ke integer itu.

// Kalau kamu tulis:
// *p = 10
// Itu artinya: masukin angka 10 ke dalam ruang memori yang ditunjuk oleh p.

// Bayangin kamu nyewa loker (lemari penyimpanan):
// new(int) = kamu sewa loker kosong buat nyimpen angka (otomatis isinya 0).
// p = kamu pegang kunci ke loker itu.
// *p = 10 = kamu buka loker dan masukin angka 10 ke dalamnya.

// kalo pake new() itu buat nyiapin ruang kosong di memori buat sebuah tipe data tertentu,
// lalu ngembaliin alamat memori-nya (alias pointer-nya).
// kalo pake & itu buat ngambil alamat memori dari sebuah variabel.
func main() {
	var i int32  // i = 0 (default)
	var p *int32 // p = nil

	fmt.Println("Nilai awal i:", i)
	fmt.Println("Alamat memori i:", &i)
	fmt.Println("Nilai awal p:", p)
	fmt.Println("Alamat memori p:", &p)

	p = &i // p menunjuk ke alamat memori i
	fmt.Println("Setelah p = &i, nilai p:", p)
	fmt.Println("Nilai yang ditunjuk p (*p):", *p) // *p adalah nilai yang ditunjuk p
	fmt.Println("Nilai i:", i)                     // i tetap 0 karena p tidak mengubah nilainya

	*p = 42 // isi dari alamat yang ditunjuk p diubah jadi 42
	fmt.Println("Setelah *p = 42, nilai i:", i)
	fmt.Println("Nilai yang ditunjuk p (*p):", *p)

	// Contoh penggunaan new(int32)
	fmt.Println("\n--- Contoh penggunaan new(int32) ---")
	// new() membuat ruang memori baru dan mengembalikan pointer ke ruang tersebut
	pNew := new(int32)                  // Membuat pointer ke int32 baru dengan nilai default 0 "Pointer ke int32" artinya: kertas yang berisi alamat dari kotak yang menyimpan angka int32
	fmt.Println("Nilai pNew:", pNew)    // Nilai pNew: 0xc00010c0ac
	fmt.Println("Nilai pNew &:", &pNew) // 	Nilai pNew: 0xc00010e068
	fmt.Println("Nilai yang ditunjuk pNew (*pNew):", *pNew)

	// 	Memori:
	// ┌─────────────────┐
	// │ 0xc00010e068    │ ← Alamat tempat pNew disimpan (&pNew)
	// │ [0xc00010c0ac]  │ ← Nilai pNew (pointer ke int32)
	// └─────────────────┘
	//         │
	//         ▼
	// ┌─────────────────┐
	// │ 0xc00010c0ac    │ ← Alamat yang ditunjuk pNew
	// │ [0]             │ ← Nilai int32 (*pNew)
	// └─────────────────┘

	// Mengubah nilai melalui pointer
	*pNew = 100
	fmt.Println("Setelah *pNew = 100, nilai *pNew:", *pNew)

	// Perbedaan dengan deklarasi var p *int32
	var pNil *int32 // Hanya deklarasi pointer, nilainya nil
	fmt.Println("Nilai pNil:", pNil)
	// fmt.Println("Nilai *pNil:", *pNil) // Ini akan error: panic: runtime error: invalid memory address
	// or nil pointer dereference
}
