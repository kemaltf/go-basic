package main

import "fmt"

func main() {
	fmt.Println("=== ARRAY DASAR DI GO ===")

	// 1. DEKLARASI ARRAY
	fmt.Println("\n--- 1. Deklarasi Array ---")

	// Deklarasi dengan ukuran tetap
	var numbers [5]int // array dengan 5 elemen, semua bernilai 0
	fmt.Printf("Array kosong: %v\n", numbers)

	// Deklarasi dengan inisialisasi
	var fruits [3]string = [3]string{"apel", "jeruk", "pisang"}
	fmt.Printf("Array buah: %v\n", fruits)

	// Deklarasi singkat dengan inisialisasi
	colors := [4]string{"merah", "biru", "hijau", "kuning"}
	fmt.Printf("Array warna: %v\n", colors)

	// Ukuran otomatis berdasarkan elemen
	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat"}
	fmt.Printf("Array hari: %v (panjang: %d)\n", days, len(days))

	// 2. MENGAKSES DAN MENGUBAH ELEMEN
	fmt.Println("\n--- 2. Mengakses dan Mengubah Elemen ---")

	// Mengakses elemen berdasarkan index (dimulai dari 0)
	fmt.Printf("Elemen pertama: %s\n", fruits[0])
	fmt.Printf("Elemen kedua: %s\n", fruits[1])

	// Mengubah nilai elemen
	fruits[2] = "mangga"
	fmt.Printf("Array setelah diubah: %v\n", fruits)

	// Mengisi array numbers
	for i := 0; i < len(numbers); i++ {
		numbers[i] = (i + 1) * 10
	}
	fmt.Printf("Array numbers setelah diisi: %v\n", numbers)

	// 3. LOOP DENGAN ARRAY
	fmt.Println("\n--- 3. Loop dengan Array ---")

	// Loop biasa dengan index
	fmt.Println("Loop dengan index:")
	for i := 0; i < len(colors); i++ {
		fmt.Printf("Index %d: %s\n", i, colors[i])
	}

	// Loop dengan range (mendapat index dan value)
	fmt.Println("\nLoop dengan range (index + value):")
	for index, value := range days {
		fmt.Printf("Hari ke-%d: %s\n", index+1, value)
	}

	// Loop dengan range (hanya value)
	fmt.Println("\nLoop dengan range (hanya value):")
	for _, fruit := range fruits {
		fmt.Printf("Buah: %s\n", fruit)
	}

	// Loop dengan range (hanya index)
	fmt.Println("\nLoop dengan range (hanya index):")
	for i := range numbers {
		fmt.Printf("Index %d memiliki nilai: %d\n", i, numbers[i])
	}

	// 4. OPERASI DASAR ARRAY
	fmt.Println("\n--- 4. Operasi Dasar Array ---")

	// Panjang array
	fmt.Printf("Panjang array colors: %d\n", len(colors))

	// Mencari elemen dalam array
	target := "biru"
	found := false
	for i, color := range colors {
		if color == target {
			fmt.Printf("Warna '%s' ditemukan di index %d\n", target, i)
			found = true
			break
		}
	}
	if !found {
		fmt.Printf("Warna '%s' tidak ditemukan\n", target)
	}

	// Menghitung jumlah elemen dalam array numbers
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	fmt.Printf("Jumlah semua elemen numbers: %d\n", sum)

	// Mencari nilai maksimum
	max := numbers[0]
	for _, num := range numbers {
		if num > max {
			max = num
		}
	}
	fmt.Printf("Nilai maksimum: %d\n", max)

	// 5. ARRAY MULTIDIMENSI
	fmt.Println("\n--- 5. Array Multidimensi ---")

	// Array 2 dimensi (matrix)
	var matrix [3][3]int

	// Mengisi matrix
	value := 1
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			matrix[i][j] = value
			value++
		}
	}

	// Menampilkan matrix
	fmt.Println("Matrix 3x3:")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}

	// Array 2D dengan inisialisasi langsung
	scores := [2][3]int{
		{85, 90, 78},
		{92, 88, 95},
	}
	fmt.Println("\nNilai siswa:")
	for i, student := range scores {
		fmt.Printf("Siswa %d: ", i+1)
		for _, score := range student {
			fmt.Printf("%d ", score)
		}
		fmt.Println()
	}

	// 6. FUNGSI DENGAN ARRAY
	fmt.Println("\n--- 6. Fungsi dengan Array ---")

	// Memanggil fungsi dengan array
	average := calculateAverage(numbers)
	fmt.Printf("Rata-rata array numbers: %.2f\n", average)

	// Array sebagai parameter (pass by value)
	original := [3]int{1, 2, 3}
	fmt.Printf("Array sebelum fungsi: %v\n", original)
	modifyArray(original)
	fmt.Printf("Array setelah fungsi: %v\n", original) // Tidak berubah!

	// Menggunakan pointer untuk mengubah array
	fmt.Printf("Array sebelum modifikasi dengan pointer: %v\n", original)
	modifyArrayWithPointer(&original)
	fmt.Printf("Array setelah modifikasi dengan pointer: %v\n", original)

	// 7. KETERBATASAN ARRAY
	fmt.Println("\n--- 7. Keterbatasan Array ---")
	fmt.Println("YANG TIDAK BISA DILAKUKAN DENGAN ARRAY:")
	fmt.Println("1. Mengubah ukuran setelah deklarasi (ukuran tetap)")
	fmt.Println("2. Menambah atau menghapus elemen secara dinamis")
	fmt.Println("3. Array dengan ukuran berbeda tidak bisa dibandingkan")
	fmt.Println("4. Tidak ada built-in method seperti append, delete, dll")
	fmt.Println("5. Ukuran harus diketahui pada compile time")

	// Contoh keterbatasan
	// var dynamicArray [n]int // ERROR: n harus konstanta
	// numbers = append(numbers, 60) // ERROR: append tidak bisa untuk array

	// Demonstrasi perbandingan array
	arr1 := [3]int{1, 2, 3}
	arr2 := [3]int{1, 2, 3}

	fmt.Printf("\nPerbandingan array:")
	fmt.Printf("arr1 == arr2: %t\n", arr1 == arr2) // true
	// fmt.Printf("arr1 == arr3: %t\n", arr1 == arr3) // ERROR: ukuran berbeda

	// ========================================
	// BAGIAN BARU: SLICE DAN PERBANDINGANNYA
	// ========================================

	fmt.Println("\n\n=== SLICE DAN PERBANDINGAN DENGAN ARRAY ===")

	// 8. DEKLARASI SLICE
	fmt.Println("\n--- 8. Deklarasi Slice ---")

	// Slice kosong (nil)
	var sliceNumbers []int
	fmt.Printf("Slice kosong: %v (len: %d, cap: %d, nil: %t)\n", sliceNumbers, len(sliceNumbers), cap(sliceNumbers), sliceNumbers == nil)

	// Slice dengan inisialisasi
	sliceFruits := []string{"apel", "jeruk", "pisang"}
	fmt.Printf("Slice buah: %v (len: %d, cap: %d)\n", sliceFruits, len(sliceFruits), cap(sliceFruits))

	// Slice dengan make
	sliceColors := make([]string, 4) // length = 4, capacity = 4
	sliceColors[0] = "merah"
	sliceColors[1] = "biru"
	sliceColors[2] = "hijau"
	sliceColors[3] = "kuning"
	fmt.Printf("Slice warna: %v (len: %d, cap: %d)\n", sliceColors, len(sliceColors), cap(sliceColors))

	// Slice dengan make (length dan capacity berbeda)
	sliceDays := make([]string, 3, 10) // length = 3, capacity = 10
	sliceDays[0] = "Senin"
	sliceDays[1] = "Selasa"
	sliceDays[2] = "Rabu"
	fmt.Printf("Slice hari: %v (len: %d, cap: %d)\n", sliceDays, len(sliceDays), cap(sliceDays))

	// 9. OPERASI DINAMIS SLICE (YANG TIDAK BISA DILAKUKAN ARRAY)
	fmt.Println("\n--- 9. Operasi Dinamis Slice ---")

	// Append - menambah elemen
	fmt.Println("\nMenambah elemen dengan append:")
	sliceNumbers = append(sliceNumbers, 10)
	fmt.Printf("Setelah append 10: %v (len: %d, cap: %d)\n", sliceNumbers, len(sliceNumbers), cap(sliceNumbers))

	sliceNumbers = append(sliceNumbers, 20, 30, 40)
	fmt.Printf("Setelah append 20,30,40: %v (len: %d, cap: %d)\n", sliceNumbers, len(sliceNumbers), cap(sliceNumbers))

	// Append slice ke slice
	moreNumbers := []int{50, 60, 70}
	sliceNumbers = append(sliceNumbers, moreNumbers...)
	fmt.Printf("Setelah append slice: %v (len: %d, cap: %d)\n", sliceNumbers, len(sliceNumbers), cap(sliceNumbers))

	// Slicing - membuat sub-slice
	fmt.Println("\nSlicing operation:")
	subSlice := sliceNumbers[1:4] // index 1 sampai 3
	fmt.Printf("Original: %v\n", sliceNumbers)
	fmt.Printf("SubSlice [1:4]: %v (len: %d, cap: %d)\n", subSlice, len(subSlice), cap(subSlice))

	// Slicing dengan berbagai cara
	firstThree := sliceNumbers[:3] // dari awal sampai index 2
	lastThree := sliceNumbers[4:]  // dari index 4 sampai akhir
	middle := sliceNumbers[2:5]    // dari index 2 sampai 4
	fmt.Printf("First three [:3]: %v\n", firstThree)
	fmt.Printf("Last three [4:]: %v\n", lastThree)
	fmt.Printf("Middle [2:5]: %v\n", middle)

	// 10. PERBANDINGAN LANGSUNG ARRAY VS SLICE
	fmt.Println("\n--- 10. Perbandingan Langsung Array vs Slice ---")

	// Deklarasi
	fmt.Println("\nDeklarasi:")
	arrayExample := [3]int{1, 2, 3} // Array - ukuran tetap
	sliceExample := []int{1, 2, 3}  // Slice - ukuran dinamis
	fmt.Printf("Array: %v (type: %T)\n", arrayExample, arrayExample)
	fmt.Printf("Slice: %v (type: %T, len: %d, cap: %d)\n", sliceExample, sliceExample, len(sliceExample), cap(sliceExample))

	// Menambah elemen
	fmt.Println("\nMenambah elemen:")
	// arrayExample = append(arrayExample, 4) // ERROR: tidak bisa append ke array
	sliceExample = append(sliceExample, 4) // OK: bisa append ke slice
	fmt.Printf("Array tetap: %v\n", arrayExample)
	fmt.Printf("Slice bertambah: %v\n", sliceExample)

	// Passing ke function
	fmt.Println("\nPassing ke function:")
	fmt.Printf("Array sebelum: %v\n", arrayExample)
	fmt.Printf("Slice sebelum: %v\n", sliceExample)

	modifyArray(arrayExample) // Pass by value - tidak berubah
	modifySlice(sliceExample) // Pass by reference - berubah

	fmt.Printf("Array sesudah: %v (tidak berubah)\n", arrayExample)
	fmt.Printf("Slice sesudah: %v (berubah!)\n", sliceExample)

	// 11. SLICE DARI ARRAY
	fmt.Println("\n--- 11. Membuat Slice dari Array ---")

	bigArray := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("Array asli: %v\n", bigArray)

	// Membuat slice dari array
	sliceFromArray := bigArray[2:7] // index 2 sampai 6
	fmt.Printf("Slice dari array [2:7]: %v (len: %d, cap: %d)\n", sliceFromArray, len(sliceFromArray), cap(sliceFromArray))

	// Mengubah slice akan mengubah array asli (shared memory)
	sliceFromArray[0] = 999
	fmt.Printf("Setelah ubah slice: %v\n", sliceFromArray)
	fmt.Printf("Array asli juga berubah: %v\n", bigArray)

	// 12. COPY SLICE
	fmt.Println("\n--- 12. Copy Slice ---")

	source := []int{1, 2, 3, 4, 5}
	destination := make([]int, len(source))

	// Copy slice
	copied := copy(destination, source)
	fmt.Printf("Source: %v\n", source)
	fmt.Printf("Destination: %v (copied %d elements)\n", destination, copied)

	// Mengubah destination tidak mengubah source
	destination[0] = 999
	fmt.Printf("Setelah ubah destination: %v\n", destination)
	fmt.Printf("Source tetap: %v\n", source)

	// 13. SLICE MULTIDIMENSI
	fmt.Println("\n--- 13. Slice Multidimensi ---")

	// Slice 2D
	sliceMatrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("Slice Matrix 2D:")
	for i, row := range sliceMatrix {
		fmt.Printf("Row %d: %v\n", i, row)
	}

	// Menambah row baru ke slice 2D
	newRow := []int{10, 11, 12}
	sliceMatrix = append(sliceMatrix, newRow)
	fmt.Println("\nSetelah menambah row:")
	for i, row := range sliceMatrix {
		fmt.Printf("Row %d: %v\n", i, row)
	}

	// 14. PERFORMA DAN MEMORY
	fmt.Println("\n--- 14. Performa dan Memory ---")

	// Demonstrasi capacity growth
	fmt.Println("\nPertumbuhan kapasitas slice:")
	growingSlice := make([]int, 0)
	for i := 0; i < 10; i++ {
		oldCap := cap(growingSlice)
		growingSlice = append(growingSlice, i)
		newCap := cap(growingSlice)
		if newCap != oldCap {
			fmt.Printf("Append %d: capacity berubah dari %d ke %d\n", i, oldCap, newCap)
		}
	}
	fmt.Printf("Final slice: %v (len: %d, cap: %d)\n", growingSlice, len(growingSlice), cap(growingSlice))

	fmt.Println("\n=== KESIMPULAN ===")
	fmt.Println("Array cocok untuk:")
	fmt.Println("- Data dengan ukuran tetap dan diketahui")
	fmt.Println("- Performa tinggi (akses langsung ke memori)")
	fmt.Println("- Struktur data sederhana")
	fmt.Println("\nSlice cocok untuk:")
	fmt.Println("- Data dengan ukuran dinamis")
	fmt.Println("- Operasi append, slicing, copy")
	fmt.Println("- Kebanyakan kasus penggunaan sehari-hari")
	fmt.Println("\nRekomendasi: Gunakan SLICE untuk fleksibilitas, Array untuk performa maksimal!")
}

// Fungsi untuk menghitung rata-rata array
func calculateAverage(arr [5]int) float64 {
	sum := 0
	for _, value := range arr {
		sum += value
	}
	return float64(sum) / float64(len(arr))
}

// Fungsi yang mencoba mengubah array (tidak berhasil karena pass by value)
func modifyArray(arr [3]int) {
	arr[0] = 999
	fmt.Printf("Di dalam fungsi array: %v\n", arr)
}

// Fungsi yang mengubah array menggunakan pointer
func modifyArrayWithPointer(arr *[3]int) {
	arr[0] = 999
	fmt.Printf("Di dalam fungsi dengan pointer: %v\n", *arr)
}

// Fungsi yang mengubah slice (berhasil karena pass by reference)
func modifySlice(slice []int) {
	if len(slice) > 0 {
		slice[0] = 888
	}
	fmt.Printf("Di dalam fungsi slice: %v\n", slice)
}
