package main

import (
	"flag"
	"fmt"
	"os"
)

func main2() {

	method := flag.String("m", "", "Method: 'enc' or 'dec'.")
	content := flag.String("c", "", "Content to encode or decode.")

	main2()

	flag.Parse()

	if *method == "enc" {
		fmt.Println(Enc(*content))
	} else if *method == "dec" {
		fmt.Println(Dec(*content))
	} else {
		flag.Usage()
		os.Exit(0)
	}

}

const (
	Usage = `
	_____        __                 
	_/ ____\____  |  | __ ___________ 
	\   __\\__  \ |  |/ // __ \_  __ \
	 |  |   / __ \|    <\  ___/|  | \/
	 |__|  (____  /__|_ \\___  >__|   
				\/     \/    \/      
				
Faker commands:

	enc <content>      Encode a string.
	dec <content>      Decode a string.

	conceal <content>      Encode a string and add a faker prefix.
	receal  <content>      Decode a string if it contains a faker prefix.
`
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println(Usage)
	} else {
		method := os.Args[1]
		content := os.Args[2]

		switch method {
		case "enc":
			fmt.Println(Enc(content))
		case "dec":
			fmt.Println(Dec(content))
		case "reveal":
			fmt.Println(Reveal(content))
		case "conceal":
			fmt.Println(Conceal(content))
		default:
			fmt.Println(Usage)
		}

	}
}
