package main

import "fmt"

func main() {
	var input1 int
	var input2 int
	fmt.Print("masukan angka alas =")
	fmt.Scan(&input1)
	fmt.Print("masukan angka tinggi =")
	fmt.Scan(&input2)
	luasSegitiga(&input1, &input2)
}

func luasSegitiga(alas, tinggi *int) {

	luas := *alas * *tinggi / 2
	fmt.Println("luas = ", luas)
}
