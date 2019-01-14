package slices

import "fmt"

func Sort() {
	fmt.Println("Sort string by length of string")
	strings := []string{"ab", "cde", "fghi", "jk"}
	fmt.Println("Strings to sort:", strings)
	fmt.Println("Solution 1: use a map")
	res := make(map[int][]string)
	for _, v := range strings {
		i := len(v)
		if res[i] == nil {
			res[i] = make([]string, 3)
		}
		res[i] = append(res[i], v)
	}
	for _, v := range res {
		fmt.Println(v)
	}

	fmt.Println("Solution 1: use a map")
	maxLength := 0
	for _, v := range strings {
		if maxLength < len(v) {
			maxLength = len(v)

		}
	}
	fmt.Println(maxLength)
	res2 := make([][]string, maxLength)
	for _, v := range strings {
		i := len(v) - 1
		res2[i] = append(res2[i], v)
	}
	fmt.Println("%v", res2)
	for _, v := range res2 {
		fmt.Println(v)
	}
}
