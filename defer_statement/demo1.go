package defer_statement

import "fmt"

func A(number int) string {

	defer Log("A")
	if number%2 != 0 {
		return "Number is odd"
	} else {
		return "Number is even"
	}

}

func Log(nameOfFunc string) {
	fmt.Println(nameOfFunc + " Start")
}

func Demo1() {
	result := A(8)
	fmt.Println("Result : ", result)
}
