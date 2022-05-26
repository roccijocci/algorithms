package main

import "fmt"

func main() {
	scores := []int{20, 10, 12, 60, 70}

	for i := 0; i < len(scores); i++ {
		//fmt.Println(scores[i])
		fmt.Printf("%d, \n", scores[i])
	}
}
