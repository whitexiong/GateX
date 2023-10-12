package main

import "fmt"

func isIsomorphic(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	m := make(map[byte]byte)
	n := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		valM, existsM := m[s[i]]
		valN, existsN := n[t[i]]

		if !existsM {
			m[s[i]] = t[i]
		} else if valM != t[i] {
			return false
		}

		if !existsN {
			n[t[i]] = s[i]
		} else if valN != s[i] {
			return false
		}

	}

	return true
}

func containsDuplicate(nums []int) bool {
	m := make(map[int]int)
	for i, num := range nums {
		if _, ok := m[num]; ok {

			return false
		}
		m[num] = i
	}
	return true
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := make(map[byte]int)
	n := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; ok {
			m[s[i]] = m[s[i]] + 1
		} else {
			m[s[i]] = 1
		}
	}

	for i := 0; i < len(t); i++ {
		if _, ok := n[t[i]]; ok {
			n[t[i]] = n[t[i]] + 1
		} else {
			n[t[i]] = 1
		}
	}

	fmt.Println(m)
	fmt.Println(n)
	for i := 0; i < len(m); i++ {
		if m[t[i]] != n[t[i]] {
			return false
		}
	}

	return true
}
func main() {

}
