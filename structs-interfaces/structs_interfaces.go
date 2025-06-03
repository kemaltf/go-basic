package main

import (
	"fmt"
	"math"
)

// ========== STRUCTS ==========
// Struct di Go adalah tipe data yang memungkinkan Anda mengelompokkan beberapa variabel
// dengan tipe data berbeda menjadi satu kesatuan. Struct bisa dianggap sebagai "template"
// untuk membuat objek yang memiliki properti-properti tertentu.
// Jika dibandingkan dengan JavaScript, struct di Go mirip dengan:
// Object Literals di JavaScript

// 1. Basic struct definition
type Person struct {
	Name string
	Age  int
	City string
}

// 2. Struct dengan embedded fields
type Address struct {
	Street  string
	City    string
	ZipCode string
}

type Employee struct {
	Person  // Embedded struct
	Address // Embedded struct
	ID      int
	Salary  float64
}

// 3. Struct methods
func (p Person) Greet() string {
	return fmt.Sprintf("Hello, my name is %s and I'm %d years old", p.Name, p.Age)
}

// Pointer receiver method (dapat memodifikasi struct)
func (p *Person) HaveBirthday() {
	// Dengan pointer receiver, method dapat mengubah nilai pada struct asli, bukan hanya salinannya. Jika kita menulis:
	// 	person := Person{Name: "Alice", Age: 25}
	// person.HaveBirthday() // Age sekarang menjadi 26
	//  Jika tanpa tanda * maka hanya merubah salinannyay saja
	// Dalam Go, ketika Anda memanggil sebuah fungsi atau method, parameter yang
	// dikirim secara default adalah salinan (copy) dari nilai asli, bukan nilai asli
	// itu sendiri. Ini disebut pass by value.
	// Value receiver - bekerja dengan salinan
	// func (p Person) HaveBirthday() {
	//     p.Age++ // Hanya mengubah salinan struct Person
	// }
	person := Person{Name: "Alice", Age: 25}
	person.HaveBirthday()
	fmt.Println(person.Age) // Tetap 25, tidak berubah!
	p.Age++
}

// Value receiver method (tidak dapat memodifikasi struct)
func (p Person) IsAdult() bool {
	return p.Age >= 18
}

// ========== INTERFACES ==========

// 1. Basic interface
type Greeter interface {
	Greet() string
}

// 2. Interface dengan multiple methods
type Shape interface {
	Area() float64
	Perimeter() float64
}

// 3. Empty interface (dapat menerima tipe apa saja)
type Any interface{}

// ========== IMPLEMENTING INTERFACES ==========

// Rectangle implements Shape interface
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Circle implements Shape interface
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// ========== INTERFACE COMPOSITION ==========

// Interface yang menggabungkan interface lain
type Drawable interface {
	Shape // Embedded interface
	Draw() string
}

// Rectangle juga implements Drawable
func (r Rectangle) Draw() string {
	return fmt.Sprintf("Drawing rectangle: %.1fx%.1f", r.Width, r.Height)
}

// ========== TYPE ASSERTION & TYPE SWITCH ==========

func describeValue(v interface{}) {
	fmt.Printf("Value: %v, Type: %T\n", v, v)

	// Type assertion
	if str, ok := v.(string); ok {
		fmt.Printf("It's a string with length: %d\n", len(str))
	}

	// Type switch
	switch val := v.(type) {
	case string:
		fmt.Printf("String value: %s\n", val)
	case int:
		fmt.Printf("Integer value: %d\n", val)
	case Person:
		fmt.Printf("Person: %s\n", val.Name)
	case Shape:
		fmt.Printf("Shape with area: %.2f\n", val.Area())
	default:
		fmt.Printf("Unknown type: %T\n", val)
	}
}

// ========== POLYMORPHISM EXAMPLE ==========

func calculateTotalArea(shapes []Shape) float64 {
	total := 0.0
	for _, shape := range shapes {
		total += shape.Area()
	}
	return total
}

func printShapeInfo(s Shape) {
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())

	// Type assertion untuk akses method spesifik
	if drawable, ok := s.(Drawable); ok {
		fmt.Println(drawable.Draw())
	}
}

