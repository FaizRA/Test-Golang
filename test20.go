package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"

)

func main() {

	// Input jumlah loker
	loker :
	fmt.Println("Masukan Jumlah loker : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
    jumlah := scanner.Text()

    //validasi jika input adalah nomor 
    // JIKA bukan angka
    if _, err := strconv.Atoi(jumlah); err != nil {
    fmt.Printf("\n%q Bukan lah angka. Masukan angka.\n", jumlah)
    goto loker
	}

	// JIKA Angka
	jumlaharray, err := strconv.Atoi(jumlah)

	//fmt.Println(err)
	//fmt.Println(jumlaharray)

	_ = err //variable yang dikeluarkan tapi tidak dipakai

	var array2 []int
	array2 = make([]int,jumlaharray)
	fmt.Printf("Berhasil dibuat loker dengan jumlah %d ", len(array2))






}