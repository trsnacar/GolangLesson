package arrays

import "fmt"

func Demo1() {
	var sayilar [5]int
	sayilar[2] = 50 //Burada dediğimiz şey sayıların 2. indexindeki elemana demek ve set işlemi gerçekleştirdir yani değeri atamak demek.Sayıların ikinci index'indeki eleman 50 demek istedim
	fmt.Println(sayilar[2])
}
