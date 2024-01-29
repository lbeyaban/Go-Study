package defer_statement

import "fmt"

type Product struct {
	productName string
	price       float64
}

func (p Product) Save() {

	fmt.Println("Product Saved : ", p.productName)

}

//Defer
//Eğer biz bir fonksiyonun kesinlikle çalışmasını istiyorsak başına defer koyabiliriz.
//Loglama işlemlerinde kullanılabilir. Stack yapısı ile çalışır. FILO

func Demo2() {

	p := Product{productName: "Laptop", price: 12500}
	defer p.Save()
	fmt.Println("I am here")

}
