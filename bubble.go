package main

type BubbleSort struct{}

func (b BubbleSort) byAge(data []Player) []Player {
	length := len(data)

	for i := 1; i < length; i++ {
		// fmt.Printf("passes = %d \n", i)
		for j := 0; j < length-i; j++ {
			// fmt.Printf("compares = %d \n", j)
			if data[j].age > data[j+1].age {
				//swap
				temp := data[j]
				data[j] = data[j+1]
				data[j+1] = temp
			}
		}
	}
	return data
}

func (b BubbleSort) byName(data []Player) []Player {
	length := len(data)

	for i := 1; i < length; i++ {
		// fmt.Printf("passes = %d \n", i)
		for j := 0; j < length-i; j++ {
			// fmt.Printf("compares = %d \n", j)
			if data[j].name > data[j+1].name {
				//swap
				temp := data[j]
				data[j] = data[j+1]
				data[j+1] = temp
			}
		}
	}
	return data
}
