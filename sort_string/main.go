package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	list := []string{"10", "9", "23", "8", "6", "2", "17"}
	fmt.Println(list)
	sort.Slice(list, func(i, j int) bool {
		numA, _ := strconv.Atoi(list[i])
		numB, _ := strconv.Atoi(list[j])
		return numA < numB
	})
	fmt.Println(list)
}
