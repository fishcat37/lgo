package main

import "fmt"

func Reverse(s string) string {
	temp := []rune(s)
	for i := 0; i < len(temp)/2; i++ {
		x := temp[i]
		temp[i] = temp[len(temp)-i-1]
		temp[len(temp)-i-1] = x
	}
	p := string(temp)
	return p
}
func IsPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	a := IsPalindrome("abca")
	fmt.Println(a)
	b := Reverse("acd")
	fmt.Println(b)
}
