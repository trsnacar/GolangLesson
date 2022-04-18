package slices

import "fmt"

func Demo1() {
	isimler := make([]string, 3)

	fmt.Println(isimler)
	isimler[0] = "Trusan"
	isimler[1] = "Miray"
	isimler[2] = "Acar"
	isimler = append(isimler, "Çakıroğlu")
	fmt.Println(isimler)
	fmt.Println(len(isimler))
}
