package string_functions

import (
	"fmt"
	s "strings"
)

//case sensitive
//ascii
func Demo1() {
	isim := "Trusan"
	fmt.Println(s.Count(isim, "r"))
	fmt.Println(s.Contains(isim, "c"))
	sonuc := s.Index(isim, "c")

	if sonuc != -1 {
		fmt.Println("c harfi var")
	} else {
		fmt.Println("c harfi yok")
	}

	fmt.Println(s.ToLower(isim))
	fmt.Println(s.ToUpper(isim))
}
