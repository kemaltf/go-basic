package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	// Dalam Go, string adalah kumpulan byte, yang secara default di-encode menggunakan UTF-8.
	// Ini artinya setiap karakter dalam string bisa direpresentasikan oleh satu atau lebih byte
	// tergantung dari karakter tersebut.

	// Indexing String
	// Kita bisa mengakses elemen string seperti mengakses array.
	// Namun, ketika kita mengaksesnya, kita tidak mendapatkan karakter, melainkan sebuah angka,
	// yaitu nilai byte dari karakter tersebut.
	s := "résumé"

	fmt.Printf("String: %s\n", s)
	fmt.Printf("Jumlah bytes: %d\n", len(s))
	fmt.Printf("Jumlah karakter: %d\n", utf8.RuneCountInString(s))
	fmt.Println()

	// Byte representation
	// 	### Apa itu Byte?
	// Byte adalah unit data terkecil yang terdiri dari 8 bit . Satu byte dapat menyimpan nilai dari 0 sampai 255 (dalam desimal) atau 00 sampai FF (dalam heksadesimal).
	fmt.Println("Bytes:", []byte(s))
	fmt.Println()

	// 	### Apa itu Unicode?
	// Unicode adalah standar internasional untuk encoding
	// karakter dari semua bahasa di dunia. Unicode memberikan
	// nomor unik (disebut code point ) untuk setiap karakter.

	// Indexing byte (BERBAHAYA untuk Unicode!)
	fmt.Printf("s[0] = %d (karakter '%c')\n", s[0], s[0])
	fmt.Println()

	// Unicode code points perlu di- encode menjadi bytes untuk disimpan di komputer. Ada beberapa encoding:
	//  1. ASCII (1 byte per karakter)
	//  - Hanya untuk karakter Latin dasar (A-Z, a-z, 0-9, simbol)
	//  - 1 karakter = 1 byte
	//  - Terbatas pada 128 karakter 2. UTF-8 (Variable-length encoding)
	//  - 1-4 bytes per karakter
	//  - Kompatibel dengan ASCII
	//  - Encoding paling populer di web

	// Rune adalah tipe data di Go yang merepresentasikan satu karakter Unicode . Secara teknis, rune adalah alias untuk tipe int32 yang menyimpan Unicode code point .

	// Range iteration (AMAN untuk Unicode)
	fmt.Println("Range iteration:")
	for i, v := range s {
		fmt.Printf("Posisi byte %d: Unicode %d (karakter '%c')\n", i, v, v)
	}
	fmt.Println()

	// Menunjukkan struktur byte
	fmt.Println("Struktur byte detail:")
	for i := 0; i < len(s); i++ {
		fmt.Printf("Byte %d: %d\n", i, s[i])
	}

	fmt.Println("=== DEMONSTRASI STRING DI GO ===")
	fmt.Println()

	// ===== 1. STRING SEBAGAI ARRAY BYTE (UTF-8) =====
	fmt.Println("1. STRING SEBAGAI ARRAY BYTE (UTF-8):")

	text := "Hello, 世界!"
	fmt.Printf("String: %s\n", text)
	fmt.Printf("Sebagai bytes: %v\n", []byte(text))
	fmt.Printf("Sebagai bytes (hex): %x\n", []byte(text))
	fmt.Println()

	// Menunjukkan encoding UTF-8
	indonesian := "Halo, dunia! 🌍"
	fmt.Printf("String Indonesia: %s\n", indonesian)
	fmt.Printf("Bytes: %v\n", []byte(indonesian))
	fmt.Printf("Jumlah bytes: %d\n", len([]byte(indonesian)))
	fmt.Println()

	// ===== 2. INDEXING STRING MENGAKSES BYTE, BUKAN KARAKTER =====
	fmt.Println("2. INDEXING STRING MENGAKSES BYTE:")

	multiLang := "ABC世界"
	fmt.Printf("String: %s\n", multiLang)
	fmt.Printf("len(string): %d bytes\n", len(multiLang))
	fmt.Println()

	// Indexing byte-by-byte (BERBAHAYA untuk Unicode!)
	fmt.Println("Indexing byte-by-byte:")
	for i := 0; i < len(multiLang); i++ {
		fmt.Printf("Index %d: byte value %d, char '%c'\n", i, multiLang[i], multiLang[i])
	}
	fmt.Println()

	// Menunjukkan masalah dengan indexing byte
	fmt.Println("MASALAH: Karakter Unicode membutuhkan multiple bytes")
	chinese := "世界"
	fmt.Printf("String '%s' memiliki %d bytes tapi hanya %d karakter\n",
		chinese, len(chinese), utf8.RuneCountInString(chinese))

	// Coba akses byte individual (akan rusak!)
	fmt.Println("Akses byte individual (SALAH):")
	for i := 0; i < len(chinese); i++ {
		fmt.Printf("Byte %d: %c (tidak valid!)\n", i, chinese[i])
	}
	fmt.Println()

	// ===== 3. GUNAKAN RANGE UNTUK MENGAKSES KARAKTER UNICODE =====
	fmt.Println("3. GUNAKAN RANGE UNTUK KARAKTER UNICODE:")

	mixed := "Hello世界🌍"
	fmt.Printf("String: %s\n", mixed)
	fmt.Printf("Bytes: %d, Karakter: %d\n", len(mixed), utf8.RuneCountInString(mixed))
	fmt.Println()

	// Range otomatis decode UTF-8
	fmt.Println("Menggunakan range (BENAR):")
	for i, char := range mixed {
		fmt.Printf("Posisi byte %d: karakter '%c' (Unicode: U+%04X)\n", i, char, char)
	}
	fmt.Println()

	// Bandingkan dengan loop biasa
	fmt.Println("Loop biasa (byte-by-byte, SALAH untuk Unicode):")
	for i := 0; i < len(mixed); i++ {
		fmt.Printf("Byte %d: %c\n", i, mixed[i])
	}
	fmt.Println()

	// ===== 4. KONVERSI KE []RUNE UNTUK MANIPULASI KARAKTER =====
	fmt.Println("4. KONVERSI KE []RUNE:")

	original := "Café 咖啡 ☕"
	fmt.Printf("Original: %s\n", original)

	// Konversi ke rune slice
	runes := []rune(original)
	fmt.Printf("Sebagai runes: %v\n", runes)
	fmt.Printf("Jumlah runes: %d\n", len(runes))
	fmt.Println()

	// Sekarang bisa akses karakter dengan index
	fmt.Println("Akses karakter dengan index rune:")
	for i, r := range runes {
		fmt.Printf("Karakter %d: '%c' (U+%04X)\n", i, r, r)
	}
	fmt.Println()

	// Manipulasi karakter
	fmt.Println("Manipulasi karakter:")
	if len(runes) > 0 {
		runes[0] = 'K' // Ubah 'C' menjadi 'K'
	}
	modified := string(runes)
	fmt.Printf("Setelah modifikasi: %s\n", modified)
	fmt.Println()

	// ===== 5. LEN() MENGUKUR BYTE, BUKAN KARAKTER =====
	fmt.Println("5. LEN() MENGUKUR BYTE:")

	testStrings := []string{
		"ABC",           // ASCII
		"café",          // Latin dengan accent
		"世界",            // Chinese
		"🌍🌎🌏",           // Emoji
		"Iñtërnâtiônàl", // Mixed
	}

	for _, s := range testStrings {
		fmt.Printf("String: %-15s | Bytes: %2d | Karakter: %2d\n",
			s, len(s), utf8.RuneCountInString(s))
	}
	fmt.Println()

	// Fungsi helper untuk menghitung karakter
	fmt.Println("Fungsi helper untuk menghitung karakter:")
	for _, s := range testStrings {
		fmt.Printf("%s: %d karakter\n", s, countCharacters(s))
	}
	fmt.Println()

	// ===== 6. STRINGS.BUILDER UNTUK EFISIENSI =====
	fmt.Println("6. STRINGS.BUILDER UNTUK EFISIENSI:")
	fmt.Println()

	// Cara TIDAK efisien (string concatenation)
	fmt.Println("Cara TIDAK efisien (string concatenation):")
	result1 := demonstrateInefficient()
	fmt.Printf("Hasil: %s\n", result1)
	fmt.Println()

	// Cara efisien dengan strings.Builder
	fmt.Println("Cara efisien (strings.Builder):")
	result2 := demonstrateEfficient()
	fmt.Printf("Hasil: %s\n", result2)
	fmt.Println()

	// Perbandingan performa
	fmt.Println("Perbandingan performa:")
	comparePerformance()
	fmt.Println()

	// ===== 7. CONTOH PRAKTIS =====
	fmt.Println("7. CONTOH PRAKTIS:")
	fmt.Println()

	// Reverse string dengan benar
	originalText := "Hello, 世界! 🌍"
	reversed := reverseString(originalText)
	fmt.Printf("Original: %s\n", originalText)
	fmt.Printf("Reversed: %s\n", reversed)
	fmt.Println()

	// Truncate string dengan aman
	longText := "This is a very long text with Unicode: 世界 🌍"
	truncated := safeSubstring(longText, 0, 20)
	fmt.Printf("Original: %s\n", longText)
	fmt.Printf("Truncated (20 chars): %s\n", truncated)
	fmt.Println()

	// Validasi UTF-8
	validateUTF8Examples()
	fmt.Println()

	// ===== 8. COMMON PITFALLS =====
	fmt.Println("8. COMMON PITFALLS:")
	fmt.Println()

	demonstratePitfalls()

	fmt.Println("=== SELESAI ===")
}

