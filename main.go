package main

import "fmt"

/*
	Run the following command for benchmarking:

	> go test -benchmem -run=^$ -bench .

*/

func main() {
	sort := BubbleSort{}

	data := []Player{
		Player{age: 12, name: "Ass"},
		Player{age: 20, name: "Frank"},
		Player{age: 2, name: "Dave"},
		Player{age: 63, name: "Carol"},
		Player{age: 45, name: "Bee"}}

	fmt.Println(sort.byAge(data))
	fmt.Println(sort.byName(data))
}
