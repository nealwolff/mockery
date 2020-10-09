package main

import "fmt"

func main() {
	theType := TestType{}
	theType.AssignType("passing in data using real implementation")
	fmt.Println(theType.GetData())
}
