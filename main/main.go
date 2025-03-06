package main

import (
	"fmt"

	"github.com/sahapranta/banglaconv"
)

func main() {

	fmt.Println(banglaconv.ToBengaliNumber(-1234))

	txt, _ := banglaconv.ToBengaliWord(-1234)
	fmt.Println(txt)
}
