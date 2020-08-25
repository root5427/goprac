package integer

import (
	"fmt"
)

//CallInteger - function that calls specific integer task
func CallInteger() {
	var id int = 0
	print("Enter the task number (1 to 10): ")
	fmt.Scan(&id)
	if id > 0 {
		switch id {
		case 1:
			integer1()
		case 2:
			integer2()
		case 3:
			integer3()
		case 4:
			integer4()
		case 5:
			integer5()
		case 6:
			integer6()
		case 7:
			integer7()
		case 8:
			integer8()
		case 9:
			integer9()
		case 10:
			integer10()
		}
	}
}

func integer1() {
	fmt.Println(`
	A distance L is given in centimeters. Find the amount of full
	meters of this distance (1 m = 100 cm). Use the operator of
	integer division.
	`)
	var l int
	print("l = ")
	fmt.Scan(&l)
	d := l / 100
	fmt.Printf("d = %d", d)
}

func integer2() {
	fmt.Println(`
	A file size is given in bytes. Find the amount of full Kbytes of
	this size (1 K = 1024 bytes). Use the operator of integer
	division
	`)
	var sizeB int
	print("size in bytes = ")
	fmt.Scan(&sizeB)
	sizekb := sizeB / 1024
	fmt.Printf("size in kb = %d", sizekb)
}
func integer3() {
	fmt.Print(`
	Two positive integers A and B are given (A > B). Segment of
	length A contains the greatest possible amount of inside
	segments of length B (without overlaps). Find the length of
	unused part of the segment A. Use the operator of taking the
	remainder after integer division.
	`)
	var a, b uint
	print("a, b = ")
	fmt.Scan(&a, &b)
	res := a % b
	fmt.Printf("res = %d", res)
}
func integer4() {
	fmt.Println(`
	A two-digit integer is given. Find the sum and the product of its
	digits.
	`)
	var digit int
	print("digit = ")
	fmt.Scan(&digit)
	sum := digit/10 + digit%10
	prod := (digit / 10) * (digit % 10)
	fmt.Printf("sum, prod = %d %d", sum, prod)
}
func integer5() {
	fmt.Println(`
	A three-digit integer is given. Using one operator of integer
	division find first digit of the given integer (a hundreds digit)
	`)
	var digit int
	print("digit = ")
	fmt.Scan(&digit)
	res := digit / 100
	fmt.Printf("res = %d", res)
}
func integer6() {
	fmt.Println(`
	A three-digit integer is given. Find the sum and the product of
	its digits.
	`)
	var digit int
	print("digit = ")
	fmt.Scan(&digit)
	d1 := digit / 100
	d2 := digit / 10 % 10
	d3 := digit % 10
	sum := d1 + d2 + d3
	prod := d1 * d2 * d3
	fmt.Printf("sum, prod = %d %d", sum, prod)
}
func integer7() {
	fmt.Println(`
	A three-digit integer is given. Output an integer obtained from
	the given one by reading it from right to left.
	`)
	var digit int
	print("digit = ")
	fmt.Scan(&digit)
	d1 := digit / 100
	d2 := digit / 10 % 10
	d3 := digit % 10
	reverse := d3*100 + d2*10 + d1
	fmt.Printf("reversed := %d", reverse)
}
func integer8() {
	fmt.Println(`
	A three-digit integer is given. Output an integer obtained from
	the given one by exchange a ones digit and a tens digit (for
	example, 123 will be changed to 132).
	`)
	var digit int
	print("digit = ")
	fmt.Scan(&digit)
	d1 := digit / 100
	d2 := digit / 10 % 10
	d3 := digit % 10
	res := d1*100 + d3*10 + d2
	fmt.Printf("res = %d", res)
}
func integer9() {
	fmt.Printf(`
	An integer greater than 999 is given. Using one operator of
	integer division and one operator of taking the remainder find a
	thousands digit of the given integer.
	`)
	var digit int
	print("digit = ")
	fmt.Scan(&digit)
	th := digit / 100 % 10
	fmt.Printf("thousands = %d", th)
}
func integer10() {
	fmt.Printf(`
	From the beginning of the day N seconds have passed (N is
	integer). Find an amount of full hours passed from the beginning
	of the day.
	`)
	var n int
	print("n = ")
	fmt.Scan(&n)
	hours := n / 3600
	fmt.Printf("hours = %d", hours)
}
