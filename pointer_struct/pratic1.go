package pointer_struct

import (
	"fmt"
)

type Product struct {
	name  string
	price float64
	brand string
}

func (p Product) save() {

	fmt.Println("Saved : ", p.name)

}

func IncNumber(number *int) {

	*number++

}

func CreatePointer() {

	number := 10
	IncNumber(&number)
	fmt.Println("Number : ", number)

	pr := Product{name: "Bilgisayar", price: 26500}

	pr.save()

	pr2 := Product{
		name:  "Telefon",
		price: 16500,
	}

	pr2.save()

}
