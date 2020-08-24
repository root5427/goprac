package begin

import "fmt"

func CallBegin() {
	var id int = 0
	println("Enter the task number (1 to 10): ")
	fmt.Scanf("%d", &id)
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
	fmt.Println("This is begin!")
}

func begin2() {
	fmt.Println("This is begin!")

}

func begin3() {
	fmt.Println("This is begin!")

}

func begin4() {
	fmt.Println("This is begin!")

}

func begin5() {
	fmt.Println("This is begin!")

}

func begin6() {
	fmt.Println("This is begin!")

}

func begin7() {

	fmt.Println("This is begin!")
}

func begin8() {
	fmt.Println("This is begin!")
}

func begin9() {
	fmt.Println("This is begin!")
}

func begin10() {
	fmt.Println("This is begin!")
}
