package main

import "fmt"

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func nonempty(strings []string) []string {
	s := strings[:0]
	for _, t := range strings {
		if t != "" {
			s = append(s, t)
		}
	}

	return s
}

func main() {
	n := [...]int{1, 5, 9, 13, 15, 0, 2, 7}
	reverse(n[:])
	fmt.Print(n)

	var runes []rune
	for _, r := range "Hello 中文" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes)
	fmt.Printf("%04X\n", runes)
	fmt.Printf("%X\n", "Hello 中文")

	runes = append(runes, runes...)
	fmt.Printf("%04X\n", runes)

	ss := []string{"abc", "", "123", "", "xxx"}
	fmt.Printf("%s", nonempty(ss))
}
