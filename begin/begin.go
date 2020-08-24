package begin

import (
	"fmt"
	"math"
)

//CallBegin - calls certain begin task
func CallBegin() {
	var id int = 0
	print("Enter the task number (1 to 10): ")
	fmt.Scan(&id)
	if id > 0 {
		switch id {
		case 1:
			begin1()
		case 2:
			begin2()
		case 3:
			begin3()
		case 4:
			begin4()
		case 5:
			begin5()
		case 6:
			begin6()
		case 7:
			begin7()
		case 8:
			begin8()
		case 9:
			begin9()
		case 10:
			begin10()
		}
	}
}

func begin1() {
	println(`
	Given the side a of a square, 
	find the perimeter P of the square:
	P = 4·a
	`)
	print("a = ")
	var a float64
	fmt.Scan(&a)
	p := 4 * a
	fmt.Printf("P = %.2f", p)
}

func begin2() {
	fmt.Println(`
	Given the side a of a square,
	find the area S of the square: 
	S = a^2
	`)
	print("a = ")
	var a float64
	fmt.Scan(&a)
	s := a * a
	fmt.Printf("S = %.2f", s)
}

func begin3() {
	fmt.Println(`
	The sides a and b of a rectangle are given.
	Find the area S = a*b and the perimeter P = 2·(a + b) 
	of the rectangle.
	`)
	var a, b float64
	print("a = ")
	fmt.Scan(&a)
	print("b = ")
	fmt.Scan(&b)
	s, p := a*b, 2*(a+b)
	fmt.Printf("S = %.2f\nP = %f", s, p)
}

func begin4() {
	fmt.Println(`
	Given the diameter d of a circle,
	find the length L of the circle:
	L = π·d. Use 3.14 for a value of π.
	`)
	var d float64
	print("d = ")
	fmt.Scan(&d)
	const pi = 3.14
	l := pi * d
	fmt.Printf("L = %.2f", l)
}

func begin5() {
	fmt.Println(`
	Given the edge a of a cube, find the volume V = a^3
	and the surface area S = 6·a^2 of the cube.
	`)
	var a float64
	print("a = ")
	fmt.Scan(&a)
	v, s := a*a*a, 6*a*a
	fmt.Printf("V = %.2f\nS = %.2f", v, s)
}

func begin6() {
	fmt.Println(`
	The coordinates (x1, y1) and (x2, y2) of two points are given.
	Find the distance between the points: 
	((x2 − x1)^2 + (y2 − y1)^2)^(1/2)
	`)
	var x1, y1, x2, y2 float64
	print("x1, y1 = ")
	fmt.Scan(&x1, &y1)
	print("x2, y2 = ")
	fmt.Scan(&x2, &y2)
	dist := math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
	fmt.Printf("Distance = %.2f", dist)
}

func begin7() {
	fmt.Println(`
	Given the area S of a circle, 
	find the diameter D and the length L
	of the circumference. Take into account that
	L = π·D, S = π·(D^2)/4. 
	Use 3.14 for a value of π.
	`)
	const pi = 3.14
	var s float64
	print("S = ")
	fmt.Scan(&s)
	d := math.Sqrt((4 * s) / pi)
	l := pi * d
	fmt.Printf("L = %.2f\nD = %.2f", d, l)
}

func begin8() {
	fmt.Println(`
	Two nonzero numbers are given. Find the sum, the difference,
	the product, and the quotient of their squares.
	`)
	var a, b float64
	print("a, b = ")
	fmt.Scan(&a, &b)
	sqra, sqrb := a*a, b*b
	sum, diff, prod, div := sqra+sqrb, sqra-sqrb, sqra*sqrb, sqra/sqrb
	fmt.Printf("Sum = %.2f\nDifference = %.2f\nProduction = %.2f\nQuitient = %.2f\n",
		sum, diff, prod, div)

}

func begin9() {
	fmt.Println(`
	Three points A, B, C are given on the real axis, 
	the point C is located between the points A and B.
	Find the product of the length of AC and 
	the length of BC.
	`)
	var a, b, c float64
	print("a, b, c = ")
	fmt.Scan(&a, &b, &c)
	prod := math.Abs(a-c) * math.Abs(b-c)
	fmt.Printf("Product = %.2f", prod)
}

func begin10() {
	fmt.Println(`
	The coordinates (x1, y1) and (x2, y2) of two 
	opposite vertices of a rectangle are given.
	Sides of the rectangle are parallel to
	coordinate axes. Find the perimeter and 
	the area of the rectangle.
	`)
	var x1, x2, y1, y2 float64
	print("x1, y1 = ")
	fmt.Scan(&x1, &y1)
	print("x2, y2 = ")
	fmt.Scan(&x2, &y2)
	aside := math.Abs(y2 - y1)
	bside := math.Abs(x2 - x1)
	p, s := 2*(aside+bside), aside*bside
	fmt.Printf("P = %.2f\nS = %.2f", p, s)
}
