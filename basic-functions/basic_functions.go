package main

import (
	"fmt"
	"strconv"
	"strings"
)

// ===== FUNGSI DASAR =====

// 1. Fungsi printMe - mencetak pesan sederhana
func printMe(message string) {
	fmt.Println("Pesan:", message)
}

// 2. Fungsi intDivision - pembagian integer dengan penanganan error
func intDivision(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("tidak bisa membagi dengan nol")
	}
	return a / b, nil
}

// 3. Fungsi isEven - mengecek apakah angka genap
func isEven(num int) bool {
	return num%2 == 0
}

// 4. Fungsi getMax - mencari nilai maksimum dari dua angka
func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 5. Fungsi calculateGrade - menghitung grade berdasarkan nilai
func calculateGrade(score int) string {
	if score >= 90 {
		return "A"
	} else if score >= 80 {
		return "B"
	} else if score >= 70 {
		return "C"
	} else if score >= 60 {
		return "D"
	} else {
		return "F"
	}
}

// 6. Fungsi factorial - menghitung faktorial
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

// 7. Fungsi sumArray - menjumlahkan semua elemen dalam array
func sumArray(numbers []int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// 8. Fungsi countVowels - menghitung jumlah huruf vokal
func countVowels(text string) int {
	vowels := "aeiouAEIOU"
	count := 0
	for _, char := range text {
		if strings.ContainsRune(vowels, char) {
			count++
		}
	}
	return count
}

// 9. Fungsi reverseString - membalik string
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// 10. Fungsi isPrime - mengecek apakah angka prima
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

// ===== STRUKTUR KONTROL DASAR =====

// Fungsi demonstrasi if-else
func demonstrateIfElse() {
	fmt.Println("\n===== DEMONSTRASI IF-ELSE =====")

	// Contoh 1: Cek umur
	umur := 18
	if umur >= 18 {
		fmt.Println("Anda sudah dewasa")
	} else {
		fmt.Println("Anda masih di bawah umur")
	}

	// Contoh 2: Cek nilai
	nilai := 85
	grade := calculateGrade(nilai)
	fmt.Printf("Nilai %d mendapat grade: %s\n", nilai, grade)

	// Contoh 3: Cek genap/ganjil
	angka := 7
	if isEven(angka) {
		fmt.Printf("%d adalah angka genap\n", angka)
	} else {
		fmt.Printf("%d adalah angka ganjil\n", angka)
	}
}

// Fungsi demonstrasi switch
func demonstrateSwitch() {
	fmt.Println("\n===== DEMONSTRASI SWITCH =====")

	// Contoh 1: Switch dengan nilai
	hari := 3
	switch hari {
	case 1:
		fmt.Println("Senin")
	case 2:
		fmt.Println("Selasa")
	case 3:
		fmt.Println("Rabu")
	case 4:
		fmt.Println("Kamis")
	case 5:
		fmt.Println("Jumat")
	default:
		fmt.Println("Akhir pekan")
	}

	// Contoh 2: Switch tanpa ekspresi
	skor := 95
	switch {
	case skor >= 90:
		fmt.Println("Excellent!")
	case skor >= 80:
		fmt.Println("Good!")
	case skor >= 70:
		fmt.Println("Fair")
	default:
		fmt.Println("Need improvement")
	}

	// Contoh 3: Switch dengan multiple values
	bulan := "Januari"
	switch bulan {
	case "Desember", "Januari", "Februari":
		fmt.Println("Musim dingin")
	case "Maret", "April", "Mei":
		fmt.Println("Musim semi")
	case "Juni", "Juli", "Agustus":
		fmt.Println("Musim panas")
	case "September", "Oktober", "November":
		fmt.Println("Musim gugur")
	}
}

// Fungsi demonstrasi for loop
func demonstrateForLoop() {
	fmt.Println("\n===== DEMONSTRASI FOR LOOP =====")

	// Contoh 1: For loop dasar
	fmt.Println("Angka 1-5:")
	for i := 1; i <= 5; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// Contoh 2: For loop dengan array
	fmt.Println("Buah-buahan:")
	buah := []string{"Apel", "Jeruk", "Mangga", "Pisang"}
	for i, nama := range buah {
		fmt.Printf("%d. %s\n", i+1, nama)
	}

	// Contoh 3: For loop dengan kondisi (seperti while)
	fmt.Println("Countdown:")
	counter := 5
	for counter > 0 {
		fmt.Print(counter, " ")
		counter--
	}
	fmt.Println("Go!")

	// Contoh 4: For loop dengan break dan continue
	fmt.Println("Angka ganjil 1-10:")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue // skip angka genap
		}
		fmt.Print(i, " ")
	}
	fmt.Println()

	// Contoh 5: Nested loop (tabel perkalian)
	fmt.Println("Tabel perkalian 3x3:")
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("%d x %d = %d\t", i, j, i*j)
		}
		fmt.Println()
	}
}

