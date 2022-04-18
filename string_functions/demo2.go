package string_functions

import (
	"fmt"
	s "strings"
)

func Demo2() {
	isim := "Trusan"
	fmt.Println(s.HasPrefix(isim, "Trus"))
	fmt.Println(s.HasSuffix(isim, "an"))
	fmt.Println(s.Index(isim, "an"))
	harfler := []string{"t", "r", "u", "s", "a", "n"}
	sonuc := s.Join(harfler, "*")
	fmt.Println(sonuc)

	fmt.Println(s.Replace(sonuc, "*", "+", -1))
	fmt.Println(s.Split(sonuc, "-"))
	fmt.Println(s.Repeat(sonuc, 5))
}
