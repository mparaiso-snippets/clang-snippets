package main

import (
	"io/ioutil"
	"os"

	"github.com/mparaiso-snippets/snippets/go/utils"
)

func main() {
	alphabet := utils.NewAlphabet(os.Args[0])
	R := alphabet.R()
	count := make([]int, R)
	in, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		print("error reading stdin", err)
		os.Exit(1)
	}
	str := string(in)
	length := len(str)
	for i := 0; i < length; i++ {
		if alphabet.Contains(str[i]) {
			count[alphabet.ToIndex(str[i])]++
		}
	}
	for c := 0; c < R; c++ {
		println(string(alphabet.ToChar(c)), " ", count[c])
	}
}
