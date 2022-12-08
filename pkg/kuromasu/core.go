package kuromasu

import "fmt"

func Generate(w, h int) {
	result := generator(w, h)
	for _, r := range result.Solution {
		fmt.Println(r)
	}
}
