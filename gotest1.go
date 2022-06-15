/*
Q1
*/

package main

import (
	"fmt"
)

func main() {
	c1 := []string{"Red", "Black", "White"}
	c2 := []string{"Black", "Yellow", "Orange"}
	l := len(c1) + len(c2)
	c3 := make([]string, l)
	m2 := make(map[string]int)
	//var s2 string
	for _, val := range c1 {

		if _, ok := m2[val]; ok {
			m2[val] += 1
		} else {
			m2[val] = 1
		}
	}
	for _, val := range c2 {
		if _, ok := m2[val]; ok {
			m2[val] += 1
		} else {
			m2[val] = 1
		}
	}

	for key, _ := range m2 {
		c3 = append(c3, key)
	}

	for _, val := range c3 {
		fmt.Println(val)
	}
}

// time complexity = O(m+n)
// Space Complexity = O(len(m2)+len(c3))
