package interfaces

import "fmt"

type creditCal interface {
	calculate() float64
}

type Mortgage struct {
	creditPaymentTotal float64
	adress             string
	rate               float64
}

type Car struct {
	creditPaymentTotal float64
	carInfo            string
	rate               float64
}

func (m Mortgage) calculate() float64 {
	return m.creditPaymentTotal * m.rate / 100 / 12
}

func (c Car) calculate() float64 {
	return c.creditPaymentTotal * c.rate / 100 / 12
}

func CalMonthlyPayment(credits []creditCal) float64 {

	paymentTotal := 0.0

	for i := 0; i < len(credits); i++ {

		paymentTotal += credits[i].calculate()

	}

	return paymentTotal

}

func Demo2() {

	credi1 := Mortgage{creditPaymentTotal: 100000, adress: "A", rate: 10}
	credi2 := Mortgage{creditPaymentTotal: 50000, adress: "A", rate: 12}

	crediArr := []creditCal{credi1, credi2}

	total := CalMonthlyPayment(crediArr)
	fmt.Println("Total Aylık Ödeme :", total)

}
