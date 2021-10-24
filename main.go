package main

import (
	"fmt"
	"github.com/Saegl/lecture2lib"
)

func main() {
	fmt.Printf("Swap case of 'A' %c\n", lecture2lib.ChangeCharCase('A'))
	fmt.Printf("Swap case of 'z' %c\n", lecture2lib.ChangeCharCase('z'))

	r1, r2, err := lecture2lib.FindQuadraticRoots(1, -5, -9)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r1, r2)
	}
}
