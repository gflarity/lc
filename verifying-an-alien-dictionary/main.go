package main

import (
	"fmt"
	"math"
)

func isAlienSorted(words []string, order string) bool {
	// edge case only one word
	if len(words) == 0 {
		return true
	}

	mp := generateMapping(order)

	// use the lessThanEqual function to check that each pair in order
	// if this is true then the the array is sorted
	for i := 0; i < len(words)-1; i++ {
		if !compare(words[i], words[i+1], mp) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isAlienSorted([]string{"hello", "leetcode"}, "hlabcdefgijkmnopqrstuvwxyz"))
	fmt.Println(isAlienSorted([]string{"word", "world", "row"}, "worldabcefghijkmnpqstuvxyz"))
	fmt.Println(isAlienSorted([]string{"apple", "app"}, "abcdefghijklmnopqrstuvwxyz"))
	fmt.Println(isAlienSorted([]string{"app", "apple"}, "abcdefghijklmnopqrstuvwxyz"))

}

func compare(a string, b string, mp map[rune]int) bool {
	min := int(math.Min(float64(len(a)), float64(len(b))))
	ar := []rune(a)
	br := []rune(b)
	for i := 0; i < min; i++ {
		if mp[ar[i]] < mp[br[i]] {
			return true
		}
		if mp[ar[i]] > mp[br[i]] {
			return false
		}
		// else they are equal
	}

	// if we made it here the substrings are equal
	// if the strings are the same this is in order
	if len(a) == len(b) {
		return true
	}

	// if the first string is a short substring of the second, this is in order
	if len(a) < len(b) {
		return true
	}

	return false
}

func generateMapping(order string) map[rune]int {
	runes := []rune(order)
	m := make(map[rune]int, len(runes))
	for i, r := range runes {
		m[r] = i
	}
	return m
}
