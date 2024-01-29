package for_range

import (

	"fmt"

)

func CreateForRange(){

	//Go dilinde Rangeleri foreach yapısına benzetebiliriz. Syntax biraz farklı

	numbers := []int{1,2,3,4,5,6,7,8,9}
	sum := 0

	for i , number := range numbers {

		fmt.Print("Number : ", number)
		fmt.Println("  Index : ", i)
		if number % 2 != 0 {

			sum += number
		}

	}

	fmt.Println("Sum of numbers : " , sum)



}