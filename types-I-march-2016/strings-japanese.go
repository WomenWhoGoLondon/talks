package main

import "fmt"

func main() {
	japanese := "語好きですよ！"
	fmt.Println(japanese, len(japanese))

	fmt.Println(len("語"))
	fmt.Println(len("好"))
	fmt.Println(len("き"))
	fmt.Println(len("で"))
	fmt.Println(len("す"))
	fmt.Println(len("よ"))
	fmt.Println(len("！"))
}
