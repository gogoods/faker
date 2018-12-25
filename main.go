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

	▄████  ██   █  █▀ ▄███▄   █▄▄▄▄ 
	█▀   ▀ █ █  █▄█   █▀   ▀  █  ▄▀ 
	█▀▀    █▄▄█ █▀▄   ██▄▄    █▀▀▌  
	█      █  █ █  █  █▄   ▄▀ █  █  
	 █        █   █   ▀███▀     █   
	  ▀      █   ▀             ▀    
                ▀                                 						   
			   
Faker commands:

	enc <content>      Encode a string.
	dec <content>      Decode a string.

	conceal <content> [prefix]     Encode a string and add a faker prefix.
	reveal  <content> [prefix]     Decode a string if it contains a faker prefix.
`
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println(Usage)
	} else {
		method := os.Args[1]
		content := os.Args[2]

		var prefix string
		if len(os.Args) > 3 {
			prefix = os.Args[3]
		}

		switch method {
		case "enc":
			fmt.Println(Enc(content))
		case "dec":
			fmt.Println(Dec(content))
		case "reveal":
			if prefix == "" {
				fmt.Println(Reveal(content))
			} else {
				fmt.Println(Reveal(content, prefix))
			}

		case "conceal":
			if prefix == "" {
				fmt.Println(Conceal(content))
			} else {
				fmt.Println(Conceal(content, prefix))
			}

		default:
			fmt.Println(Usage)
		}

	}
}
