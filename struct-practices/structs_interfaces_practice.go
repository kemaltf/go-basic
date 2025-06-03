package main

import "fmt"

type GasEngine struct {
	gallons uint8 // gallons of gas left
	mpg     uint8 // miles per gallon
}

// Fungsi milesLeft() yang Anda tunjukkan adalah sebuah method yang terkait dengan struct GasEngine .
func (g GasEngine) milesLeft() uint8 {
	return g.gallons * g.mpg // jarak tersisa = jumlah galon × mil per galon
}

type ElectricEngine struct {
	kwh   uint8
	mpkwh uint8
}

// Keduanya (Gas dan Electric) punya method yang sama: milesLeft().
func (e ElectricEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh // jarak tersisa = jumlah kwh × mil per kwh
}

// Tantangannya:
// Fungsi seperti ini hanya menerima GasEngine:
// Kalau mau juga menerima ElectricEngine, tidak bisa — harus overload
// (yang tidak tersedia di Go)
// func canMakeIt(e GasEngine, miles uint8) bool {
// 	return e.milesLeft() >= miles
// }

//Solusinya: Gunakan interface
// Engine interface defines the behavior for different engine types
type Engine interface {
	milesLeft() uint8
}

func canMakeIt(e Engine, miles uint8) bool {
	return e.milesLeft() >= miles
}

func main() {
	// Declare as Engine interface type to hold different engine implementations
	var engine Engine
	// Assign GasEngine
	engine = GasEngine{gallons: 10, mpg: 20}
	fmt.Println(engine.milesLeft()) // Output: 200

	// Assign ElectricEngine
	engine = ElectricEngine{kwh: 10, mpkwh: 20}
	fmt.Println(engine.milesLeft()) // Output: 200

	// Dengan ElectricEngine
	electricEngine := ElectricEngine{kwh: 10, mpkwh: 20}
	fmt.Println(canMakeIt(electricEngine, 150)) // Output: true (karena milesLeft = 200)
	fmt.Println(canMakeIt(electricEngine, 250)) // Output: false (karena milesLeft = 200)
}
