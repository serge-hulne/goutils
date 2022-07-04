package main

import (
	fi "github.com/serge-hulne/goutils/files"
	it "github.com/serge-hulne/goutils/iter"
)

func main() {

	a := []int{1, 2, 3}
	println(a)
	b := it.Iterable_from_Array(a)

	for item := range b {
		println(item)
	}

	println("- - - ")

	_, txt := fi.Read("test.txt")
	println(txt)

	_, f := fi.Open("test.txt")

	lines := fi.ReadLines(f)

	println("- - - ")

	for line := range lines {
		println(line)
	}

}
