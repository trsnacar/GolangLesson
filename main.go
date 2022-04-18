package main

import (
	"fmt"
	"golesson/project"
)

func main() {
	//variables.Demo1()
	//fmt.Print()
	//conditionals.Demo5()
	//loops.Demo3()
	//arrays.Demo4()
	//slices.Demo2()
	//functions.SelamVer("Trusan")
	//functions.SelamVer("Miray")
	//var sonuc = functions.Topla(2, 6)
	//fmt.Println(sonuc * 10)
	// sonuc1, sonuc2, sonuc3, _ := functions.DortIslem(10, 6)
	// fmt.Println("Toplam : ", sonuc1)
	// fmt.Println("Çıkarım : ", sonuc2)
	// fmt.Println("Çarpım : ", sonuc3)
	//fmt.Println("Bölüm : ", sonuc4)
	// fmt.Println(functions.ToplaVariadic(1, 4, 6, 3, 10))
	// fmt.Println(functions.ToplaVariadic(1, 4))
	// fmt.Println(functions.ToplaVariadic())

	// sayilar := []int{4, 6, 7, 0, 11}
	// fmt.Println(functions.ToplaVariadic(sayilar...))

	// sayi := 20
	// pointers.Demo1(&sayi)
	// fmt.Println("Demodaki sayi : ", sayi)

	//sayilar := []int{1, 2, 3}
	//pointers.Demo2(sayilar)
	//fmt.Println("Maindeki sayi : ", sayilar[0])

	//structs.Demo2()
	// ciftSayiCn := make(chan int)
	// tekSayiCn := make(chan int)
	// go channeles.CiftSayilar(ciftSayiCn)
	// go channeles.TekSayilar(tekSayiCn)

	// ciftSayiToplam, tekSayiToplam := <-ciftSayiCn, <-tekSayiCn

	// carpim := ciftSayiToplam * tekSayiToplam
	// fmt.Println("Çarpım : ", carpim)

	//fmt.Println(error_handling.TahminEt2(99))

	//restful.Demo2()

	product, _ := project.AddProduct()
	fmt.Println(product)
	products, _ := project.GetAllProducts()

	for i := 0; i < len(products); i++ {
		fmt.Println(products[i].ProductName)
	}

}
