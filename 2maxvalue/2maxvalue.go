package main

import (
	"fmt"
)

//declare func maxvalue that takes 2 values and
//returns an int
func maxvalue(arrays []int, length int) int {
	for i := 0; i < length-1; i++ {
		if arrays[i] > arrays[i+1] {
			temp := arrays[i]

			arrays[i] = arrays[i+1]
			arrays[i+1] = temp
		}
	}
	maxvalue := arrays[length-1]
	return maxvalue
}

func main() {
	scores := []int{100, 50, 32, 15, 89, 32}
	length := len(scores)
	maxvalue := maxvalue(scores, length)

	fmt.Println(maxvalue)

}
