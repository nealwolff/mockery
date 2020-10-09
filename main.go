package main

import "fmt"

func main() {
	theType := TestType{}
	theType.AssignData("passing in data using real implementation")
	fmt.Println(theType.GetData())
}
