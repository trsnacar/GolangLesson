package conditionals

import "fmt"

func Demo1() {
	var hesap float64 = 1000
	var cekilmekIstenen float64 = 900

	var durum bool
	newFunction(durum, cekilmekIstenen, hesap)

	if cekilmekIstenen > hesap {
		fmt.Printf("Hesaptaki para yetersiz")
	}

	if cekilmekIstenen <= hesap {
		fmt.Println("Paranız hazırlanıyor")
		hesap = hesap - cekilmekIstenen
	}

	fmt.Println("Bitti. Hesaptaki para : " + fmt.Sprintf("%v", hesap))
}

func newFunction(durum bool, cekilmekIstenen float64, hesap float64) {
	durum = cekilmekIstenen > hesap
}
