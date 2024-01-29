package maps 

import (

	"fmt"

)


func CreateMap(){

	//Map oluşturma, içine anahtar ve değer ekleme
	dictionary := make(map[string]string)
	dictionary["Add"] = "ekle"

	//Map oluştururken anahtarları ve değerleri ekleme
	dictionary2 := map[string]int{"fenerbahce" : 1907 , "Gs" : 1905}

	//Map içinden direk veri alma
	fmt.Println("fenerbahçe : " , dictionary2["fenerbahce"])

	//Map içinde deper var mı kontrolü?
	//Mapler 2 deger geri döndürürler bunlardan bir tanesi değer ikincisi ise içinde var mı yok mu true false degerleri
	_ , varMi := dictionary2["bjk"]
	fmt.Println("Var mi : " , varMi)

	//Direk mapleri yazdırma
	fmt.Println("Dictionary :" , dictionary)
	fmt.Println("Dictionary 2 :" , dictionary2)


}