// Fungsi demonstrasi penggunaan fungsi-fungsi yang sudah dibuat
func demonstrateFunctions() {
	fmt.Println("\n===== DEMONSTRASI FUNGSI =====")

	// Test printMe
	printMe("Hello, Go!")

	// Test intDivision
	result, err := intDivision(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("10 / 2 = %d\n", result)
	}

	// Test intDivision dengan error
	result, err = intDivision(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("10 / 0 = %d\n", result)
	}

	// Test getMax
	max := getMax(15, 25)
	fmt.Printf("Maksimum dari 15 dan 25 adalah: %d\n", max)

	// Test factorial
	fact := factorial(5)
	fmt.Printf("Faktorial dari 5 adalah: %d\n", fact)

	// Test sumArray
	numbers := []int{1, 2, 3, 4, 5}
	sum := sumArray(numbers)
	fmt.Printf("Jumlah dari %v adalah: %d\n", numbers, sum)

	// Test countVowels
	text := "Hello World"
	vowelCount := countVowels(text)
	fmt.Printf("Jumlah huruf vokal dalam '%s': %d\n", text, vowelCount)

	// Test reverseString
	original := "Go Programming"
	reversed := reverseString(original)
	fmt.Printf("'%s' dibalik menjadi: '%s'\n", original, reversed)

	// Test isPrime
	num := 17
	if isPrime(num) {
		fmt.Printf("%d adalah bilangan prima\n", num)
	} else {
		fmt.Printf("%d bukan bilangan prima\n", num)
	}
}

// Fungsi untuk input sederhana (simulasi)
func simulateUserInput() {
	fmt.Println("\n===== SIMULASI INPUT PENGGUNA =====")

	// Simulasi input nama dan umur
	nama := "Budi"
	umurStr := "25"

	umur, err := strconv.Atoi(umurStr)
	if err != nil {
		fmt.Println("Error konversi umur:", err)
		return
	}

	fmt.Printf("Nama: %s\n", nama)
	fmt.Printf("Umur: %d\n", umur)

	if umur >= 18 {
		fmt.Printf("%s sudah dewasa\n", nama)
	} else {
		fmt.Printf("%s masih di bawah umur\n", nama)
	}

	// Simulasi kalkulator sederhana
	a, b := 15, 4
	fmt.Printf("\nKalkulator sederhana untuk %d dan %d:\n", a, b)
	fmt.Printf("Penjumlahan: %d + %d = %d\n", a, b, a+b)
	fmt.Printf("Pengurangan: %d - %d = %d\n", a, b, a-b)
	fmt.Printf("Perkalian: %d * %d = %d\n", a, b, a*b)

	if hasil, err := intDivision(a, b); err == nil {
		fmt.Printf("Pembagian: %d / %d = %d\n", a, b, hasil)
	} else {
		fmt.Printf("Pembagian: Error - %s\n", err)
	}
}

func main() {
	fmt.Println("===== PROGRAM PEMBELAJARAN FUNGSI DAN STRUKTUR KONTROL DASAR =====")

	// Demonstrasi fungsi-fungsi yang sudah dibuat
	demonstrateFunctions()

	// Demonstrasi struktur kontrol
	demonstrateIfElse()
	demonstrateSwitch()
	demonstrateForLoop()

	// Simulasi input pengguna
	simulateUserInput()

	fmt.Println("\n===== PROGRAM SELESAI =====")
}
