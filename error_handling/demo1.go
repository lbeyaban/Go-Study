package error_handling

import (
	"errors"
	"fmt"
	"os"
)

// Hatamız için struct bir yapı oluşturuyoruz.
// İçinde parametre ve mesaj değişkenleri var
type specialException struct {
	message   string
	parameter int
}

// Hata structımıza interface olarak Error veriyoruz.
// Bu sayede onu hata olarak tanımlayabiliyoruz.
func (s *specialException) Error() string {
	return fmt.Sprintf("%d-------%s", s.parameter, s.message)
}

// Go dilinde hataları nil yani boş mu değil mi diye kontrol ederek ayıklarız.
// Hata döndürebilecek fonksiyonlar err değişkenine değer atar
func CatchErr() {

	f, err := os.Open("deneme.txt")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(f.Name())
	}

}

// Type assertion ile hata yakalama
func CatchErr2() {

	f, err := os.Open("deneme.txt")

	if err != nil {

		if pErr, ok := err.(*os.PathError); ok {
			fmt.Println("File Path couldn't find : ", pErr.Path)
		} else {
			fmt.Println(err)
		}

	} else {
		fmt.Println(f.Name())
	}

}

// Özel olarak error fırlatabiliyoruz.
func A(number int) (string, error) {

	if number > 100 || number < 0 {

		return "", errors.New("1 - 100 arasında bir sayi giriniz.")

	}

	return "Bildiniz", nil

}

// Özel struct hata yapımız.
func B(number int) (string, error) {

	if number > 100 || number < 0 {

		return "", &specialException{parameter: 14, message: "Sınırların dışında kaldınız."}

	}

	return "Bildiniz", nil

}

func Demo1() {

	fmt.Println(B(105))

}
