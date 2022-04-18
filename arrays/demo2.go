package arrays

import "fmt"

func Demo2() {
	var sehirler [5]string // 5 tane verdiğimiz için 5.'yide yazmak isteyebiliriz ama 0'dan başladığı için zaten toplamda 5 tane oluyor buna dikkat etmeliyiz.
	sehirler[0] = "Ankara"
	sehirler[1] = "İstanbul"
	sehirler[2] = "İzmir"
	sehirler[3] = "Adana"
	sehirler[4] = "Diyarbakır"
	fmt.Println(sehirler)
}
