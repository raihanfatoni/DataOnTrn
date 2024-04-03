package main

import "fmt"

type Mobil struct {
	Ban [4]int
	PintuKanan,Pintukiri string
}

func (m Mobil) gantiBan(posisi int, jenis string){
	if posisi >= 0 && posisi < 4 {
        m.Ban[posisi] = jenisBan
		fmt.Println(posisi)
		fmt.Println(jenis)
    } else {
        fmt.Println("Posisi ban tidak valid.")
    }
}
func (m Mobil) ketukPintu(pintu string){
	if pintu == "kanan" || pintu == "KANAN"{
		fmt.Println("Knock")
	} else {
		fmt.Println("Beep")
	}
}
func (m Mobil) bukaPintu(pintu string){
	if pintu == "kanan" || pintu == "KANAN"{
		fmt.Println("Beep")
	} else {
		fmt.Println("Knock")
	}
}



func main() {
	var mobil Mobil
	mobil.PintuKanan = "beep"
    mobil.PintuKiri = "Knock"
	
    // Mengganti ban pada posisi 1 dengan ban kayu
	mobil.gantiBan(1, "Ban Kayu")
    mobil.GantiBan(2, "ban kayu")

    // Mengganti ban pada posisi 3 dengan ban besi
    mobil.GantiBan(3, "ban besi")

    // Mengendalikan perilaku dari kedua pintu
    mobil.KetukPintu("kanan")
    mobil.BukaPintu("kiri")

}