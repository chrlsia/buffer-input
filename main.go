package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func main(){
	//request user input
	fmt.Print("\nEnter Text: ")

	//store user input in a memory buffer
	scanner:=bufio.NewScanner(os.Stdin)

	//read user input
	scanner.Scan()

	//output an error if one occured
	if scanner.Err()!=nil{
		fmt.Println(scanner.Err())
	} else {
		// returns the input that has been read as a single string value
		fmt.Println(scanner.Text())
	}

	fmt.Print("\nEnter Text: ")
	scanner=bufio.NewScanner(os.Stdin)
	scanner.Scan()

	words:=strings.Fields(scanner.Text())

	if scanner.Err()!=nil{
		fmt.Println(scanner.Err())
	} else {
		for index, value:=range words{
			fmt.Printf("index: %d ,value: %s\n",index,value)
		}
	}
}