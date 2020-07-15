package main

import "fmt"

func main() {

	l, s := lcs("bcg", "abcd")
	fmt.Println(l)
	fmt.Println(s)
}

func lcs(a, b string) (int, string) {
	arunes := []rune(a)
	brunes := []rune(b)
	aLen := len(arunes)
	bLen := len(brunes)
	fmt.Println(aLen + 1)
	fmt.Println(bLen + 1)
	T := make([][]int, aLen+1)
	for i := 0; i < aLen+1; i++ {
		T[i] = make([]int, bLen+1)
	}

	fmt.Println(T)

	// row 0 and column 0 are initialized to 0 already
	for i := 0; i < aLen; i++ {
		for j := 0; j < bLen; j++ {
			if arunes[i] == brunes[j] {
				T[i+1][j+1] = T[i][j] + 1
			} else if T[i+1][j] > T[i][j+1] {
				T[i+1][j+1] = T[i+1][j]
			} else {
				T[i+1][j+1] = T[i][j+1]
			}
		}
	}

	// read the substring out from the matrix
	s := make([]rune, 0, T[aLen][bLen])
	for x, y := aLen, bLen; x != 0 && y != 0; {
		if T[x][y] == T[x-1][y] {
			x--
		} else if T[x][y] == T[x][y-1] {
			y--
		} else {
			s = append(s, arunes[x-1])
			x--
			y--
		}
	}
	// reverse string
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return len(s), string(s)
}
