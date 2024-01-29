package string_functions

import (
	"fmt"
	s "strings"
)

// Go dili case sensitive bir dildir.
func Demo1() {

	sentence := "Hello my friends."

	//String içerindeki aranan harf/kelime kaç defa geçtiğiniz verir.
	fmt.Println(s.Count(sentence, "e"))

	//String içerinde aranan kelime veya harfin olup olmadığını döndürür. Var ise "True" yok ise -1 döndürür.
	fmt.Println(s.Contains(sentence, "Hello"))

	//String içinde aranan kelime veya harfin ilk bulunduğu yerin index numarasını verir.
	fmt.Println(s.Index(sentence, "my"))

	//Stringi küçük harflere çevirir.
	fmt.Println(s.ToLower(sentence))

	//Stringi büyük harflere çevirir.
	fmt.Println(s.ToUpper(sentence))

	//Bir string'in başında var mı yok mu onu kontrol eder. Var ise "True" yok ise "False" döner.
	fmt.Println(s.HasPrefix(sentence, "Hello"))

	//Bir string'in sonunda var mı yok mu onu kontrol eder. Var ise "True" yok ise "False" döner.
	fmt.Println(s.HasSuffix(sentence, "s."))

	//Bir string dizisini birleştirir. 2. parametre birleştirirken araya bir şey koyup koymayacağımız.
	harfler := []string{"o", "m", "e", "r"}
	fmt.Println("Harfler Birleşmiş hali : ", s.Join(harfler, ""))

	//Bir stringin harf veya kelimesinin yerine yeni bir kelime veya harf koymaya yarar.
	//İlk parametre stringi 2. parametre eski harf / kelime 3. parametre yeni harf / kelime 4. parametre ise kaç tanesinde değişiklik olacak.
	fmt.Println("Cumlenin boşluklar yerine - getirilmiş hali : ", s.Replace(sentence, " ", "-", -1))

	//Bir stringi bir ayraça göre ayırır ve geriye bir string dizisi geri döner.
	adSoyad := "Omer AYDIN"
	kelimeler := s.Split(adSoyad, " ")
	fmt.Println("Dizi olarak ayrılmış hali : ", kelimeler)

}
