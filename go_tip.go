package main

import (
	"fmt"
)

// https://devhints.io/go

// 常量
const pai = 3.14

// Named return values
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// Vertex is struct
type Vertex struct {
	X int
	Y int
}

// Abs is
func (v Vertex) Abs() float64 {
	return float64(v.X*v.X + v.Y*v.Y)
}

// Scale func for update v
func (v *Vertex) Scale(f float64) {
	v.X = v.X * 1
	v.Y = v.Y * 1
}

func main() {
	// 变量定义
	var msg string
	msg = "hello"
	msg1 := "world"
	msg2 := `many 
	line`
	message := msg + msg1 + msg2
	fmt.Println(message)

	num := 3          // int
	num1 := 3.        // float64
	num2 := 3 + 4i    // complex128
	num3 := byte('a') // byte (alias for uint8)
	// uint (unsigned)
	// float32
	f := float64(num)
	fmt.Printf("%d %f %d %d %f", num, num1, num2, num3, f)

	// Arrays have a fixed size
	var nums [5]int
	nums1 := [...]int{0, 0, 0, 0, 0}
	// Slices have a dynamic size, unlike arrays
	slice := []int{2, 3, 4}
	slice1 := []byte("Hello")
	fmt.Printf("%d %d %d %d %f", nums, nums1, slice, slice1)

	entry := []string{"Jack", "John", "Jones"}
	for i, val := range entry {
		fmt.Printf("At position %d, the character %s is present\n", i, val)
	}

}
