package main

import "testing"

func BenchmarkSortByAge(b *testing.B) {
	// run the Fib function b.N times
	sort := BubbleSort{}

	data := []Player{
		Player{age: 12, name: "Ass"},
		Player{age: 20, name: "Frank"},
		Player{age: 2, name: "Dave"},
		Player{age: 63, name: "Carol"},
		Player{age: 45, name: "Bee"}}
	for n := 0; n < b.N; n++ {
		sort.byAge(data)
	}
}

func BenchmarkSortByName(b *testing.B) {
	// run the Fib function b.N times
	sort := BubbleSort{}

	data := []Player{
		Player{age: 12, name: "Ass"},
		Player{age: 20, name: "Frank"},
		Player{age: 2, name: "Dave"},
		Player{age: 63, name: "Carol"},
		Player{age: 45, name: "Bee"}}
	for n := 0; n < b.N; n++ {
		sort.byName(data)
	}
}
