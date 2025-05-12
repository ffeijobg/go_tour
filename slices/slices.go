package main

import (
	"fmt"
	"strings"
)

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printSlice2(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)

	sa := []int{2, 3, 5, 7, 11, 13}

	sa = sa[1:4]
	fmt.Println(sa)

	sa = sa[:2]
	fmt.Println(sa)

	sa = sa[1:]
	fmt.Println(sa)

	sa = sa[:]
	fmt.Println(sa)

	sb := []int{2, 3, 5, 7, 11, 13}
	printSlice(sb)

	// Slice the slice to give it zero length.
	sb = sb[:0]
	printSlice(sb)

	// Extend its length.
	sb = sb[:4]
	printSlice(sb)

	// Drop its first two values.
	sb = sb[3:]
	printSlice(sb)

	var sc []int
	fmt.Println(sc, len(sc), cap(sc))
	if sc == nil {
		fmt.Println("nil!")
	}

	aa := make([]int, 5)
	printSlice2("a", aa)

	bb := make([]int, 0, 5)
	printSlice2("b", bb)

	cc := bb[:2]
	printSlice2("c", cc)

	dd := cc[2:5]
	printSlice2("d", dd)

	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	var sd []int
	printSlice(sd)

	// append works on nil slices.
	sd = append(sd, 0)
	printSlice(sd)

	// The slice grows as needed.
	sd = append(sd, 1)
	printSlice(sd)

	// We can add more than one element at a time.
	sd = append(sd, 2, 3, 4)
	printSlice(sd)
}
