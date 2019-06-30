package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
  "strconv"

)

type Host struct {
	Text string
}

func main() {

	fmt.Println("List command : \n 1. init [jumlah] -> Membuat loker \n 2. input [Tipe Identitas] [No Identitas] -> Menyimpan barang \n 3. leave [No Loker] -> Kosongkan loker dengan nomor loker tersebut \n 4. find [No Identitas] -> Mencari nomor loker menggunakan No Identitas")

arraystatus := 0
var array2 []string

//Proses input
start :
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
    string1 := scanner.Text()
    arrayofstring := strings.Fields(string1)

    if (strings.EqualFold("init", arrayofstring[0]) == true) {
		fmt.Println(">> init initiated")
		//Buat Loker
		
		//cek jika jumlah ada setelah init
		jumlahtext := len(arrayofstring)

		if (jumlahtext < 2){
			fmt.Println("!> Jumlah belum dimasukan")
			goto start
		}

		if (jumlahtext > 2){
			fmt.Println("!> argumen setelah angka tidak akan digunakan")
		}

		// cek jika jumlah adalah angka
		jumlaharray := arrayofstring[1]
		if _, err := strconv.Atoi(jumlaharray); err != nil {
	    fmt.Printf("!> %q Bukan lah angka. Masukan angka \n", jumlaharray)
	    // jika bukan angka balik ke awal
	 	goto start
		}

		//jika angka
		jumlaharrayint, err := strconv.Atoi(jumlaharray) //convert jumlah string ke int
		_ = err //variable yang dikeluarkan tapi tidak dipakai
		
		array2 = make([]string, jumlaharrayint)
		arraystatus = 1
		fmt.Printf(">> Berhasil dibuat loker dengan jumlah %d \n", len(array2))
		goto start



	}else if (strings.EqualFold("input", arrayofstring[0]) == true) {
		fmt.Println(">> input initiated")
		//cek loker
		if (arraystatus == 0 ){
			fmt.Println("Jumlah loker belum ditentukan")
			goto start
		}


		//value checking
		/*
		for i := 0 ; i < len(array2) ; i++ {
			if (array2[i] == "") {
				indeks := i+1
				s := strconv.Itoa(indeks)
				array2[i][0] = []Host{Host{s}}
				array2[i][1] = []Host{Host{arrayofstring[1]}}
				array2[i][2] = []Host{Host{arrayofstring[2]}}
				

			}
		}
		*/

	}else if (strings.EqualFold("leave", arrayofstring[0]) == true) {
		fmt.Println(">> leave initiated")
	}else if (strings.EqualFold("find", arrayofstring[0]) == true) {
		fmt.Println(">> find initiated")
	}else if (strings.EqualFold("end", arrayofstring[0]) == true) {
		fmt.Println(">> end initiated")
		return
	}else {
		fmt.Println("!>command tidak diketahui")
		goto start
	}
    










}