func main() {
	fmt.Println("=== STRUCTS EXAMPLES ===")

	// 1. Creating structs
	person1 := Person{Name: "Alice", Age: 25, City: "Jakarta"}
	person2 := Person{"Bob", 30, "Bandung"} // Positional initialization

	fmt.Println("Person 1:", person1)
	fmt.Println("Person 2:", person2)

	// 2. Accessing struct fields
	fmt.Printf("%s is %d years old\n", person1.Name, person1.Age)

	// 3. Struct methods
	fmt.Println(person1.Greet())
	fmt.Println("Is adult?", person1.IsAdult())

	// 4. Pointer receiver method
	fmt.Printf("Before birthday: %d\n", person1.Age)
	person1.HaveBirthday()
	fmt.Printf("After birthday: %d\n", person1.Age)

	// 5. Embedded structs
	emp := Employee{
		Person:  Person{Name: "Charlie", Age: 28, City: "Surabaya"},
		Address: Address{Street: "Jl. Merdeka", City: "Surabaya", ZipCode: "60111"},
		ID:      12345,
		Salary:  75000.0,
	}

	fmt.Printf("Employee: %s, ID: %d, Lives in: %s\n", emp.Name, emp.ID, emp.Address.City)
	fmt.Println(emp.Greet()) // Method dari embedded struct

	fmt.Println("\n=== INTERFACES EXAMPLES ===")

	// 6. Interface implementation
	var greeter Greeter = person1
	fmt.Println(greeter.Greet())

	// 7. Shape interface
	rect := Rectangle{Width: 10, Height: 5}
	circle := Circle{Radius: 3}

	fmt.Println("\nRectangle:")
	printShapeInfo(rect)

	fmt.Println("\nCircle:")
	printShapeInfo(circle)

	// 8. Polymorphism
	shapes := []Shape{rect, circle, Rectangle{Width: 4, Height: 6}}
	totalArea := calculateTotalArea(shapes)
	fmt.Printf("\nTotal area of all shapes: %.2f\n", totalArea)

	// 9. Empty interface examples
	fmt.Println("\n=== TYPE ASSERTION & SWITCH ===")
	values := []interface{}{"Hello", 42, person1, rect, 3.14}

	for i, val := range values {
		fmt.Printf("\nValue %d:\n", i+1)
		describeValue(val)
	}

	// 10. Interface composition
	fmt.Println("\n=== INTERFACE COMPOSITION ===")
	var drawable Drawable = rect
	fmt.Printf("Area: %.2f\n", drawable.Area())
	fmt.Println(drawable.Draw())

	fmt.Println("\n=== PRACTICAL EXAMPLES ===")

	// 11. Slice of interfaces
	greeters := []Greeter{person1, person2}
	for i, g := range greeters {
		fmt.Printf("Greeter %d: %s\n", i+1, g.Greet())
	}

	// 12. Interface sebagai parameter
	processShapes := func(shapes ...Shape) {
		for i, shape := range shapes {
			fmt.Printf("Shape %d - Type: %T, Area: %.2f\n", i+1, shape, shape.Area())
		}
	}

	processShapes(rect, circle)
}

// ========== ADVANCED CONCEPTS ==========

// 1. Struct tags (untuk JSON, database, dll)
type User struct {
	ID       int    `json:"id" db:"user_id"`
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email_address"`
}

// 2. Anonymous structs
func demonstrateAnonymousStruct() {
	config := struct {
		Host string
		Port int
	}{
		Host: "localhost",
		Port: 8080,
	}

	fmt.Printf("Config: %+v\n", config)
}

// 3. Interface dengan type constraints (Go 1.18+)
type Numeric interface {
	int | int64 | float64
}

// 4. Struct comparison
func demonstrateStructComparison() {
	p1 := Person{Name: "Alice", Age: 25, City: "Jakarta"}
	p2 := Person{Name: "Alice", Age: 25, City: "Jakarta"}
	p3 := Person{Name: "Bob", Age: 30, City: "Bandung"}

	fmt.Printf("p1 == p2: %t\n", p1 == p2) // true
	fmt.Printf("p1 == p3: %t\n", p1 == p3) // false
}

/*
=== KONSEP PENTING ===

1. STRUCTS:
   - Tipe data yang mengelompokkan field-field terkait
   - Dapat memiliki methods (receiver functions)
   - Mendukung embedding untuk komposisi
   - Value types (disalin saat assignment)

2. INTERFACES:
   - Kontrak yang mendefinisikan method signatures
   - Implementasi implisit (tidak perlu keyword 'implements')
   - Mendukung polymorphism
   - Empty interface{} dapat menerima tipe apa saja

3. METHODS:
   - Value receiver: func (t Type) method()
   - Pointer receiver: func (t *Type) method()
   - Pointer receiver dapat memodifikasi struct

4. TYPE ASSERTION:
   - value.(Type) untuk mengakses concrete type
   - value.(Type) mengembalikan (value, ok)

5. TYPE SWITCH:
   - switch v := value.(type) untuk handling multiple types

=== BEST PRACTICES ===

1. Gunakan pointer receivers untuk:
   - Method yang memodifikasi struct
   - Struct yang besar (untuk efisiensi)
   - Konsistensi jika ada method lain yang menggunakan pointer

2. Interface design:
   - Keep interfaces small and focused
   - Define interfaces di package yang menggunakannya
   - Prefer composition over large interfaces

3. Struct design:
   - Group related fields together
   - Use embedding untuk code reuse
   - Consider using struct tags untuk metadata

4. Error handling:
   - Selalu check ok value dari type assertion
   - Handle default case di type switch

=== COMMON PITFALLS ===

1. Nil interface vs nil pointer
2. Interface{} bukan generic (gunakan dengan hati-hati)
3. Method set rules (value vs pointer receivers)
4. Struct comparison dengan slice/map fields
5. Interface pollution (terlalu banyak interface kecil)

*/
