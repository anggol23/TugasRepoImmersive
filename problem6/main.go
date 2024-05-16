package main

import "fmt"

func main() {
	var baju int = 370000
	var diskon float64 = 15.0 / 100

	hargaDiskon := int(diskon * float64(baju))
	hargaAkhir := baju - hargaDiskon

	fmt.Println("Harga yang harus dibayar adalah Rp", hargaAkhir)
}
