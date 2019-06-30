package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
)

func main() {

		scanner := bufio.NewScanner(os.Stdin)
	    fmt.Print("Enter your text: ")
        scanner.Scan()
        text := scanner.Text()
        fmt.Println("Your text was: ", text)
        words := strings.Fields(text) // jadikan array
        fmt.Println(words, len(words)) // [one two three four] 4
		for i := 0 ; i < len(words) ; i++{ //proses output input  satu persatu
			fmt.Println(words[i])
		}

		jawab := strings.EqualFold("salam", words[0]) // cek inputan pertama jika sama true lainnya false

		fmt.Println(jawab) //cek true false

		if (jawab == true) {
			fmt.Println("waalaikumsalam")
		}else {
			fmt.Println("BAJINGAN")
		}


}