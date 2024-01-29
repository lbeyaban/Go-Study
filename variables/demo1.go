package variables

import (

	"fmt"

)


func CreateVar() {

	var name string = "James"
	fmt.Println("Name : " , name)

}

func Condi(){

	var number int = 16
	var number2 int = 34

	if number < number2 {

		fmt.Println("Dogru")

	} else {

		fmt.Println("Yanlis")

	}


}

func LoopGame(){

	tuttugumSayi := 10
	soylenenSayi := 0


	for tuttugumSayi != soylenenSayi {
		fmt.Println("Tahminde Bulununuz.")
		fmt.Scanln(&soylenenSayi)
	}

	fmt.Println("Tebrikler Bildiniz.")


}

func IsPrime(number int){

	var isPrime bool = true
	for i := 2; i <= number / 2; i++ {

		if number % i == 0 {

			isPrime = false
			
		}

	}

	if isPrime == true {

		fmt.Println("Number is Prime")

	} else {

		fmt.Println("Number isnt Prime")

	}



} 

func FriendsNumbers(number int , number2 int) {

	var sum int = 0
	var sum2 int = 0

	for i := 1; i < number; i++ {
		
		if (number % i == 0) {

			sum += i
			fmt.Println("i : " , i )
			
		}

	}

	for j := 1; j < number2; j++ {
		
		if (number2 % j == 0) {

			sum2 += j
			fmt.Println("j : " , j )
			
		}

	}

	fmt.Println("Sum : " , sum)
	fmt.Println("Sum2 : " , sum2)

	if sum == number2 && sum2 == number {

		fmt.Println("Numbers are friends number")

	} else {

		fmt.Println("Numbers arent friends number")

	}


}

func Arrays(){

	numbers := [5]int{1,2,3,4,5}

	fmt.Println("Numbers :" , numbers)


}

func BiggestNumberInArray(){

	numbers := [5]int{2,5,12,4,20}

	enBuyuk := numbers[0]

	for i := 1; i < len(numbers); i++ {	
		
		if numbers[i] > enBuyuk {

			enBuyuk = numbers[i]

		}
		
	}

	fmt.Println("En buyuk sayi : " , enBuyuk)

}

func Slices(){

	slice := make([]int, 3, 5)
	capacity := cap(slice)

	slice = append(slice , 1 , 2, 3, 4, 5, 6)
	
	fmt.Println(capacity)
	fmt.Println(slice)

}