// Helper functions

func countCharacters(s string) int {
	return utf8.RuneCountInString(s)
}

func demonstrateInefficient() string {
	// JANGAN lakukan ini untuk banyak concatenation!
	result := ""
	for i := 0; i < 5; i++ {
		result += fmt.Sprintf("Item %d ", i)
	}
	return result
}

func demonstrateEfficient() string {
	// Gunakan strings.Builder untuk efisiensi
	var builder strings.Builder

	// Optional: pre-allocate jika tahu ukuran perkiraan
	builder.Grow(50)

	for i := 0; i < 5; i++ {
		builder.WriteString(fmt.Sprintf("Item %d ", i))
	}

	return builder.String()
}

func comparePerformance() {
	// Simulasi perbandingan (untuk demo)
	fmt.Println("String concatenation: Lambat untuk banyak operasi (O(n²))")
	fmt.Println("strings.Builder: Cepat dan efisien (O(n))")
	fmt.Println("Rekomendasi: Gunakan strings.Builder untuk >5 concatenations")
}

func reverseString(s string) string {
	// Konversi ke rune untuk handle Unicode dengan benar
	runes := []rune(s)

	// Reverse slice
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func safeSubstring(s string, start, length int) string {
	// Konversi ke rune untuk akses karakter yang aman
	runes := []rune(s)

	if start >= len(runes) {
		return ""
	}

	end := start + length
	if end > len(runes) {
		end = len(runes)
	}

	return string(runes[start:end])
}

func validateUTF8Examples() {
	fmt.Println("Validasi UTF-8:")

	validStrings := []string{
		"Hello",
		"世界",
		"🌍",
		"Café",
	}

	for _, s := range validStrings {
		if utf8.ValidString(s) {
			fmt.Printf("✓ '%s' adalah UTF-8 valid\n", s)
		} else {
			fmt.Printf("✗ '%s' bukan UTF-8 valid\n", s)
		}
	}

	// Contoh invalid UTF-8 (dibuat dari byte array)
	invalidBytes := []byte{0xff, 0xfe, 0xfd}
	invalidString := string(invalidBytes)
	if utf8.ValidString(invalidString) {
		fmt.Printf("✓ String dari bytes invalid adalah UTF-8 valid\n")
	} else {
		fmt.Printf("✗ String dari bytes invalid bukan UTF-8 valid\n")
	}
}

func demonstratePitfalls() {
	fmt.Println("PITFALL 1: Substring dengan byte index")
	text := "Hello, 世界!"
	// SALAH: bisa memotong di tengah karakter Unicode
	// wrongSub := text[0:8] // Bisa rusak!
	fmt.Printf("Text: %s (bytes: %d, chars: %d)\n",
		text, len(text), utf8.RuneCountInString(text))
	fmt.Println("JANGAN: text[0:8] - bisa memotong karakter Unicode!")
	fmt.Println("GUNAKAN: safeSubstring() atau konversi ke []rune")
	fmt.Println()

	fmt.Println("PITFALL 2: Loop dengan index untuk Unicode")
	unicodeText := "🌍🌎🌏"
	fmt.Printf("Text: %s\n", unicodeText)
	fmt.Println("SALAH: for i := 0; i < len(text); i++")
	fmt.Println("BENAR: for i, char := range text")
	fmt.Println()

	fmt.Println("PITFALL 3: Mengasumsikan 1 byte = 1 karakter")
	emoji := "👨‍👩‍👧‍👦" // Family emoji (complex)
	fmt.Printf("Emoji: %s\n", emoji)
	fmt.Printf("Bytes: %d, Runes: %d, Visual chars: 1\n",
		len(emoji), utf8.RuneCountInString(emoji))
	fmt.Println("Catatan: Emoji kompleks bisa terdiri dari multiple runes!")
}
