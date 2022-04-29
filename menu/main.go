package main

import "fmt"

func init() {
	fmt.Println("\nwelcome to golang menu!")
	fmt.Println("***********************")
}

func exit() {
	fmt.Println("byb!")
}

func main() {

	defer exit()

	var cmd string
	for {
		fmt.Println("Please input a command here:")
		fmt.Scan(&cmd)
		switch cmd {
		case "help":
			fmt.Println("hello help quit")
			break
		case "hello":
			fmt.Println("hello")
			break
		case "quit":
			return
		default:
			fmt.Println("your input error! you can input help to know more!")
			break
		}
	}
}
