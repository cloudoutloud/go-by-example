package main

import (
	"fmt"
	"math"
	
)
	
// act like a var but can define outside func
const s string = "constant"

func main()  {
	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}

